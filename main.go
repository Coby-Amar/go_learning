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
	router.Use(handlers.LoggingMiddleware())
	router.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	v1Router := chi.NewRouter()
	router.Mount("/api/v1", v1Router)
	v1Router.Use(apiConf.AuthenticationMiddleware)

	v1AuthRouter := chi.NewRouter()
	v1Router.Mount("/auth", v1AuthRouter)

	v1AuthRouter.Post("/register", apiConf.HandleRegister)
	v1AuthRouter.Post("/login", apiConf.HandleLogin)
	v1AuthRouter.Post("/logout", apiConf.HandleLogout)

	v1Router.Get("/healthz", apiConf.HandleHealthZ)

	v1Router.Get("/products", apiConf.HandleGetProducts)
	v1Router.Post("/products", handlers.ParseJSONAndValidateMiddleware[database.CreateProductParams](apiConf.HandleCreateProduct))
	v1Router.Put("/products/{productId}", handlers.ParseJSONAndValidateMiddleware[database.UpdateProductParams](apiConf.HandleUpdateProduct))

	v1Router.Get("/reports", apiConf.HandleGetReports)
	v1Router.Post("/reports", handlers.ParseJSONAndValidateMiddleware[handlers.UserCreateReportWithEntries](apiConf.HandleCreateReport))

	initServer(router)
}
