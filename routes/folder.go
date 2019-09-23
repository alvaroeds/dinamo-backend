package routes

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func folder(e *echo.Echo)  {
        //crud
        e.POST("/api/v1/create/folder", controllers.CreateFolder)

        //traer la data
        e.PUT("/api/v1/folder", controllers.UpdateFolder)
        e.GET("/api/v1/folder", controllers.DataFolder)
}
