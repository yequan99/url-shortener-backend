package main

import (
	"fmt"
	// "net/http"

	// chi "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/rs/cors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"helpers/awsservice"
	"helpers/dynamodbops"
)

type Item struct {
	UserID string `json:"userID"`
	Hash   string `json:"hash"`
}

func ReadDB() {
	svc := awsservice.GetDBConn()

	tableName := "UserAuth"
	userID := "4"
	keyAttributes := map[string]*dynamodb.AttributeValue{
		"userID": {
			S: aws.String(userID),
		},
	}

	item := Item{}

	result, err := dynamodbops.ReadItems(svc, tableName, keyAttributes)

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.Hash == "" {
		fmt.Println("not found!")
	} else {
		fmt.Println("Found item:")
		fmt.Println("hash:  ", item.Hash)
	}
}

func WriteDB() {
	svc := awsservice.GetDBConn()

	tableName := "UserAuth"
	item := Item{
		UserID: "6",
		Hash:   "cvb",
	}

	condition := []string{"userID"}

	err := dynamodbops.InsertItems(svc, tableName, item, condition)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// router := chi.NewRouter()

	// router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.URLFormat)
	// cors := cors.New(cors.Options{
	// 	AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
	// 	AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	AllowCredentials:       true,
	// 	MaxAge:                 3599, // Maximum value not ignored by any of major browsers
	// })

	// router.Use(cors.Handler)

	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to URL Shortener Backend Authentication Microservice"))
	// })

	// fmt.Println("Starting server at port 5050")
	// fmt.Println(http.ListenAndServe(":5050", router))

	// ReadDB()
	WriteDB()
}
