package main

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Filename     string   `json:"filename"`
	ModifiedDate string   `json:"modified_date"`
	Content      string   `json:"content"`
	Tags         []string `json:"tags,omitempty"`
}
