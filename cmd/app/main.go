package main

import (
	//"fmt"
	"net/http"

	"github.com/Shyyw1e/effective-mobile-subs/internal/db"
	handle "github.com/Shyyw1e/effective-mobile-subs/internal/delivery/http"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/config"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"github.com/go-chi/chi/v5"

	_ "github.com/Shyyw1e/effective-mobile-subs/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Subscription API
// @version 1.0
// @description REST-сервис для работы с подписками
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Load()
	logger.InitLog(cfg.LogLevel)

	r := chi.NewRouter()
	r.Get("/docs/*", httpSwagger.WrapHandler)

	logger.Log.Infof("🚀 Starting app on port %s...", cfg.AppPort)

	database, err := db.InitDB(cfg)
	if err != nil {
		logger.Log.Fatalf("failed to init database: %v", err)
	}

	handle.RegisterRoutes(r, database)

	err = http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
	}
}
