package main

import (
	"log"
	"net/http"
	"restapi03/app"
	"restapi03/app/database"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting the services.")

	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Open()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("The services is ready to listen and serve.")
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

//This function to print help with error handling.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
