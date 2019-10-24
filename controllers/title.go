package controllers

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
        "github.com/labstack/echo"
        "log"
        "net/http"
)

func CreateTitle(c echo.Context) error {
        //variables
        title := models.Title{}

        //volcaldo la data entrante
        err := c.Bind(&title)
        if err != nil {
                log.Println(err)

                msg.ErrorCode = "title_created_bind"
                msg.Message = "Error al aceptar la data entrante"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //se crea un nuevo titulo
        resp, err := models.Create_Title(&title)
        if err != nil {
                msg.ErrorCode = "title_created_bad"
                msg.Message = "Algo salio mal"
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

        return c.JSON(http.StatusOK, echo.Map{
                "message":   resp,
                "title": title,
        })
}
