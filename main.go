package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func initServer(router *chi.Mux) {
	port := os.Getenv("PORT")
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	slog.Info(fmt.Sprintln("Staring server on port:", port))
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Fatal server couldnt start port:")
	}
}

func main() {
	godotenv.Load()
	apiConf := handlers.CreateApiConfig()
	if apiConf == nil {
		slog.Error("CreateApiConfig returned nil")
		return
	}

	router := chi.NewRouter()
	v1router := chi.NewRouter()
	router.Use(handlers.LoggingMiddleware())

	v1router.Get("/products", apiConf.HandleGetProducts)
	v1router.Post("/products", handlers.ParseJSONAndValidateMiddleware[database.CreateProductParams](apiConf.HandleCreateProduct))
	v1router.Put("/products/{productId}", handlers.ParseJSONAndValidateMiddleware[database.UpdateProductParams](apiConf.HandleUpdateProduct))

	v1router.Get("/reports", apiConf.HandleGetReports)
	v1router.Post("/reports", apiConf.HandleCreateReport)
	router.Mount("/api/v1", v1router)

	initServer(router)

}
