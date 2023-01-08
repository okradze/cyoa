package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)



func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "JSON file with the story")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("Could not read file: %s\n", *filename)
		os.Exit(1)
	}

	defer file.Close()

	story := JSONStory(file)

	http.HandleFunc("/", NewHandler(story))

	fmt.Printf("Starting the server on port: %d", *port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
}