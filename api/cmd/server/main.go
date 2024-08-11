package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/manururu/manuru-ru/api/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.Home)
}
