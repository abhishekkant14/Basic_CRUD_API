package models

type User struct {
	ID    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Phone int    `json:"Phone"`
}
