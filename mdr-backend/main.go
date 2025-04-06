package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Load config options
	config, err := LoadConfig("config.yaml")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Create router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/recipes", RecipeRoutes(config))
	http.ListenAndServe(":3000", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RecipeRoutes(config *AppConfig) chi.Router {

	folderPath := config.FolderPath
	fmt.Println("Using folder path:", folderPath)

	storage := NewRecipeFileStore(config)
	service := &RecipeService{storage: storage}
	controller := &RecipeController{service: service}

	r := chi.NewRouter()
	r.Get("/", controller.ListRecipes)
	r.Post("/", controller.CreateRecipe)
	r.Get("/{id}", controller.GetRecipe)
	r.Put("/{id}", controller.UpdateRecipe)
	r.Delete("/{id}", controller.DeleteRecipe)
	return r
}
