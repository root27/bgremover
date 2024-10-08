package main

import (
	"context"
	"flag"
	"github.com/root27/bgremover/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {

	var file string

	flag.StringVar(&file, "file", "", "The image file to be processed")

	flag.Parse()

	//TODOS: Grpc client here and send image to python grpc server

	if file == "" {

		log.Fatal("Please provide an image file to process")

	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(

		insecure.NewCredentials(),
	))
	if err != nil {

		log.Fatal("Could not connect to the server", err)
	}

	defer conn.Close()

	client := pb.NewRemoveClient(conn)

	// Client send grpc request to python

	bytesRead, err := os.ReadFile(file)

	if err != nil {
		log.Fatal("Could not read file ", err)
	}

	res, err := client.RemoveBG(context.Background(), &pb.ImageRequest{Image: bytesRead})

	if err != nil {

		log.Fatal("Could not process image ", err)
	}

	err = os.WriteFile("animal-2-out.jpg", res.ProcessedImage, 0644)

	if err != nil {

		log.Fatal("Could not save processed image ", err)

	}
}
