package routes

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
)

func calendar(e *echo.Echo)  {
        //crud
        e.POST("api/v1/create/calendar", controllers.CreateCalendar)

}
