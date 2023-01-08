package main

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
