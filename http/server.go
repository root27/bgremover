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

	conn, err := grpc.Dial("bgremover-grpc-server:50051", grpc.WithTransportCredentials(
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

		ctx, cancel := context.WithTimeout(context.Background(), 30)

		defer cancel()

		stream, err := client.RemoveBG(ctx)

		if err != nil {

			http.Error(w, "Error streaming", http.StatusInternalServerError)
			return
		}

		buf := make([]byte, 1024*1024)

		reader := bytes.NewReader(bytesRead)

		for {

			n, err := reader.Read(buf)

			if err == io.EOF {
				break
			}

			if err != nil {

				http.Error(w, "Error reading image", http.StatusInternalServerError)
				return
			}

			if err := stream.Send(&pb.ImageRequest{Image: buf[:n]}); err != nil {
				http.Error(w, "Error sending image", http.StatusInternalServerError)
				return
			}

		}

		response, err := stream.CloseAndRecv()

		if err != nil {

			http.Error(w, "Error receiving response", http.StatusInternalServerError)
			return
		}

		io.Copy(w, bytes.NewReader(response.ProcessedImage))

	})

	log.Println("Server starting at port: ", port)

	log.Fatal(http.ListenAndServe(port, nil))

}
