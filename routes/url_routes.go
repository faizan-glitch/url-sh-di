package routes

import (
	"net/http"
	"url-sh/controllers"
	"url-sh/storage"
)

func SetupRoutes(m *http.ServeMux) {

	memoryStore := storage.NewMemoryStore()

	urlController := controllers.NewUrlController(memoryStore)

	m.HandleFunc("POST /shorten", urlController.ShortenUrl)

	m.HandleFunc("GET /resolve", urlController.ResolveUrl)
}
