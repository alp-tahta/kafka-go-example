package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
}
