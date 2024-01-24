package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	// "helpers/dstruct"
	"helpers/general"
	"webservice/handler"

	jwt "github.com/dgrijalva/jwt-go"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.URLFormat)
	cors := cors.New(cors.Options{
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedHeaders:         []string{"Accept", "X-Auth-Token", "Content-Type", "X-CSRF-Token"},
		AllowCredentials:       true,
		MaxAge:                 3599, // Maximum value not ignored by any of major browsers
		// OptionsPassthrough:     true, // Allow preflight requests to pass through
	})

	router.Use(cors.Handler)

	// Add JWT token authentication middleware
	// router.Use(authMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to URL Shortener Backend Webservice Microservice"))
		fmt.Println("/")
	})

	router.Post("/{userId}/shorten-url", handler.GenerateShortenedURLs)
	router.Get("/{userId}/shorten-url", handler.GetShortenedURLs)

	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", router))
}

// JWT token authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Auth-Token")
		if tokenString == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			jwtSecret, err := general.GetJWTSecret("JWT_TOKEN")
			if err != nil {
				return jwtSecret, err
			}
			return jwtSecret, nil
		})

		if err != nil {
			log.Errorf("[Auth Middleware] Error parsing token: %s", err)
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Error("[Auth Middleware] Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
