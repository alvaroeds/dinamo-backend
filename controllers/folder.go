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
		fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")

		msg.ErrorCode = "user_created_bind"
		msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
		msg.Error = err.Error()

		return c.JSON(400, msg)
	}

	//se crea el nuevo usuario
	resp, err := models.Create_Folder(&folder)

	return c.JSON(http.StatusOK, echo.Map{
		"message":   resp,
		"folder_id": folder.Id,
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

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "folder_id": folder.Id,
        })
}

/*
func DataFolder(c echo.Context) error  {

        id := c.QueryParam("user")
        fmt.Printf("%s", id)


        //se crea el nuevo usuario
        folder, err := models.Data_Folder(id)

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "folder_id": folder.Id,
        })
}
*/
