package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/root27/bgremover/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strings"
)

const welcome = `

 _                                                  
| |__   __ _ _ __ ___ _ __ ___   _____   _____ _ __ 
| '_ \ / _  | '__/ _ \ '_   _ \ / _ \ \ / / _ \ '__|
| |_) | (_| | | |  __/ | | | | | (_) \ V /  __/ |   
|_.__/ \__, |_|  \___|_| |_| |_|\___/ \_/ \___|_|   
       |___/                                        

bgremover-cli is a tool to remove background from images

`

func checkExtension(file string) string {

	elems := strings.Split(file, "/")

	ext := strings.Split(elems[len(elems)-1], ".")

	return ext[len(ext)-1]

}

func main() {

	fmt.Println(welcome)

	var file string

	var outfile string

	flag.StringVar(&file, "file", "", "The image file to be processed")

	flag.StringVar(&outfile, "out", "", "The output file to save the processed image")

	flag.Parse()

	if file == "" {

		fmt.Println("Please provide an image file to process")
		return
	}

	fileExt := checkExtension(file)

	if outfile == "" {

		outfile = "processed_." + fileExt

	}

	fmt.Println("out file: ", outfile)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(

		insecure.NewCredentials(),
	))
	if err != nil {

		fmt.Println("Could not connect to the server")
		return
	}

	defer conn.Close()

	client := pb.NewRemoveClient(conn)

	bytesRead, err := os.ReadFile(file)

	if err != nil {

		fmt.Println("Could not read image file")
		return
	}

	res, err := client.RemoveBG(context.Background(), &pb.ImageRequest{Image: bytesRead})

	if err != nil {

		fmt.Println("Could not process image")
		return
	}

	err = os.WriteFile(outfile, res.ProcessedImage, 0644)

	if err != nil {

		fmt.Println("Could not save processed image")
		return
	}
}
