package main

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Filename     string   `json:"filename"`
	ImagePath    string   `json:"image_path,omitempty"`
	ModifiedDate string   `json:"modified_date"`
	Content      string   `json:"content"`
	Tags         []string `json:"tags,omitempty"`
}
