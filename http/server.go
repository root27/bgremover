package main

import (
	"bytes"
	"context"
	"github.com/root27/bgremover/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {

		log.Fatal("Could not connect to the server", err)

	}

	defer conn.Close()

	client := pb.NewRemoveClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		image, _, err := r.FormFile("image")

		if err != nil {

			http.Error(w, "Could not read image", http.StatusBadRequest)

			return

		}

		bytesRead, err := io.ReadAll(image)

		if err != nil {

			http.Error(w, "Could not read image", http.StatusBadRequest)
			return
		}

		response, err := client.RemoveBG(context.Background(), &pb.ImageRequest{Image: bytesRead})

		if err != nil {

			http.Error(w, "Error processing Image", http.StatusInternalServerError)
			return
		}

		io.Copy(w, bytes.NewReader(response.ProcessedImage))

	})

	log.Println("Server starting at port 10000")

	log.Fatal(http.ListenAndServe(":10000", nil))

}
