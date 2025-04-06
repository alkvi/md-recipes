package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// Generic interface for a recipe storage
type RecipeStorage interface {
	List() []*Recipe
	Get(string) *Recipe
	Update(string, Recipe) (Recipe, error)
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
		ID:           id,
		Title:        filename,
		Filename:     filename,
		Content:      string(content),
		ModifiedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
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
			ID:           id,
			Title:        file.Name(),
			Filename:     file.Name(),
			Content:      string(content),
			ModifiedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			Tags:         tags,
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

func (b RecipeFileStore) Update(id string, recipe Recipe) (Recipe, error) {
	// Get the current filename
	oldFilename, exists := b.metadata[id]
	if !exists {
		return Recipe{}, fmt.Errorf("recipe not found")
	}

	// Generate new filename based on title
	newFilename := sanitizeFilename(recipe.Title) + ".md"
	newPath := filepath.Join(b.folderPath, newFilename)

	// If the title has changed, we need to rename the file
	if oldFilename != newFilename {
		oldPath := filepath.Join(b.folderPath, oldFilename)

		// Check if new filename already exists
		if _, err := os.Stat(newPath); err == nil {
			return Recipe{}, fmt.Errorf("a recipe with this title already exists")
		}

		// Rename the file
		if err := os.Rename(oldPath, newPath); err != nil {
			return Recipe{}, fmt.Errorf("failed to rename recipe file: %w", err)
		}

		// Update the filename mapping
		b.metadata[id] = newFilename
	}

	// Write the new content to the file
	// File permissions: owner can read and write (6), group and others can only read (4)
	// 0644 = -rw-r--r--
	if err := os.WriteFile(newPath, []byte(recipe.Content), 0644); err != nil {
		return Recipe{}, fmt.Errorf("failed to write recipe file: %w", err)
	}

	// Get the file info to update the modification time
	info, err := os.Stat(newPath)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to get file info: %w", err)
	}

	// Update the recipe with the new modification time
	recipe.ModifiedDate = info.ModTime().Format(time.RFC3339)
	return recipe, nil
}

// Utility function to generate a unique ID
func generateID(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

func sanitizeFilename(title string) string {
	// Replace spaces with hyphens and remove invalid characters
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return '-'
		}
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-' && r != '_' {
			return -1
		}
		return r
	}, strings.ToLower(title))
}
