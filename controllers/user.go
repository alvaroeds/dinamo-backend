package controllers

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/dinamo-backend/commons"
        "github.com/alvaroenriqueds/dinamo-backend/configuration"
        "github.com/alvaroenriqueds/dinamo-backend/models"
        "github.com/labstack/echo"
        "net/http"
)

//CreateUser funcion para crear un usuario
func CreateUser(c echo.Context) error  {
        user := models.User{}

        //se lee el json entrante y vuelca en el modelo user
        err := json.NewDecoder(c.Request().Body).Decode(&user)
        if err != nil {
                fmt.Printf("Error al leer el usuario a registrar: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        //se confirma que las contraseñas seas iguales
        if user.Password != user.ConfirmPassword {
                fmt.Printf("Las contraseñas no coinciden: %s | %s", user.Password, user.ConfirmPassword)
                return c.NoContent(http.StatusBadRequest)
        }

        //se codifica la contraseña en sha256
        pass := sha256.Sum256([]byte(user.Password))
        pwd := fmt.Sprintf("%x", pass)
        user.Password = pwd

        //agregar validacion para la imagen del usuario

        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "insert into usuario (email, password, name, lastname, numero) values ($1, $2, $3, $4, $5) RETURNING id;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        row := stmt.QueryRow(user.Email,user.Password, user.Name, user.LastName, user.Numero)
        err = row.Scan(&user.Id)
        if err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }
        user.Password = ""
        user.ConfirmPassword = ""

        //generamos el token
        token := commons.GenerateJWT(user)
        result := token

        return c.JSON(http.StatusOK, echo.Map{
                "token": result,
                "usuario": user,
        })
}

//LoginUs er es para que se logueen lo usuarios
func LoginUser(c echo.Context) error {
        user := models.User{}

        //se lee el json entrante y se vuelva en user
        err := json.NewDecoder(c.Request().Body).Decode(&user)
        if err != nil {
                fmt.Fprintf(c.Response(), "Error: %s\n", err)
                return c.NoContent(http.StatusBadRequest)
        }

        fmt.Println(user)
        //se codifica la contraseña a sha256
        pass := sha256.Sum256([]byte(user.Password))
        pwd := fmt.Sprintf("%x", pass)
        user.Password = pwd

        fmt.Println(user)
        fmt.Println(user.Email)


        //se abre una conexion con al BD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "SELECT u.id, u.name, u.lastname, u.numero FROM usuario u WHERE u.email=$1 AND u.password=$2;"


        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }


        row := stmt.QueryRow(user.Email, user.Password)
        if row == nil {
                fmt.Printf("Usuario o clave no valido: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }
        user.Password = ""

        err = row.Scan(&user.Id, &user.Name, &user.LastName, &user.Numero)
        if err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        //generamos el token
        token := commons.GenerateJWT(user)
        result := token

        /*
        jsonr, err := json.Marshal(result)
        if err != nil {
                log.Fatalf("Error al convertir el token a json: %s", err)
        }
        c.Response().WriteHeader(http.StatusOK)
        c.Response().Header().Set("Content-Type", "application/json")
        c.Response().Write(jsonr)
        */

        return c.JSON(http.StatusOK, echo.Map{
                "token": result,
                "usuario": user,
        })
}
