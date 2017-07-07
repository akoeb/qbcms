package model

/*
Page BO
*/
type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	Author      User   `json:"author"`
}
