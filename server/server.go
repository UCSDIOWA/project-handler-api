package main

import (
	"context"
	"log"
	"net"

	pb "github.com/UCSDIOWA/project-handler-api/protos"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"google.golang.org/grpc"
)

type server struct{}

type mongo struct {
	Operation *mgo.Collection
}

// DB is a pointer to mongo struct
var DB *mongo

func main() {
	// Host mongo server
	m, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()

	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

    //Using email to find users projects in database. Won't allow duplicate emails.
    DB.Operation.EnsureIndex(mgo.Index{
        Key:        []string{"email"},
        Unique:     true,
        DropDups:   true,
        Background: true,
        Sparse:     true,})

	// Host grpc server
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	log.Println("Hosting server on", listen.Addr().String())

	s := grpc.NewServer()
	pb.RegisterProjectCreatorAPIServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) CreateProject(ctx context.Context, crProjReq *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {

    //Get user projects
	projects := &pb.UserProjects{}
	err1 := DB.Operation.Find(bson.M{ "email": crProjReq.Email}).One(projects)

    //User projects not found
	if err1 != nil {
		//Create user projects
        projects.Email = crProjReq.Email
        DB.Operation.Insert(projects.Email)
	}

	//if a project was already found, then we can't create this project
	if  projects.Projects[crProjReq.Title] != nil {
		return &pb.CreateProjectResponse{Success: false}, nil
	}

	//otherwise, go ahead and create the project
	projects.Projects[crProjReq.Title] = &pb.Project{
			Title:          crProjReq.Title,
			Users:          []string{crProjReq.Email},
			Description:    crProjReq.Description,
			Deadline:       crProjReq.Deadline,
			Private:        crProjReq.Private,
			Announcements:  []string{},
			Size:           1,
			ProjectLeader:  crProjReq.Email,
			ProgressBar:    0,
			Done:           false,
			Milestones:     []string{}}

    //Insert to database
    err := DB.Operation.Update( bson.M{ "email": crProjReq.Email}, projects )

	if err != nil {
		return &pb.CreateProjectResponse{Success: false}, err
	}

	return &pb.CreateProjectResponse{Success: true}, nil

}
