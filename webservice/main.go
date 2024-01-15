package main

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.URLFormat)
	cors := cors.New(cors.Options{
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials:       true,
		MaxAge:                 3599, // Maximum value not ignored by any of major browsers
	})

	router.Use(cors.Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to URL Shortener Backend Webservice Microservice"))
	})

	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", router))
}
