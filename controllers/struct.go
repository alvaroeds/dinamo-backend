package controllers

type User struct {
        Email           string `json:"email,omitempty"`
        PhoneNumber     string `json:"phone_number"`
        Password        string `json:"password,omitempty"`
        ConfirmPassword string `json:"confirmpassword,omitempty"`

}

type Error struct {
        Message string `json:"string"`
        ErrorCode string `json:"error_code"`
        Error string `json:"error"`
}
var msg Error
