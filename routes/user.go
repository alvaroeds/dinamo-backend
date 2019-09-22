package routes

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func user(e *echo.Echo)  {
        //crud
        e.POST("/api/v1/register", controllers.CreateUser)
        e.POST("/api/v1/login", controllers.LoginUser)

        //traer la data
        e.PUT("/api/v1/user", controllers.UpdateUser)
        e.GET("/api/v1/user", controllers.DataUser)
}
