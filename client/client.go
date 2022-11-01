package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"
)

type Client struct {
	id               int
	name             string
	portNumber       int
	connectionString string
	stream           gRPC.TimeAskService_GetTimeClient
	lamportTime      int64
}

var (
	clientPort = flag.Int("clientPort", 8081, "client port number")
	serverPort = flag.Int("serverPort", 8080, "server port number")
)

// go run server/server.go -port=8083

func main() {

	fmt.Println("Please enter a nickname")
	scannerName := ""
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Split(bufio.ScanBytes)

	for scanner1.Scan() {
		if scanner1.Text() == "\n" {
			break
		} else {
			// If the character is not \n, add to the input buffer
			scannerName += scanner1.Text()
		}
	}

	flag.Parse()

	client := &Client{
		name:        scannerName,
		portNumber:  *clientPort,
		lamportTime: 0,
	}

	go startClient(client)

	for {

	}
}

// Method that setup the client - starts the send- and recieveMessage method
func startClient(client *Client) {
	serverConnection, err := getServerConnection()

	stream, err := serverConnection.GetTime(context.Background(), grpc.CustomCodecCallOption{})
	if err != nil {
		log.Printf("4, an error occured in client class")
	}

	//Setup of the stream
	client.stream = stream

	stream.Send(&gRPC.Message{
		Clientname:       client.name,
		Message:          client.name + " joined the chatroom",
		LamportTimestamp: client.lamportTime,
	})
	// rqConn(client, serverConnection)

	if err != nil {
		log.Printf("Error..")
	}
	//Method that handles the sending-part of a client -> sends message to server
	go sendMessage(client, serverConnection)
	//Method that handles the recieving-part of a client <- recieves message from server
	go receiveMessage(client, client.stream)

	for {

	}
}

func receiveMessage(client *Client, stream gRPC.TimeAskService_GetTimeClient) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println("error")
		}

		if client.name != msg.GetClientname() {
			log.Println(client.lamportTime, " - ", msg.GetLamportTimestamp())
			if msg.GetLamportTimestamp() > client.lamportTime {
				client.lamportTime = msg.GetLamportTimestamp() + 1
			} else {
				client.lamportTime++
			}
		} else {
			client.lamportTime++
		}
		log.Printf("%v %v, lamport: %v", msg.GetClientname(), msg.GetMessage(), client.lamportTime)

	}
}

func sendMessage(client *Client, serverConnection gRPC.TimeAskServiceClient) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "exit" {
			client.stream.Send(&gRPC.Message{
				Message:    "exit",
				Clientname: client.name,
			})
			os.Exit(0)

		} else {
			log.Printf("(Message sent from this client: '%s')", input)
			client.stream.Send(&gRPC.Message{
				Clientname:       client.name,
				Message:          input,
				LamportTimestamp: client.lamportTime,
			})
		}
	}
}

func getServerConnection() (gRPC.TimeAskServiceClient, error) {
	//conn, err := grpc.Dial(":"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//dial the server to get a connection to it
	//conn, err := grpc.DialContext(timeContext, fmt.Sprintf(":%s", *serverPort), opts...)

	// conn, err := grpc.Dial("192.168.0.145:5400", opts...) //smusHosKure
	// conn, err := grpc.Dial("172.20.10.5:5400", opts...) //smusiphone
	// conn, err := grpc.Dial("192.168.8.112:5400", opts...) //bemihjem
	// conn, err := grpc.Dial("10.26.16.44:5400", opts...) //itu
	conn, err := grpc.Dial("192.168.0.153:5400", opts...) //smusHjemme
	// conn, err := grpc.Dial("172.20.10.6:5400", opts...) //bemiiphone
	// conn, err := grpc.Dial("192.168.0.110:5400", opts...) //kure
	//get ip to dial from $ ipconfig getifaddr en0

	if err != nil {
		log.Fatalln("Could not dial")
	}

	log.Printf("--- Client is connected to the server ---")

	return gRPC.NewTimeAskServiceClient(conn), err
}
