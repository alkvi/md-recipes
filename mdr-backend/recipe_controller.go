package main

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
)

type RecipeController struct {
    service *RecipeService
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
    query := r.URL.Query().Get("query")
    if query == "" {
        http.Error(w, "Query parameter is required", http.StatusBadRequest)
        return
    }
    recipes := h.service.SearchRecipes(query)
    err := json.NewEncoder(w).Encode(recipes)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
    }
}