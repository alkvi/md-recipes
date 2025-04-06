package main

import (
	"strings"
)

type RecipeService struct {
	storage RecipeStorage
}

func (s *RecipeService) ListRecipes() []*Recipe {
	return s.storage.List()
}

func (s *RecipeService) GetRecipe(id string) *Recipe {
	return s.storage.Get(id)
}

func (s *RecipeService) CreateRecipe(recipe Recipe) Recipe {
	s.storage.Create(recipe)
	return recipe
}

func (s *RecipeService) UpdateRecipe(id string, recipe Recipe) *Recipe {
	updatedRecipe, err := s.storage.Update(id, recipe)
	if err != nil {
		return nil
	}
	return &updatedRecipe
}

func (s *RecipeService) DeleteRecipe(id string) *Recipe {
	return s.storage.Delete(id)
}

func (s *RecipeService) GetRecipeByFilename(filename string) *Recipe {
	recipes := s.storage.List()
	for _, recipe := range recipes {
		if recipe.Filename == filename {
			return recipe
		}
	}
	return nil
}

func (s *RecipeService) SearchRecipes(query string) []*Recipe {
	allRecipes := s.storage.List()
	var results []*Recipe
	for _, recipe := range allRecipes {
		if strings.Contains(recipe.Title, query) || strings.Contains(recipe.Content, query) || contains(recipe.Tags, query) {
			results = append(results, recipe)
		}
	}
	return results
}

func contains(slice []string, term string) bool {
	for _, item := range slice {
		if strings.Contains(item, term) {
			return true
		}
	}
	return false
}
