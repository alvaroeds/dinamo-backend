package controllers

import (
	"fmt"
	"github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func CreateFolder(c echo.Context) error {
	folder := models.Folder{}

	//se vuelca la data
	err := c.Bind(&folder)
	if err != nil {
		log.Print(err)
		fmt.Println("ERROR AL VOLCAR LA DATA ENTRANTE")

		msg.ErrorCode = "folder_created_bind"
		msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
		msg.Error = err.Error()

		return c.JSON(400, msg)
	}

	//se crea el nuevo usuario
	resp, err := models.Create_Folder(&folder)
        if err != nil {
                log.Print(err)

                msg.ErrorCode = "folder_created_bind"
                msg.Message = "ERROR AL CREAR TU ESPACIO"
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

	return c.JSON(http.StatusOK, echo.Map{
		"message":   resp,
		"folder": folder,
	})
}

func UpdateFolder(c echo.Context) error  {
        folder := models.Folder{}

        //se vuelca la data
        err := c.Bind(&folder)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")

                msg.ErrorCode = "user_created_bind"
                msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //se crea el nuevo usuario
        resp, err := models.Update_Folder(&folder)
        if err != nil {
                log.Print(err)

                msg.ErrorCode = "folder_update_null"
                msg.Message = "ERROR"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "folder_id": folder.Id,
        })
}


func DataFolder(c echo.Context) error  {

        id := c.QueryParam("user")
        fmt.Printf("%s", id)

        folder := []models.Folder{}

        //se crea el nuevo usuario
        resp, err := models.Data_Folder(id, &folder)
        if err != nil {
                return c.NoContent(500)
        }

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "folder": folder,
        })
}

