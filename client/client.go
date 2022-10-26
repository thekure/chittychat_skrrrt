package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strconv"

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
	conn, err := grpc.Dial(":"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Could not dial")
	}

	log.Printf("Dialed")

	return gRPC.NewTimeAskServiceClient(conn), err
}
