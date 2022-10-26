package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gRPC "github.com/thekure/chittychat_skrrrt/proto"
)

type Client struct {
	id         int
	portNumber int
}

var (
	clientPort = flag.Int("clientPort", 8081, "client port number")
	serverPort = flag.Int("serverPort", 8080, "server port number")
)

// go run server/server.go -port=8083

func main() {
	flag.Parse()

	client := &Client{
		id:         1,
		portNumber: *clientPort,
	}

	go startClient(client)

	for {

	}
}

func startClient(client *Client) {
	serverConnection, err := getServerConnection()

	if err != nil {
		log.Printf("Error..")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		log.Printf("Client input %s\n", input)

		timeMessage, err := serverConnection.GetTime(context.Background(), &gRPC.AskForTimeMessage{ClientId: int64(client.id)})

		if err != nil {
			log.Printf("Could not get time")
		}

		log.Printf("Server %s says that the time is %s\n", timeMessage.Time)
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
	conn, err := grpc.Dial("192.168.0.145:5400", opts...)

	if err != nil {
		log.Fatalln("Could not dial")
	}

	log.Printf("Dialed")

	return gRPC.NewTimeAskServiceClient(conn), err
}
