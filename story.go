package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func JSONStory(reader io.Reader) Story {
	decoder := json.NewDecoder(reader)
	var story Story
	err := decoder.Decode(&story)

	if err != nil {
		fmt.Println("Could not parse JSON into story. Invalid JSON")
		os.Exit(1)
	}

	return story
}

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []ChapterOption `json:"options"`
}

type ChapterOption struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}
