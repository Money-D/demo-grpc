package main

import (
	context "context"
	"log"
	"net"

	"https://github.com/Money-D/demo-grpc/testAPI"

	"google.golang.org/grpc"
)

type myTestAPIServer struct {
	testAPI.UnimplementedTestAPIServer
}

func (server mytestAPIServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return &testAPI.GetDataResponse{
		Pdf: []byte("test"),
		Png: []byte("test"),
	}, nil
}

func main() {
	// This opens a connection to receive requests and send reposnes
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myTestAPIServer{}
	testApi.RegisterTestAPIServer(serverRegistrar, service)
	err = serverRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("cannot serve: %s", err)
	}
}
