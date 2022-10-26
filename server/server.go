package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"

	"google.golang.org/grpc"
)

type Server struct {
	// .Unimplemented has to be there for Go reasons. Google it...
	gRPC.UnimplementedTimeAskServiceServer
	name string
	port int
}

var port = flag.Int("port", 8080, "server port number")

func main() {
	flag.Parse()

	server := &Server{
		name: "server1",
		port: *port,
	}

	go startServer(server)

	// The loop makes the program keep going..

	for {

	}
}

func startServer(server *Server) {
	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalln("Could not start listener.")
	}

	log.Printf("Server started.")

	gRPC.RegisterTimeAskServiceServer(grpcServer, server)
	serverError := grpcServer.Serve(listener)

	if serverError != nil {
		log.Printf("Could not register server")
	}

	grpcServer.Serve(listener)
}

// c *Server means thats

func (c *Server) GetTime(ctx context.Context, in *gRPC.AskForTimeMessage) (*gRPC.TimeMessage, error) {
	log.Printf("Client with ID %d asked for the time \n", in.ClientId)

	return &gRPC.TimeMessage{
		Time:       time.Now().String(),
		ServerName: c.name,
	}, nil
}
