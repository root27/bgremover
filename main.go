package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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

P.S: File can not be larger than 10MB


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

	buf := &bytes.Buffer{}

	mWriter := multipart.NewWriter(buf)

	f, err := os.Open(file)

	fileStats, _ := f.Stat()

	if fileStats.Size() > 10*1024*1024 {

		fmt.Println("File is too large")
		return
	}

	if err != nil {

		fmt.Println("Error reading file")
		return
	}

	defer f.Close()

	fileWriter, err := mWriter.CreateFormFile("image", file)

	if err != nil {
		fmt.Println("Error creating form file")
		return
	}

	_, err = io.Copy(fileWriter, f)

	if err != nil {
		fmt.Println("Error writing file")
		return
	}

	err = mWriter.Close()

	if err != nil {
		fmt.Println("Error closing writer")
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api-bgremover/root27.dev/api/bgremove", buf)

	if err != nil {

		fmt.Println("Error creating request")
		return
	}

	req.Header.Add("Content-Type", mWriter.FormDataContentType())

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error processing request")
		return
	}

	defer response.Body.Close()

	outData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response")
		return
	}

	err = os.WriteFile(outfile, outData, 0644)

	if err != nil {

		fmt.Println("Could not save processed image")
		return
	}

	outStats, _ := os.Stat(outfile)

	if outStats.Size() == 0 {
		fmt.Println("Could not process image")
		return
	}

	fmt.Println("out file: ", outfile)

}
