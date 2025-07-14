package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type RecipeService struct {
	storage RecipeStorage
	logger *logrus.Logger
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

func (s *RecipeService) SearchRecipes(filters map[string]string) []*Recipe {
	allRecipes := s.storage.List()
	var results []*Recipe

	for _, recipe := range allRecipes {
		if matchesFilters(recipe, filters) {
			results = append(results, recipe)
		}
	}
	return results
}

func matchesFilters(recipe *Recipe, filters map[string]string) bool {
	for key, value := range filters {
		switch key {
		case "title":
			if !strings.Contains(strings.ToLower(recipe.Title), strings.ToLower(value)) {
				return false
			}
		case "tag":
			if !containsIgnoreCase(recipe.Tags, value) {
				return false
			}
		case "id":
			if recipe.ID != value {
				return false
			}
		}
	}
	return true
}

func contains(slice []string, term string) bool {
	for _, item := range slice {
		if strings.Contains(item, term) {
			return true
		}
	}
	return false
}

func containsIgnoreCase(slice []string, term string) bool {
	term = strings.ToLower(term)
	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), term) {
			return true
		}
	}
	return false
}
