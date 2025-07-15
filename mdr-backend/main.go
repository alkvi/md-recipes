package main

import (
	"os"
	"net/http"
	"strings"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sirupsen/logrus"
)

func main() {

	// Load config options
	config, err := LoadConfig("config.yaml")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Create logger
	log := SetupLogger(config)
	log.Debug("Logger initiated")

	// Create router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/config", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(config)
    })
	r.Mount("/recipes", RecipeRoutes(config, log))
	http.ListenAndServe(":3000", r)
}

// Log specific setup
func SetupLogger(config *AppConfig) *logrus.Logger {
    logger := logrus.New()
    logger.SetOutput(os.Stdout)

    // Use colored, timestamped text formatter
    logger.SetFormatter(&logrus.TextFormatter{
        ForceColors:     true,
        FullTimestamp:   true,
        TimestampFormat: "2006-01-02 15:04:05",
    })

    // Parse log level from config (case-insensitive)
    level, err := logrus.ParseLevel(strings.ToLower(config.LogLevel))
    if err != nil {
        logger.Warnf("Invalid log level '%s', defaulting to 'info'", config.LogLevel)
        level = logrus.InfoLevel
    }

    logger.SetLevel(level)
    return logger
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

func RecipeRoutes(config *AppConfig, logger *logrus.Logger) chi.Router {

	folderPath := config.FolderPath
	logger.Infof("Using folder path: %s", folderPath)

	storage := NewRecipeFileStore(config, logger)
	service := &RecipeService{storage: storage, logger: logger}
	controller := &RecipeController{service: service, logger: logger}

	r := chi.NewRouter()
	r.Get("/", controller.ListRecipes)
	r.Post("/", controller.CreateRecipe)
	r.Get("/{id}", controller.GetRecipe)
	r.Put("/{id}", controller.UpdateRecipe)
	r.Delete("/{id}", controller.DeleteRecipe)
	r.Get("/search", controller.SearchRecipes)
	return r
}
