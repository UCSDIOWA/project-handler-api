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
	m, err := mgo.Dial("mongodb://tea:cse110IOWA@ds159263.mlab.com:59263/tea")
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
		Sparse:     true})

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

/* This function creates a project for a given user */
func (s *server) CreateProject(ctx context.Context, crProjReq *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {

	/* first, update the project's field for the User collection */
	DB = &mongo{m.DB("tea").C("users")} //change collection to projects

	//Get user projects
	usrProjects := &pb.UserProjects{}

	//get all user's project titles
	usrErr := DB.Operation.Find(bson.M{"email": crProjReq.Email}).One(usrProjects)

	//in case something goes awry
	if usrErr != nil {
		return &pb.CreateProjectResponse{Success: false}, usrErr
	}

	//if a project was already found, then we can't create this project
	for _, proj := range usrProjects.Projects {
		if proj == crProjReq.Title {
			return &pb.CreateProjectResponse{Success: false}, nil
		}
	}

	//otherwise, go ahead and add the project
	usrProjects.Projects = append(usrProjects.Projects, crProjReq.Title)

	//Insert to Users collection
	usrErr = DB.Operation.Update(bson.M{"email": crProjReq.Email}, usrProjects)

	if usrErr != nil {
		return &pb.CreateProjectResponse{Success: false}, usrErr
	}

	/* then, update the Projects collection */
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	//make sure it's not a duplicate project
	newProj := &pb.Project{}
	projErr := DB.Operation.Find(bson.M{"title": crProjReq.Title}).One(newProj)

	if (projErr != nil || newProj != {&pb.Project{})} {
		return &pb.CreateProjectResponse{Success: false}, projErr
	}

	//now create the new project entry
	newProj = &pb.Project{
		Title:         crProjReq.Title,
		Users:         []string{crProjReq.Email},
		Description:   crProjReq.Description,
		Deadline:      crProjReq.Deadline,
		Private:       crProjReq.Private,
		Announcements: []string{},
		Size:          1,
		ProjectLeader: crProjReq.Email,
		ProgressBar:   0,
		Done:          false,
		Milestones:    []string{},
		Tags:          crProjReq.Tags,
	}

	//Insert to the Projects collection
	projErr = DB.Operation.Update(bson.M{"title": crProjReq.Title}, newProj)

	if projErr != nil {
		return &pb.CreateProjectResponse{Success: false}, projErr
	}

	/* then, update the Tags collection */
	DB = &mongo{m.DB("tea").C("tags")} //change collection to projects

	//Insert the new title to each in tags collection
	tagProjs := &pb.TagProjects{}
	for _, v := range crProjReq.Tags {
		//see if tag is already in collection
		tagErr := DB.Operation.Find(bson.M{"name": v}).One(tagProjs)
		if tagErr != nil {
			return &pb.CreateProjectResponse{Success: false}, tagErr
		}
		//if tag doesn't exist, create it. Otherwise, update.
		if tagProjs == {&pb.TagProjects{}) {
			tagErr := DB.Operation.Insert(bson.M{"name": v, "Projects": []string{crProjReq.Title}})
			if tagErr != nil {
				return &pb.CreateProjectResponse{Success: false}, tagErr
			}
		} else {
			//append new project title to tag's array of project titles
			tagProjs.Projects = append(tagProjs.Projects, crProjReq.Title)
			tagErr := DB.Operation.Update(bson.M{"name": v}, tagProjs)
			if tagErr != nil {
				return &pb.CreateProjectResponse{Success: false}, tagErr
			}
		}
	}

	//if we made it here, we successfully updated all 3 collections, so return
	return &pb.CreateProjectResponse{Success: true}, nil

}

/* This function fetches a project which belongs to a certain user */
func (s *server) FetchProject(ctx context.Context, fProjReq *pb.FetchProjectRequest) (*pb.FetchProjectResponse, error) {

	//Get user projects
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	fetched := &pb.Project{}

	err := DB.Operation.Find(bson.M{"title": fProjReq.Title}).One(fetched)

	if err != nil {
		return &pb.FetchProjectResponse{Success: false, Project: &pb.Project{}}, err
	}

	//if a project not found, return nothing and but no error
	if fetched == (&pb.Project{}) {
		return &pb.FetchProjectResponse{Success: false, Project: &pb.Project{}}, nil
	}

	//otherwise, go ahead and return the project
	return &pb.FetchProjectResponse{Success: true, Project: fetched}, nil

}
