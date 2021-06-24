package app

import (
	"restapi03/app/database"

	"github.com/gorilla/mux"
)

// App struct will represent the structure of our app.
type App struct {
	Router *mux.Router
	DB     database.PostDB
}

//This function will be in charge of returning the actual application based on our struct.
func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.InitRoutes()
	return a
}

//This function will be a receiver of our App and called inside of our New() func
func (a *App) InitRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostHandler()).Methods("GET")
}
