package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// Generic interface for a recipe storage
type RecipeStorage interface {
	List() []*Recipe
	Get(string) *Recipe
	Update(string, Recipe) *Recipe
	Create(Recipe)
	Delete(string) *Recipe
}

// Storage for recipes in a folder
type RecipeFileStore struct {
	folderPath string
	metadata   map[string]string
}

// Create a new file store
func NewRecipeFileStore(config *AppConfig) *RecipeFileStore {
	return &RecipeFileStore{
		folderPath: config.FolderPath,
		metadata:   make(map[string]string),
	}
}

func (b RecipeFileStore) Get(id string) *Recipe {
	filename, exists := b.metadata[id]
	if !exists {
		return nil
	}

	filePath := b.folderPath + "/" + filename
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil
	}

	return &Recipe{
		ID:          id,
		Title:       filename,
		Filename:    filename,
		Content:     string(content),
		CreatedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
	}
}

func (b RecipeFileStore) List() []*Recipe {

	// Read the directory contents using os.ReadDir
	files, err := os.ReadDir(b.folderPath)
	if err != nil {
		// Handle the error (e.g., log it or return an empty list)
		return nil
	}

	var recipes []*Recipe

	// Iterate through each file in the directory
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Get file information (to retrieve creation/modification time)
		fileInfo, err := file.Info()
		if err != nil {
			continue // Skip files that can't provide info
		}

		// Read the file content using os.ReadFile
		content, err := os.ReadFile(b.folderPath + "/" + file.Name())
		if err != nil {
			continue // Skip files that can't be read
		}

		// Add basic tags based on filename or content (as an example)
		// TODO: add based in metadata section in the markdown
		tags := []string{}

		// Generate an ID for this file by hashing
		// the filename and store the id-filename mapping
		id := generateID(file.Name())
		b.metadata[id] = file.Name()

		// Create a new Recipe struct with file details
		recipe := &Recipe{
			ID:          id,
			Title:       file.Name(),
			Filename:    file.Name(),
			Content:     string(content),
			CreatedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			Tags:        tags,
		}

		// Add the recipe to the slice
		recipes = append(recipes, recipe)
	}

	return recipes
}

func (b RecipeFileStore) Create(recipe Recipe) {

}

func (b RecipeFileStore) Delete(id string) *Recipe {
	return nil
}

func (b RecipeFileStore) Update(id string, recipeUpdate Recipe) *Recipe {
	filename, exists := b.metadata[id]
	if !exists {
		return nil
	}

	filePath := b.folderPath + "/" + filename

	// Write the new content to the file with permissions:
	// -rw-r--r--
	err := os.WriteFile(filePath, []byte(recipeUpdate.Content), 0644)
	if err != nil {
		return nil
	}

	// Get the updated file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil
	}

	// Add basic tags based on filename or content (as an example)
	// TODO: add based in metadata section in the markdown
	tags := []string{}

	return &Recipe{
		ID:          id,
		Title:       filename,
		Filename:    filename,
		Content:     recipeUpdate.Content,
		CreatedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
		Tags:        tags,
	}
}

// Utility function to generate a unique ID
func generateID(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
