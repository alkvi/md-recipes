package main

import (
    "os"
)

type RecipeStorage interface {
	List() []*Recipe
	Get(string)*Recipe
	Update(string, Recipe) *Recipe
	Create(Recipe)
	Delete(string)*Recipe
}

type RecipeFileStore struct {
    folderPath string
}

func NewRecipeFileStore(config *AppConfig) *RecipeFileStore {
    return &RecipeFileStore{
        folderPath: config.FolderPath,
    }
}

func (b RecipeFileStore) Get(id string) *Recipe {
    return nil
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

        // Create a new Recipe struct with file details
        recipe := &Recipe{
            Title:       file.Name(),
            CreatedDate: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
            Content:     string(content),
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
    return nil
}

