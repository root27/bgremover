package main

import (
	"flag"
	"log"
)

func main() {

	var file string

	flag.StringVar(&file, "file", "", "The image file to be processed")

	flag.Parse()

	//TODOS: Grpc client here and send image to python grpc server

	if file == "" {

		log.Fatal("Please provide an image file to process")

	}

}
