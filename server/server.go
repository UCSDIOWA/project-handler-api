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

	//first check if this project has already been created
	project := &pb.CreateProjectRequest{}
	err1 := DB.Operation.Find(
		bson.M{
			"title":      crProjReq.Title,
			"descripton": crProjReq.Description,
			"deadline":   crProjReq.Deadline,
			"private":    crProjReq.Deadline,
		}).One(project)

	if err1 != nil {
		return &pb.CreateProjectResponse{Success: false}, err1
	}

	//if a project was already found, then we can't create this project
	if (&pb.CreateProjectRequest{}) != project {
		return &pb.CreateProjectResponse{Success: false}, err1
	}

	//otherwise, go ahead and create the project
	err := DB.Operation.Insert(
		bson.M{
			"title":          crProjReq.Title,
			"Users":          [1]string{crProjReq.Email},
			"descripton":     crProjReq.Description,
			"deadline":       crProjReq.Deadline,
			"private":        crProjReq.Private,
			"Announcements":  []string{},
			"size":           1,
			"project_leader": crProjReq.Email,
			"progress_bar":   0,
			"done":           false,
			"Milestones":     []string{},
		})
	if err != nil {
		return &pb.CreateProjectResponse{Success: false}, err
	}
	return &pb.CreateProjectResponse{Success: true}, nil

}
