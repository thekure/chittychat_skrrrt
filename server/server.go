package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	//"time"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"

	"google.golang.org/grpc"
)

type Server struct {
	// .Unimplemented has to be there for Go reasons. Google it...
	gRPC.UnimplementedTimeAskServiceServer
	name                    string
	port                    int
	clientConnectionStrings map[string]gRPC.TimeAskService_GetTimeServer
	listen                  net.Listener
}

var (
	testServerName          = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
	port                    = flag.Int("port", 5400, "Server port number")   // set with "-port <port>" in terminal
	counter                 = 1
	clientConnectionStrings = make(map[string]gRPC.TimeAskService_GetTimeServer)
)

func main() {

	flag.Parse()
	//Set up remote server
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))

	//listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port)) //sets up remote server
	if err != nil {
		log.Fatalln("Could not start listener.")
	}

	//Create server instance
	server := &Server{
		name:                    "server1",
		port:                    *port,
		clientConnectionStrings: clientConnectionStrings,
		listen:                  listener,
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// RegisterService registers a service and its implementation to the
	// concrete type implementing this interface.  It may not be called
	gRPC.RegisterTimeAskServiceServer(grpcServer, server)
	log.Println("Server started.")

	// The loop makes the program keep going..
	for {
		//Method listens for inputs and broadcasts message to all clients
		server.readConnection(grpcServer)
	}
}

// GetTime = SendMessages
func (s *Server) GetTime(stream gRPC.TimeAskService_GetTimeServer) error {

	var online bool = true
	for online {
		//reads messages from client from the stream
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalln("an error occurred")
		}

		// if it is the first time the client is sending a message to the server (a request to join)
		// the client name and server-stream is saved to the hashmap
		// this happends in the else statement.
		// If it isn´t the first time the server broadcast the message further to all stream
		// stored in the connectionStrings hashmap

		if _, ok := s.clientConnectionStrings[msg.GetClientname()]; ok {

			log.Printf("---Inside server: %v says %v ", msg.GetClientname(), msg.GetMessage())

			if msg.GetMessage() == "exit" {
				delete(s.clientConnectionStrings, msg.Clientname)
				log.Printf("%v disconnected", msg.GetClientname())
				for key := range s.clientConnectionStrings {
					s.clientConnectionStrings[key].Send(&gRPC.MessageAck{
						Message:    "left the chat",
						Clientname: msg.GetClientname(),
					})
				}
				online = false
				break
			} else {
				//broadcast message to all clients connected to server via the hashmap storing streams
				for key := range s.clientConnectionStrings {
					s.clientConnectionStrings[key].Send(&gRPC.MessageAck{
						Message:    "says: " + msg.GetMessage(),
						Clientname: msg.GetClientname(),
					})
				}
			}
		} else {

			//adds the "new client" to the hashmap by key = clientname, val = stream
			s.clientConnectionStrings[msg.GetClientname()] = stream
			log.Printf("%v joined the chatroom", msg.GetClientname())
			for key := range s.clientConnectionStrings {
				s.clientConnectionStrings[key].Send(&gRPC.MessageAck{
					Message:    "joined the chatroom",
					Clientname: msg.GetClientname(),
				})
			}
		}
	}
	return nil
}

func (s *Server) readConnection(grpcServer *grpc.Server) (*gRPC.MessageAck, error) {

	//accepts all incoming connections
	grpcServer.Serve(s.listen)

	return &gRPC.MessageAck{
		Message: "client succesfully requested connection",
	}, nil
}
