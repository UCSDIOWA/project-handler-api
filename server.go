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
	echoEndpoint = flag.String("echo_endpoint", "localhost:50052", "endpoint of project-handler")
}

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
	m, err := mgo.Dial("mongodb://tea:cse110IOWA@ds159263.mlab.com:59263/tea")
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

	return err
}

func startHTTP() error {
	/*ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterLoginAPIHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts) //should
	if err != nil {
		return err
	}
	log.Println("Listening on port 8080")

	herokuPort := os.Getenv("PORT")
	if herokuPort == "" {
		herokuPort = "8080"
	}

	return http.ListenAndServe(":"+herokuPort, mux)*/
	return error
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
		return &pb.CreateProjectResponse{Success: false}, nil
	}

	//if a project was already found, then we can't create this project
	if (&pb.CreateProjectRequest{}) != project {
		return &pb.CreateProjectResponse{Success: false}, nil
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
		return &pb.CreateProjectResponse{Success: false}, nil
	}
	return &pb.CreateProjectResponse{Success: true}, nil

}
