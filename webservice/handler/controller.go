package handler

import (
	"fmt"

	"helpers/awsservice"
	"helpers/dynamodbops"
	"helpers/models"

	log "github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

func GetShortURL(username string, longURL string) (string, error) {

	// Generate shortURL id
	generator := shortid.MustNew(1, shortid.DefaultABC, 2342)
	shortID, err := generator.Generate()
	if err != nil {
		log.Errorf("[ID generator] Unable to generate a short ID")
	}
	shortURL := "http://127.0.0.1:8080/" + shortID

	// Store Short URL into DB
	svc := awsservice.GetDBConn()

	tableName := "UrlCode"
	item := models.UrlCode{
		Username: username,
		ShortURL: shortURL,
		LongURL:  longURL,
	}

	err = dynamodbops.InsertItems(svc, tableName, item, []string{})
	if err != nil {
		log.Errorf("[Short URL Generation] Unable to create new URL: %s", err)
		return "", fmt.Errorf("Unable to create new URL")
	}

	return shortURL, nil
}
