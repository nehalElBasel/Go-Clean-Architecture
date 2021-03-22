package models

type Post struct {
	ID     int    `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
}
