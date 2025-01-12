
package main

import (
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

var fakeRecipes = []*Recipe{{
    Title:            "Avocados",
    CreatedDate:    "24/12/2024",
    Content: "Create avocados",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *Recipe {
    return fakeRecipes[0]
}

func (s fakeStorage) Delete(_ string) *Recipe {
    return nil
}  

func (s fakeStorage) List() []*Recipe {
    return fakeRecipes
}

func (s fakeStorage) Create(_ Recipe) {
    return
}

func (s fakeStorage) Update(string, Recipe) *Recipe {
    return fakeRecipes[1]
}

// TODO: implement mocked service layer
func TestGetRecipesController(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/recipes/1", nil)
    w := httptest.NewRecorder()
    service := &RecipeService{storage: fakeStorage{}}
    recipeController := &RecipeController{service: service}
    recipeController.GetRecipe(w, req)
    res := w.Result()
    defer res.Body.Close()
    data, err := io.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    recipe := Recipe{}
    json.Unmarshal(data, &recipe)
    if recipe.Title != "Avocados" {
        t.Errorf("expected ABC got %v", string(data))
    }
}