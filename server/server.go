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
	gRPC.UnimplementedChatServiceServer // .Unimplemented has to be there for Go reasons. Google it...
	name                                string
	port                                int
}

var (
	testServerName = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
	port           = flag.Int("port", 5400, "Server port number")   // set with "-port <port>" in terminal
)

func main() {

	flag.Parse()

	go startServer()

	// The loop makes the program keep going..

	for {

	}
}

func startServer() {

	server := &Server{
		name: "server1",
		port: *port,
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port)) //sets up remote server
	if err != nil {
		log.Fatalln("Could not start listener.")
	}

	log.Printf("Server started.")

	gRPC.RegisterChatServiceServer(grpcServer, server)

	serverError := grpcServer.Serve(listener)

	if serverError != nil {
		log.Printf("Could not register server")
	}

	grpcServer.Serve(listener)
}

func (c *Server) SendMessageAck(ctx context.Context, in *gRPC.Message, opts ...grpc.CallOption) (*gRPC.MessageAck, error) {
	log.Printf("Client with nickname %v is trying to send a message", in.Clientname)

	return &gRPC.MessageAck{
		MessageAck:  "server received message",
		LamportTime: 1,
	}, nil

}

func (c *Server) SendMessage(ctx context.Context, in *gRPC.Message) (*gRPC.MessageAck, error) {
	log.Printf("Client with nickname %v is trying to send a message", in.Clientname)

	log.Printf("Received message from %v: %v ", in.Clientname, in.Message)

	return &gRPC.MessageAck{
		MessageAck:  "server received message",
		LamportTime: 1,
	}, nil

}

// c *Server means thats

// func (c *Server) GetTime(ctx context.Context, in *gRPC.AskForClientName) (*gRPC.TimeMessage, error) {

// 	log.Printf("Client with nickname %v asked for the time", in.Clientname)

// 	return &gRPC.TimeMessage{
// 		Time:       time.Now().String(),
// 		ServerName: c.name,
// 	}, nil
// }
