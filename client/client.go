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
	id         int
	name       string
	portNumber int
}

var senderName = flag.String("sender", "default", "Senders name")
var global = 1

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

	// clients = append(clients, Client{
	// 	id:         global,
	// 	portNumber: *clientPort,
	// })

	client := &Client{
		name:       scannerName,
		portNumber: *clientPort,
	}

	//clients = append(clients, *client)

	fmt.Println(global)
	global := global + 1
	fmt.Println(global)

	go startClient(client)

	for {

	}
}

func startClient(client *Client) {
	serverConnection, err := getServerConnection()

	if err != nil {
		log.Printf("Error..")
	}

	for {
		sendMessage(client, serverConnection)
	}
}

func sendMessage(client *Client, serverConnection gRPC.TimeAskServiceClient) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		log.Printf("Client input %s\n", input)

		msg, err := serverConnection.GetTime(context.Background(), &gRPC.Message{
			Clientname: string(client.name),
			Message:    input,
		})

		if err != nil {
			log.Printf("Could not get time")
		}

		k, err := msg.Recv()

		log.Printf("Server has received message: %v from %v: ", k.Message, k.Clientname)
		// log.Printf("Server has received message: %v from %v: ", msg.Message, msg.Clientname)
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
	conn, err := grpc.Dial("172.20.10.5:5400", opts...) //smusHosÅbenBar
	// conn, err := grpc.Dial("192.168.0.110:5400", opts...) //kure
	//get ip to dial from $ ipconfig getifaddr en0

	if err != nil {
		log.Fatalln("Could not dial")
	}

	log.Printf("Dialed")

	return gRPC.NewTimeAskServiceClient(conn), err
}
