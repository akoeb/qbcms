package model

/*
User the user BO
*/
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
