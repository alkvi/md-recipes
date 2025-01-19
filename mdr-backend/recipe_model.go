package main

type Recipe struct {
    ID               string `json:"id"`
    Title            string `json:"title"`
    Filename         string `json:"filename"`
    CreatedDate      string `json:"created_date"`
    Content          string `json:"content"`
    Tags             []string `json:"tags,omitempty"`
}