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

const (
	port = ":10000"
)

func main() {

	conn, err := grpc.Dial("bgremover-grpc-server", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {

		log.Println("Could not connect to the server", err)

		return
	}

	log.Println("Connected to grpc server")

	defer conn.Close()

	client := pb.NewRemoveClient(conn)

	http.HandleFunc("/api/bgremove", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Content-Type", "image/*")

		err := r.ParseMultipartForm(10 << 20)

		if err != nil {

			http.Error(w, "File is too large", http.StatusBadRequest)
			return
		}

		image, _, err := r.FormFile("image")

		if err != nil {

			http.Error(w, "Could not read image", http.StatusBadRequest)
			return

		}

		defer image.Close()

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

	log.Println("Server starting at port: ", port)

	log.Fatal(http.ListenAndServe(port, nil))

}
