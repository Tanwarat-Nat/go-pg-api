package database

import (
	"log"
	"restapi03/app/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//it's like a contract (methods to implement) that we have to respect
//if we want our database to be a PostDB database. Let's start by saying that
//a PostDB should implement a Open() and a Close method that can return an error.
type PostDB interface {
	Open() error
	Close() error
	CreatePost(p *models.Post) error
	GetPost() ([]*models.Post, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Conected to database.")
	pg.MustExec(createSchemas)
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
