package routes

import (
	"net/http"

	"github.com/ZuhybDev/waqt/internal/handlers"
)

func Register() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/prayers", handlers.HandlePrayers)
	mux.HandleFunc("/api/quran", handlers.HandlePrayers)

	return mux
}
