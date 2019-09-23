package controllers

type Error struct {
        Message string `json:"message"`
        ErrorCode string `json:"error_code"`
        Error string `json:"error"`
}
var msg Error
