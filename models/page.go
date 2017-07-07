package model

/*
Page BO
*/
type Page struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Template Template `json:"template"`
	Posts    []Post   `json:"posts"`
}
