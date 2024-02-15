package handler

import (
	"encoding/json"
	"net/http"

	"helpers/dstruct"

	chi "github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func GenerateShortenedURLs(w http.ResponseWriter, r *http.Request) {
	var longUrlInfo dstruct.GenerateShortURL
	userID := chi.URLParam(r, "userId")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&longUrlInfo)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	shortURL, err := GenShortURL(userID, longUrlInfo.LongURL)
	if err != nil {
		log.Errorf("Error generating shortURL: ", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response := dstruct.ReturnShortURL{
			ShortURL: shortURL,
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func GetShortenedURLs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")

	shortUrlArray, err := GetStoredUrls(userID)
	if err != nil {
		log.Errorf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		jsonResponse, _ := json.Marshal(shortUrlArray)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func DeleteUserURL(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "urlId")

	err := DeleteURL(urlId)
	if err != nil {
		log.Errorf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
