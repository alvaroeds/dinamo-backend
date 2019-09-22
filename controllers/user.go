package controllers

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/commons"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
        "github.com/labstack/echo"
        "log"
        "net/http"
)

//CreateUser funcion para crear un usuario
func CreateUser(c echo.Context) error  {
        user := models.User{}

        //se vuelca la data
        err := c.Bind(&user)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")

                msg.ErrorCode = "user_created_bind"
                msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //se confirma que las contraseñas seas iguales
        if user.Password != user.ConfirmPassword {
                msg.ErrorCode = "user_created_confirm"
                msg.Message = "LAS CONTRASEÑAS NO COINCIDEN"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //codificamos la contraseña antes de introducirla a la BD
        user.Password = commons.Coding(user.Password)

        //agregar validacion para la imagen del usuario

        //se crea el nuevo usuario
        resp, err := models.Create_User(&user)

        //generamos el token
        token := commons.GenerateJWT(user)
        result := token

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "token": result,
                "user": user,
        })
}

//LoginUs er es para que se logueen lo usuarios
func LoginUser(c echo.Context) error {
        //variables
        user := models.User{}

        //volcando al data entrante
        err := c.Bind(&user)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")

                msg.ErrorCode = "user_login_bind"
                msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //codificamos la contraseña antes de introducirla a la BD
        user.Password = commons.Coding(user.Password)

        //se ingresa
        resp, err := models.Login_User(&user)
        if err != nil {
                log.Print(err)
                fmt.Println("ALGO ASALIO MAL")

                msg.ErrorCode = "user_login_bd"
                if (resp != ""){
                        msg.Message = resp
                }
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

        //generamos el token
        token := commons.GenerateJWT(user)
        result := token

        return c.JSON(http.StatusOK, echo.Map{
                "message": resp,
                "token": result,
                "usuario": user,
        })
}


//
func UpdateUser(c echo.Context) error {
        user := models.User{}

        //se lee el json entrante y vuelca en el modelo user
        err := c.Bind(&user)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")
                msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        resp, err := models.Update_User(&user)


        return c.JSON(200, echo.Map{
                "message": resp,
        })
}

//
func DataUser(c echo.Context) error {
        user := models.User{}

        id := c.QueryParam("user")
        fmt.Printf("%s", id)


        resp, err := models.Data_User(id, &user)
        if err != nil {

        }

        return c.JSON(200, echo.Map{
                "message": resp,
                "usuario": user,
        })
}
