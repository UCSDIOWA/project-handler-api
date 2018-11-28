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
	m, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()
	log.Println("Connected to MongoDB server")

	// Accessing user collection in tea database
	DB = &mongo{m.DB("tea").C("projects")}
	// Using email to find users in the database. Won't allow duplicated emails.
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

	return err
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterProjectCreatorAPIHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	log.Println("Listening on port 8080")

	herokuPort := os.Getenv("PORT")
	if herokuPort == "" {
		herokuPort = "8080"
	}

	return http.ListenAndServe(":"+herokuPort, mux)
}

/* This function creates a project for a given user */
func (s *server) CreateProject(ctx context.Context, request *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	request.Xid = bson.NewObjectId().Hex()
	err := DB.Operation.Insert(request)
	if err != nil {
		return &pb.CreateProjectResponse{Success: false}, nil
	}

	return &pb.CreateProjectResponse{Success: true}, nil
}

func (s *server) FetchProject(ctx context.Context, request *pb.FetchProjectRequest) (*pb.FetchProjectResponse, error) {
	var response pb.FetchProjectResponse
	err := DB.Operation.Find(bson.M{"xid": request.Xid}).One(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
