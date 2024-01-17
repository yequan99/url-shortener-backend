package handler

import (
	"fmt"

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
	shortURL := fmt.Sprintf("http://127.0.0.1:8080/%s", shortID)
	fmt.Printf("Shortened URL: %s\n", shortURL)

	// Store Short URL into DB

	return shortURL, nil
}
