package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/martinyonathann/grpc-eg-go/machine"
	"github.com/martinyonathann/grpc-eg-go/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC server should listen TCP conn.")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	machine.RegisterMachineServer(grpcServer, &server.MachineServer{})
	grpcServer.Serve(lis)
	log.Printf("Initializing gRPC server on port %d", &*port)
}
