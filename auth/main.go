package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"auth/handler"
	"helpers/dstruct"
	"helpers/general"
)

// func ReadDB() {
// 	svc := awsservice.GetDBConn()

// 	tableName := "UserAuth"
// 	userID := "3"
// 	keyAttributes := map[string]*dynamodb.AttributeValue{
// 		"UserID": {
// 			S: aws.String(userID),
// 		},
// 	}

// 	item := models.UserAuth{}

// 	result, err := dynamodbops.ReadItems(svc, tableName, keyAttributes)
// 	if err != nil {
// 		log.Error("Failed to read DB: ", err)
// 	} else {
// 		err = dynamodbattribute.UnmarshalMap(result.Item, &item)
// 		if err != nil {
// 			log.Error("Failed to read DB %s", err)
// 		}

// 		if item.HashedPwd == "" {
// 			fmt.Println("not found!")
// 		} else {
// 			fmt.Println("Found item:", item)
// 			fmt.Println("username:  ", item.Username)
// 			fmt.Println("hash:  ", item.HashedPwd)
// 		}
// 	}
// }

// func WriteDB() {
// 	svc := awsservice.GetDBConn()

// 	tableName := "UserAuth"
// 	item := models.UserAuth{
// 		Username:  "hello",
// 		HashedPwd: "cvb",
// 	}

// 	condition := []string{"UserID"}

// 	err := dynamodbops.InsertItems(svc, tableName, item, condition)
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Error(err)
// 	}
// }

// func DeleteDB() {
// 	svc := awsservice.GetDBConn()
// 	tableName := "UserAuth"
// 	UserID := "5"
// 	keyAttributes := map[string]*dynamodb.AttributeValue{
// 		"UserID": {
// 			S: aws.String(UserID),
// 		},
// 	}
// 	err := dynamodbops.DeleteItems(svc, tableName, keyAttributes)
// 	if err != nil {
// 		fmt.Println("Failed to delete: ", err)
// 		log.Error(err)
// 	} else {
// 		fmt.Println("Deleted userID" + " (" + UserID + ") from table " + tableName)
// 	}

// }

// func UpdateDB() {
// 	svc := awsservice.GetDBConn()
// 	tableName := "UserAuth"
// 	UserID := "4"
// 	keyAttributes := map[string]*dynamodb.AttributeValue{
// 		"UserID": {
// 			S: aws.String(UserID),
// 		},
// 	}
// 	expressionAttributes := map[string]*dynamodb.AttributeValue{
// 		":HashedPwd": {
// 			S: aws.String("cvbn"),
// 		},
// 		":Salt": {
// 			S: aws.String("xcv"),
// 		},
// 		":Username": {
// 			S: aws.String("world"),
// 		},
// 	}
// 	change := []string{"HashedPwd", "Salt", "Username"}
// 	err := dynamodbops.UpdateItems(svc, tableName, keyAttributes, expressionAttributes, change)
// 	if err != nil {
// 		fmt.Println("Failed to update: ", err)
// 		log.Error(err)
// 	} else {
// 		fmt.Println("Item updated")
// 	}
// }

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
		w.Write([]byte("Welcome to URL Shortener Backend Authentication Microservice"))
	})

	router.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		var credentials dstruct.UserLoginCredentials
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&credentials)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		err = handler.Authenticate(credentials)
		if err != nil {
			response := dstruct.ErrorMessage{
				Error: err.Error(),
			}
			jsonResponse, _ := json.Marshal(response)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonResponse)
		} else {
			token, err := generateJWTToken(credentials.Username)
			if err != nil {
				http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Auth-Token", token)
			w.WriteHeader(http.StatusOK)
		}
	})

	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var credentials dstruct.UserLoginCredentials
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&credentials)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		err = handler.Register(credentials)
		if err != nil {
			response := dstruct.ErrorMessage{
				Error: err.Error(),
			}
			jsonResponse, _ := json.Marshal(response)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			w.Write(jsonResponse)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})

	fmt.Println("Starting server at port 5050")
	fmt.Println(http.ListenAndServe(":5050", router))

	// ReadDB()
	// WriteDB()
	// DeleteDB()
	// UpdateDB()

}

// GenerateJWTToken generates a new JWT token for the given username
func generateJWTToken(username string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(30 * time.Second).Unix(),
	})

	// Sign the token with the secret key
	signingKey, err := general.GetJWTSecret("JWT_TOKEN")
	if err != nil {
		log.Errorf("Error 1: %s", err)
		return "", err
	}
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Errorf("Error 2: %s", err)
		return "", err
	}

	return tokenString, nil
}
