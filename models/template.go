package model

/*
Template BO
*/
type Template struct {
	ID      int    `json:"id"`
	Path    string `json:"path"`
	Content string `json:"content"`
}
