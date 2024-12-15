package main

import (
	"adsboard/internal/ads"
	"adsboard/internal/config"
	"adsboard/internal/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация БД
	store, err := storage.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	svc := ads.NewService(store)

	// Создаем маршрутизатор
	r := mux.NewRouter()

	// Регистрируем маршруты
	ads.RegisterRoutes(r, svc)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
