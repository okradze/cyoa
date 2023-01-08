package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"okradze/cyoa/templates"
	"os"
	"strings"
	"text/template"
)


func NewHandler(story Story) http.HandlerFunc {
	tpl := template.Must(template.New("").Parse(templates.DefaultTemplate))

	return func (w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSpace(r.URL.Path)

		if path == "" || path == "/" {
			path = "/intro"
		}

		path = path[1:]

		chapter, ok := story[path]

		if !ok {
			http.Error(w, "Chapter not found", http.StatusNotFound)
			return
		}

		err := tpl.Execute(w, chapter)

		if err != nil {
			fmt.Printf("Could not execute template: %s", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}

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
