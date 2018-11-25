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

// Host mongo server. Updating scope of of m
var (
	m, err = mgo.Dial("mongodb://tea:cse110IOWA@ds159263.mlab.com:59263/tea")
)

// DB is a pointer to mongo struct
var DB *mongo

func main() {

	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()

	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	//Using email to find users projects in database. Won't allow duplicate emails.
	DB.Operation.EnsureIndex(mgo.Index{
		Key:        []string{"ProjectID"},
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

	//Get user projects. Also, we will be using this generated id as the index for this project
	usrProjects := &pb.UserProjects{}
	id := bson.NewObjectId()

	//get all user's project titles
	usrErr := DB.Operation.Find(bson.M{"email": crProjReq.Email}).One(usrProjects)

	//in case something goes awry
	if usrErr != nil {
		return &pb.CreateProjectResponse{Success: false}, usrErr
	}

	//if a project was already found, then we can't create this project
	for _, proj := range usrProjects.Projects {
		if proj == string(id) {
			return &pb.CreateProjectResponse{Success: false}, nil
		}
	}

	//otherwise, go ahead and add the project
	usrProjects.Projects = append(usrProjects.Projects, string(id))

	//Insert to Users collection
	usrErr = DB.Operation.Update(bson.M{"email": crProjReq.Email}, usrProjects)

	if usrErr != nil {
		return &pb.CreateProjectResponse{Success: false}, usrErr
	}

	/* then, update the Projects collection */
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	//make sure it's not a duplicate project
	newProj := &pb.Project{}
	projErr := DB.Operation.Find(bson.M{"_id": id}).One(newProj)

	if (projErr != nil || newProj != &pb.Project{}) {
		return &pb.CreateProjectResponse{Success: false}, projErr
	}

	//now create the new project entry
	newProj = &pb.Project{
		ProjectID:     string(id),
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
		Calendar:      "",
		Milestones:    []string{},
		Tags:          crProjReq.Tags,
		PendingUsers:  []string{},
	}

	//Insert to the Projects collection with the new unique id
	projErr = DB.Operation.Insert(bson.M{"_id": id}, newProj)

	if projErr != nil {
		return &pb.CreateProjectResponse{Success: false}, projErr
	}

	/* then, update the Tags collection */
	DB = &mongo{m.DB("tea").C("tags")} //change collection to projects

	//add project title to tags
	crProjReq.Tags = append(crProjReq.Tags, newProj.ProjectID)

	//Insert the new title to each in tags collection
	tagProjs := &pb.TagProjects{}
	for _, v := range crProjReq.Tags {
		//see if tag is already in collection
		tagErr := DB.Operation.Find(bson.M{"Name": v}).One(tagProjs)
		if tagErr != nil {
			return &pb.CreateProjectResponse{Success: false}, tagErr
		}
		//if tag doesn't exist, create it. Otherwise, update.
		if (tagProjs == &pb.TagProjects{}) {
			tagErr := DB.Operation.Insert(bson.M{"Name": v, "Projects": []string{newProj.ProjectID}})
			if tagErr != nil {
				return &pb.CreateProjectResponse{Success: false}, tagErr
			}
		} else {
			//append new project title to tag's array of project titles
			tagProjs.Projects = append(tagProjs.Projects, newProj.ProjectID)
			tagErr := DB.Operation.Update(bson.M{"Name": v}, tagProjs)
			if tagErr != nil {
				return &pb.CreateProjectResponse{Success: false}, tagErr
			}
		}
	}

	//if we made it here, we successfully updated all 3 collections, so return
	return &pb.CreateProjectResponse{Success: true}, nil

}

/* This function creates a project for a given user */
func (s *server) EditProject(ctx context.Context, edProjReq *pb.EditProjectRequest) (*pb.EditProjectResponse, error) {

	/* Update the Projects collection */
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	//Update the correct project in the Projects collection
	err := DB.Operation.Update(bson.M{"ProjectID": edProjReq.ProjectID}, edProjReq)

	if err != nil {
		return &pb.EditProjectResponse{Success: false}, err
	}

	//if we made it here, we successfully edited all 3 collections, so return
	return &pb.EditProjectResponse{Success: true}, nil

}

/* This function creates a project for a given user */
func (s *server) JoinProject(ctx context.Context, jProjReq *pb.JoinProjectRequest) (*pb.JoinProjectResponse, error) {

	/* Update the Projects collection */
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	//Check current users and pending users
	currProj := &pb.Project{}
	err := DB.Operation.Find(bson.M{"ProjectID": jProjReq.ProjectID}).One(currProj)
	if err != nil {
		return &pb.JoinProjectResponse{Success: false}, err
	}
	//make sure the user is not already in the group
	for _, v := range currProj.Users {
		if v == jProjReq.NewEmail {
			return &pb.JoinProjectResponse{Success: false}, err
		}
	}
	//make sure the user is not already pending
	for _, v := range currProj.PendingUsers {
		if v == jProjReq.NewEmail {
			return &pb.JoinProjectResponse{Success: false}, err
		}
	}

	//add the new email to the pending list
	currProj.PendingUsers = append(currProj.PendingUsers, jProjReq.NewEmail)
	err = DB.Operation.Update(bson.M{"ProjectID": jProjReq.ProjectID}, bson.M{"PendingUsers": currProj.PendingUsers})

	//if we made it here, we successfully added the user to pending users list
	return &pb.JoinProjectResponse{Success: true}, nil

}

/* This function fetches a project which belongs to a certain user */
func (s *server) FindProjects(ctx context.Context, findProjReq *pb.FindProjectsRequest) (*pb.FindProjectsResponse, error) {

	//used to access collections
	DB = &mongo{m.DB("tea").C("tags")} //change collection to tags

	rtnProjects := &pb.FindProjectsResponse{}

	//map of projects and their counts
	projs := make(map[string]bool)

	//go through each Tag that the user searched by
	for _, currTag := range findProjReq.Tags {
		//get the projects pertaining to this tag
		currProjs := &pb.TagProjects{}
		err := DB.Operation.Find(bson.M{"Name": currTag}).One(currProjs)
		if err != nil {
			//return false and empty array
			rtnProjects = &pb.FindProjectsResponse{}
			rtnProjects.Success = false
			return rtnProjects, err
		}
		//go through each project and add each unique one to our return Response
		for _, projID := range currProjs.Projects {
			//check if we've already found this project
			exist := projs[projID]
			if exist {
				continue
			}
			projs[projID] = true
			//get project information from projects collection
			DB = &mongo{m.DB("tea").C("projects")}
			currProj := &pb.Project{}
			err := DB.Operation.Find(bson.M{"ProjectID": projID}).One(currProj)
			if err != nil {
				//return false and empty array
				rtnProjects = &pb.FindProjectsResponse{}
				rtnProjects.Success = false
				return rtnProjects, err
			}
			//add this project information to our return response
			rtnProjects.Projects = append(rtnProjects.Projects, currProj)
		}
		//go back to tags collection
		DB = &mongo{m.DB("tea").C("tags")} //change collection to tags

	}

	//otherwise, go ahead and return the project
	return rtnProjects, nil

}

/* This function fetches a project which belongs to a certain user */
func (s *server) FetchProject(ctx context.Context, fProjReq *pb.FetchProjectRequest) (*pb.FetchProjectResponse, error) {

	//Get user projects
	DB = &mongo{m.DB("tea").C("projects")} //change collection to projects

	fetched := &pb.Project{}

	err := DB.Operation.Find(bson.M{"ProjectID": fProjReq.ProjectID}).One(fetched)

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
