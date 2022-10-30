package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"

	"google.golang.org/grpc"
)

type Server struct {
	// .Unimplemented has to be there for Go reasons. Google it...
	gRPC.UnimplementedTimeAskServiceServer
	name      string
	port      int
	clientMap map[int]int
}

var (
	testServerName = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
	port           = flag.Int("port", 5400, "Server port number")   // set with "-port <port>" in terminal
	counter        = 1
)

func main() {

	flag.Parse()

	server := &Server{
		name:      "server1",
		port:      *port,
		clientMap: make(map[int]int),
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

	log.Println("1")
	log.Println(listener.Addr().String())
	log.Println(grpcServer.GetServiceInfo())
	k, _ := listener.Accept()
	l := k.LocalAddr()
	log.Println("3", k, l)

	gRPC.RegisterTimeAskServiceServer(grpcServer, server)
	serverError := grpcServer.Serve(listener)

	log.Println("2", grpcServer.GetServiceInfo())

	if serverError != nil {
		log.Printf("Could not register server")
	}

	grpcServer.Serve(listener)

}

// c *Server means thats
func (c *Server) GetTime(stream gRPC.TimeAskService_GetTimeServer) error {

	msg, err := stream.Recv()
	if err != nil {
		log.Fatalln("an error occurred")
	}

	log.Printf("1 %v says %v ", msg.GetClientname(), msg.GetMessage())

	return stream.Send(&gRPC.MessageAck{
		Message:    "Server received message" + msg.Message,
		Clientname: msg.GetClientname(),
	})
}

func (c *Server) BroadcastMessage(message gRPC.Message) (stream gRPC.TimeAskService_GetTimeServer) {

	input := message.GetMessage()

	log.Printf("Client input %s\n", input)

	stream, err := serverConnection.GetTime(context.Background(), grpc.CustomCodecCallOption{})
	if err != nil {
		log.Printf("an error occured in client class")
	}

	stream.Send(&gRPC.Message{
		Clientname: string(client.name),
		Message:    input,
		PortNumber: string(rune(client.portNumber)),
	})

	stream.CloseSend()
	var msg, _ = stream.Recv()
	log.Printf("%v from %v ", msg.Message, msg.Clientname)

}

// // c *Server means thats
// func (c *Server) GetTime(ctx context.Context, in *gRPC.Message) (*gRPC.MessageAck, error) {

// 	log.Printf("%v says %v ", in.Clientname, in.Message)

// 	// counter := 1
// 	// c.clientMap[in.Clientname] = counter
// 	// counter++

// 	//sætte nedenståenxe ind i et for loop
// 	return &gRPC.MessageAck{
// 		Clientname: in.Clientname,
// 		Message:    in.Message,
// 	}, nil
// }

func (c *Server) broadcastInput(message string) {
	//Take message and send it to every client
	for key, element := range c.clientMap {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
