package app

import (
	"fmt"
	"log"
	"net/http"
	"restapi03/app/models"
)

//This is where we are going to store the handlers for our routes.
//This will return an http.HandlerFunc printing a response "Welcome to Rest API".
func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Rest API")
	}
}

//this function will parse the request body and use that data
//to create a new post and save it in our database.
func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := Parse(w, r, &req)
		if err != nil {
			log.Printf("handlers: CreatePostHandler: Cannot pares post body. err=%v\n", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// create th post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		//save in DB
		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("handlers: CreatePostHandler: Cannot save post in DB. err = %v\n", err)
			SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := MapPostToJSON(p)
		SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPost()
		if err != nil {
			log.Printf("jhandlers: GetPostHandler: Cannot get post, err=%v\n", err)
			SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = MapPostToJSON(post)
		}
		SendResponse(w, r, resp, http.StatusOK)
	}
}
