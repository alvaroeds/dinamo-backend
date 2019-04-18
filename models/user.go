package models

type User struct {
	Id              uint   `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	LasName         string `json:"las_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Numero          uint   `json:"numero"`
}
