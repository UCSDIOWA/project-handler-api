package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/globalsign/mgo/bson"

	pb "github.com/UCSDIOWA/project-handler-api/protos"
	"github.com/globalsign/mgo"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type server struct{}

type mongo struct {
	Operation *mgo.Collection
}

type getUserProjects struct {
	CurrentProjects []string `json:"currentprojects" bson:"currentprojects"`
}

type getAllProjects struct {
	Xid           string   `json:"xid" bson:"xid"`
	Title         string   `json:"title" bson:"title"`
	Projectleader string   `json:"projectleader" bson:"projectleader"`
	Percentdone   int32    `json:"percentdone" bson:"percentdone"`
	Groupsize     int32    `json:"groupsize" bson:"groupsize"`
	Isprivate     bool     `json:"isprivate" bson:"isprivate"`
	Tags          []string `json:"tags" bson:"tags"`
	Deadline      string   `json:"deadline" bson:"deadline"`
	Calendarid    string   `json:"calendarid" bson:"calendarid"`
	Description   string   `json:"description" bson:"description"`
	Done          bool     `json:"done" bson:"done"`
	Joinrequests  []string `json:"joinrequests" bson:"joinrequests"`
	Memberslist   []string `json:"memberslist" bson:"memberslist"`
	Milestones    []string `json:"milestones" bson:"milestones"`
	Announcements []string `json:"announcements" bson:"announcements"`
}

// Host mongo server. Updating scope of of m
var (
	m            *mgo.Session
	echoEndpoint = flag.String("echo_endpoint", "localhost:50052", "endpoint of project-handler-api")
)

// DB is a pointer to mongo struct
var DB *mongo

func main() {
	errors := make(chan error)

	go func() {
		errors <- startGRPC()
	}()

	go func() {
		flag.Parse()
		defer glog.Flush()

		errors <- startHTTP()
	}()

	for err := range errors {
		log.Fatal(err)
		return
	}
}

func startGRPC() error {
	// Host mongo server
	var err error
	m, err = mgo.Dial("mongodb://tea:cse110IOWA@ds159263.mlab.com:59263/tea")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()
	log.Println("Connected to MongoDB server")

	// Accessing user collection in tea database
	DB = &mongo{m.DB("tea").C("projects")}

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

	return err
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterProjectCreatorAPIHandlerFromEndpoint(ctx, gwmux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	log.Println("Listening on port 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/createproject", func(w http.ResponseWriter, r *http.Request) {
		Set(w, AccessControl{
			Origin:         "*",
			AllowedMethods: []string{"POST"},
		})
		Set(w, ContentType("application/json"))
	})
	mux.Handle("/", gwmux)

	herokuPort := os.Getenv("PORT")
	if herokuPort == "" {
		herokuPort = "8080"
	}

	return http.ListenAndServe(":"+herokuPort, mux)
}

/* This function creates a project for a given user */
func (s *server) CreateProject(ctx context.Context, request *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	DB = &mongo{m.DB("tea").C("projects")}
	request.Xid = bson.NewObjectId().Hex()
	err := DB.Operation.Insert(request)
	if err != nil {
		return &pb.CreateProjectResponse{Success: false}, nil
	}

	return &pb.CreateProjectResponse{Success: true}, nil
}

func (s *server) GetAllProjects(ctx context.Context, request *pb.GetAllProjectsRequest) (*pb.GetAllProjectsResponse, error) {
	var userProjects getUserProjects
	var allProjects []getAllProjects
	var response pb.GetAllProjectsResponse

	DB = &mongo{m.DB("tea").C("users")}
	err := DB.Operation.Find(bson.M{"email": request.Email}).One(&userProjects)
	if err != nil {
		return &pb.GetAllProjectsResponse{Success: false}, nil
	}

	DB = &mongo{m.DB("tea").C("projects")}
	iter := DB.Operation.Find(nil).Iter()
	err = iter.All(&allProjects)
	if err != nil {
		return &pb.GetAllProjectsResponse{Success: false}, nil
	}

	invalidProjects := make(map[int]bool)

	for _, currentXid := range userProjects.CurrentProjects {
		for i := 0; i < len(allProjects); i++ {
			if currentXid == allProjects[i].Xid {
				invalidProjects[i] = true
			}
		}
	}

	for i := 0; i < len(allProjects); i++ {
		if !invalidProjects[i] && !allProjects[i].Done && !allProjects[i].Isprivate {
			var newProject pb.Projects
			newProject.Xid = allProjects[i].Xid
			newProject.Title = allProjects[i].Title
			newProject.Projectleader = allProjects[i].Projectleader
			newProject.Percentdone = allProjects[i].Percentdone
			newProject.Groupsize = allProjects[i].Groupsize
			newProject.Isprivate = allProjects[i].Isprivate
			newProject.Tags = allProjects[i].Tags
			newProject.Deadline = allProjects[i].Deadline
			newProject.Calendarid = allProjects[i].Calendarid
			newProject.Description = allProjects[i].Description
			newProject.Done = allProjects[i].Done
			newProject.Joinrequests = allProjects[i].Joinrequests
			newProject.Memberslist = allProjects[i].Memberslist
			newProject.Milestones = allProjects[i].Milestones
			newProject.Announcements = allProjects[i].Announcements
			response.Projects = append(response.Projects, &newProject)
		}
	}

	response.Success = true

	return &response, nil
}

func (s *server) GetProjects(ctx context.Context, request *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	var allProjects []getAllProjects
	var response pb.GetProjectsResponse

	DB = &mongo{m.DB("tea").C("projects")}
	iter := DB.Operation.Find(nil).Iter()
	err := iter.All(&allProjects)
	if err != nil {
		return &pb.GetProjectsResponse{Success: false}, nil
	}

	for _, currentXid := range request.Xid {
		for i := 0; i < len(allProjects); i++ {
			if currentXid == allProjects[i].Xid {
				var newProject pb.Projects
				newProject.Xid = allProjects[i].Xid
				newProject.Title = allProjects[i].Title
				newProject.Projectleader = allProjects[i].Projectleader
				newProject.Percentdone = allProjects[i].Percentdone
				newProject.Groupsize = allProjects[i].Groupsize
				newProject.Isprivate = allProjects[i].Isprivate
				newProject.Tags = allProjects[i].Tags
				newProject.Deadline = allProjects[i].Deadline
				newProject.Calendarid = allProjects[i].Calendarid
				newProject.Description = allProjects[i].Description
				newProject.Done = allProjects[i].Done
				newProject.Joinrequests = allProjects[i].Joinrequests
				newProject.Memberslist = allProjects[i].Memberslist
				newProject.Milestones = allProjects[i].Milestones
				newProject.Announcements = allProjects[i].Announcements
				response.Projects = append(response.Projects, &newProject)
			}
		}
	}

	response.Success = true
	if len(response.Projects) == 0 {
		response.Success = false
	}

	return &response, nil
}
