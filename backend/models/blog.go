package models

type Blog struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorId int    `json:"author_id"`
}
