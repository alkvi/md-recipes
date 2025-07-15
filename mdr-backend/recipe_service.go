package main

import (
	"strings"
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type RecipeService struct {
	storage RecipeStorage
	logger *logrus.Logger
}

func (s *RecipeService) enrichTags(recipes []*Recipe) {
	for _, r := range recipes {
		s.logger.Debugf("Parsing tags for recipe %s", r.Title)
		tags := parseFrontMatter(r.Content, s.logger)
		if tags != nil {
			r.Tags = tags
		}
	}
}

func (s *RecipeService) ListRecipes() []*Recipe {
	recipes := s.storage.List()
	s.logger.Debugf("Found %d recipes", len(recipes))
	s.enrichTags(recipes)
	return recipes
}

func (s *RecipeService) GetRecipe(id string) *Recipe {
	recipe := s.storage.Get(id)
	if recipe != nil {
		s.enrichTags([]*Recipe{recipe})
	}
	return recipe
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
	s.enrichTags(recipes)
	for _, recipe := range recipes {
		if recipe.Filename == filename {
			return recipe
		}
	}
	return nil
}

func (s *RecipeService) SearchRecipes(filters map[string]string) []*Recipe {
	allRecipes := s.storage.List()
	s.enrichTags(allRecipes)
	var results []*Recipe

	for _, recipe := range allRecipes {
		if matchesFilters(recipe, filters) {
			results = append(results, recipe)
		}
	}
	s.logger.Debugf("Found %d matches", len(results))
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

func parseFrontMatter(content string, logger *logrus.Logger) []string {
	content = strings.TrimSpace(content)

	if !strings.HasPrefix(content, "---") {
		logger.Debug("Markdown has no prefix")
		return nil
	}

	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		logger.Debug("Markdown cannot be split on --- into 3")
		return nil
	}

	rawYaml := parts[1]

	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(rawYaml), &metadata); err != nil {
		return nil
	}

	tags := []string{}
	for k, v := range metadata {
		tag := fmt.Sprintf("%s:%v", k, v)
		logger.Debugf("Found tag: %s", tag)
		tags = append(tags, tag)
	}
	return tags
}
