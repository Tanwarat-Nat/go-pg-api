package database

import (
	"fmt"
	"log"
	"restapi03/app/models"
)

//method to interact with database
func (d *DB) CreatePost(p *models.Post) error {
	res, err := d.db.Exec(insertPostSchema, p.Title, p.Content, p.Author)
	if err != nil {
		log.Println("methods: CreatePost: ", err)
	}
	res.LastInsertId()
	return err
}

//method to interact with database
func (d *DB) GetPost() ([]*models.Post, error) {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return posts, fmt.Errorf("methods: Getpost: err=%v ", err)
	}
	return posts, nil
}
