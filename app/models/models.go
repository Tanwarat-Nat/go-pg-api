package models

type Post struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

//we are later going to be mapping our responses to JSON
//sometimes the fields from our database might not correspond to our json fields
//we want to have flexibility to add new fields or remove fields.
type JsonPost struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

//When we add a post to our database the ID is auto incremented,
//meaning we do not have to pass the ID when making a request
//to create a new post.
type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
