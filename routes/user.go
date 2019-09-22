package routes

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func user_Login(e *echo.Echo)  {
        e.POST("/login", controllers.LoginUser)
        e.POST("/register", controllers.CreateUser)

        //
        e.PUT("/user", controllers.UpdateUser)
        e.GET("/user", controllers.DataUser)
}
