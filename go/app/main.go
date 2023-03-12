package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	tRouter := chi.NewRouter()
	tRouter.Use(middleware.Logger)
	tRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"))
	})
	http.ListenAndServe(":8080", tRouter)
}
