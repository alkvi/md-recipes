package main

import (
    "encoding/json"
    "net/http"
    "fmt"

    "github.com/go-chi/chi/v5"
    "github.com/sirupsen/logrus"
)

type RecipeController struct {
    service *RecipeService
    logger *logrus.Logger
}

func (h *RecipeController) ListRecipes(w http.ResponseWriter, r *http.Request) {
    recipes := h.service.ListRecipes()
    err := json.NewEncoder(w).Encode(recipes)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}

func (h *RecipeController) GetRecipe(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    recipe := h.service.GetRecipe(id)
    if recipe == nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }
    err := json.NewEncoder(w).Encode(recipe)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}

func (h *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
    var recipe Recipe
    err := json.NewDecoder(r.Body).Decode(&recipe)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createdRecipe := h.service.CreateRecipe(recipe)
    err = json.NewEncoder(w).Encode(createdRecipe)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}

func (h *RecipeController) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var recipe Recipe
    err := json.NewDecoder(r.Body).Decode(&recipe)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updatedRecipe := h.service.UpdateRecipe(id, recipe)
    if updatedRecipe == nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }
    err = json.NewEncoder(w).Encode(updatedRecipe)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}

func (h *RecipeController) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    deletedRecipe := h.service.DeleteRecipe(id)
    if deletedRecipe == nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *RecipeController) GetRecipeByFilename(w http.ResponseWriter, r *http.Request) {
    filename := chi.URLParam(r, "filename")
    recipe := h.service.GetRecipeByFilename(filename)
    if recipe == nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }
    err := json.NewEncoder(w).Encode(recipe)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}

func (h *RecipeController) SearchRecipes(w http.ResponseWriter, r *http.Request) {

    // Currently allowed queries
    allowedParams := map[string]bool{
        "title":    true,
        "author":   true,
        "category": true,
        "tag":      true,
        "id":       true,
    }

    // Parse the query param ?
    query := r.URL.Query()

    // Prepare a map query:value
    validParams := make(map[string]string)

    // Validate parameters and put into map
    for key, values := range query {
        if !allowedParams[key] {
            http.Error(w, fmt.Sprintf("Invalid query parameter: %s", key), http.StatusBadRequest)
            return
        }

        // Only consider the first value for each key
        validParams[key] = values[0]

        // Log info message if multiple values
        if len(values) > 1 {
            h.logger.Info("Only first value for query parameter read")
        }
    }

    if len(validParams) == 0 {
        http.Error(w, "At least one allowed query parameter is required (title, author, category, tag, id)", http.StatusBadRequest)
        return
    }

    // Log valid parameters
    for k, v := range validParams {
        h.logger.Debugf("Searching with %s: %s", k, v)
    }

    recipes := h.service.SearchRecipes(validParams)
    err := json.NewEncoder(w).Encode(recipes)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }

}