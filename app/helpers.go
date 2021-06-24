package app

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi03/app/models"
)

//the parse, the request body that the user entered to initialize an empty PostRequest struct
func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

//This function sends a JSON response with an http status
func SendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("helpers: SendResponse: Cannot format json. err=%v\n", err)
	}
}

//map data to JsonPost struct in case we want that flexibility
func MapPostToJSON(p *models.Post) models.JsonPost {
	return models.JsonPost{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}
