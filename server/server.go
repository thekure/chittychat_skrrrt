package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"

	"google.golang.org/grpc"
)

type Server struct {
	// .Unimplemented has to be there for Go reasons. Google it...
	gRPC.UnimplementedTimeAskServiceServer
	name string
	port int
}

var (
	testServerName = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
	port           = flag.Int("port", 5400, "Server port number")   // set with "-port <port>" in terminal
)

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

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	//listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port)) //sets up remote server
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
func (c *Server) GetTime(ctx context.Context, in *gRPC.Message) (*gRPC.MessageAck, error) {

	log.Printf("%v says %v ", in.Clientname, in.Message)

	return &gRPC.MessageAck{
		Clientname: in.Clientname,
		Message:    in.Message,
	}, nil
}
