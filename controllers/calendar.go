package controllers

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
        "github.com/labstack/echo"
        "log"
        "net/http"
)

func CreateCalendar(c echo.Context) error  {
        //se crea las variables
        cal := models.Calendar{}

        err := c.Bind(&cal)
        if err != nil {
                //log del error
                log.Println(err)
                //generamos la estructura error
                msg.ErrorCode = "calendar_created_bind"
                msg.Message= "La data que enviaste "
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        resp, err := models.Create_Calendar(&cal, cal.IdFolder)
        if err != nil {
                log.Println(err)

                msg.ErrorCode = "calendar_created_model"
                msg.Message = "Algo salio mal, intenelo luego"
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "calendar": cal,
        })
}
