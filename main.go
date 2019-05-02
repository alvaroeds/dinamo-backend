package main

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func main(){
        e := echo.New()

        e.Static("/", "public")

        e.POST("/login", controllers.LoginUser)
        e.POST("/register", controllers.CreateUser)
        e.GET("/validate", controllers.ValidateToken)
        e.POST("/upload", controllers.Upload)

        e.Start(":2020")
}
