package models

type User struct {
	Id              uint   `json:"id"`
	Email           string `json:"email,omitempty"`
	Name            string `json:"name,omitempty"`
	LastName         string `json:"lastname,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	Numero          uint   `json:"numero,omitempty"`
}
