package main

import (
	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/handlers"
	"github.com/coby-amar/go_learning/main/middleware"
	"github.com/coby-amar/go_learning/main/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func setupRouter(config *utils.ApiConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://10.100.102.41:8080", "http://10.100.102.41:8080", "http://localhost:2318"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}))
	router.Route("/api/v1", func(r chi.Router) {
		v1Router(r, config)
	})
	return router
}

func v1Router(router chi.Router, config *utils.ApiConfig) {
	router.Use(middleware.AuthenticationMiddleware(config))

	router.Get("/user", middleware.ConfigInjectorMiddleware(config, handlers.HandleGetUser))

	router.Get("/products", middleware.ConfigInjectorMiddleware(config, handlers.HandleGetProducts))
	router.Delete("/products/{productId}", middleware.ConfigInjectorMiddleware(config, handlers.HandleDeleteProduct))
	router.Post("/products",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[database.CreateProductParams](handlers.HandleCreateProduct),
		),
	)
	router.Put("/products",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[database.UpdateProductParams](handlers.HandleUpdateProduct),
		),
	)

	router.Get("/reports", middleware.ConfigInjectorMiddleware(config, handlers.HandleGetReports))
	router.Get("/reports/{reportId}/entries", middleware.ConfigInjectorMiddleware(config, handlers.HandleGetReportEntries))
	router.Delete("/reports/{reportId}", middleware.ConfigInjectorMiddleware(config, handlers.HandleDeleteReport))
	router.Put("/reports",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[utils.UserUpdateReportWithEntries](handlers.HandleUpdateReport),
		),
	)
	router.Post("/reports",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[utils.UserCreateReportWithEntries](handlers.HandleCreateReport),
		),
	)

	router.Route("/auth", func(r chi.Router) {
		authRouter(r, config)
	})
}

func authRouter(router chi.Router, config *utils.ApiConfig) {
	router.Get("/healthz", middleware.ConfigInjectorMiddleware(config, handlers.HandleHealthZ))

	router.Post("/logout", middleware.ConfigInjectorMiddleware(config, handlers.HandleLogout))

	router.Post("/register",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[utils.RegistrationJson](handlers.HandleRegister),
		),
	)
	router.Post("/login",
		middleware.ConfigInjectorMiddleware(
			config,
			middleware.ParseJSONAndValidateMiddleware[utils.LoginJson](handlers.HandleLogin),
		),
	)
}
