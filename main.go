package main

import (
	"flag"
	"fmt"
	"os"
)



func main() {
	filename := flag.String("file", "gopher.json", "JSON file with the story")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("Could not read file: %s\n", *filename)
		os.Exit(1)
	}

	defer file.Close()

	story := JSONStory(file)

	_, ok := story["intro"]

	if !ok {
		fmt.Println("There is no intro chapter in story")
		os.Exit(1)
	}

	fmt.Printf("%+v\n", story)
}