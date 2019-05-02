package main

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func main(){
        e := echo.New()

        e.POST("/login", controllers.LoginUser)
        e.POST("/register", controllers.CreateUser)
        e.GET("/validate", controllers.ValidateToken)

        e.Start(":2020")
}
