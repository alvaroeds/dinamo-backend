package controllers

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/commons"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
        "github.com/labstack/echo"
        "log"
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
        //variables/obejtos
        user := models.User{}
        msg := models.Error{}

        //volvando al data entrante
        err := c.Bind(&user)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL VOLCAS LA DATA ENTRANTE")

                msg.ErrorCode = "user_login_bind"
                msg.Message = "ERROR AL ACEPTAR LA DATA ENTRANTE"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        /*
           //confirmammos ek pass y confirm_pass
           if user.Password != user.ConfirmPassword {
                   msg.ErrorCode = "user_login_confirm"
                   msg.Message = "LAS CONTRASEÑAS NO COINCIDEN"
                   msg.Error = err.Error()

                   return c.JSON(400, msg)
           }
        */

        //codificamos la contraseña antes de introducirla a la BD
        user.Password = commons.Coding(user.Password)

        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "SELECT u.id, u.name, u.lastname, u.numero FROM usuario u WHERE u.email=$1 AND u.password=$2;"

        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                msg.ErrorCode = "user_login_prepare"
                msg.Message = "ERROR AL PREPARAR EL REGISTRO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        row := stmt.QueryRow(user.Email, user.Password)
        if row == nil {
                log.Print(err)
                fmt.Println("USUARIO O CLAVE NO VALIDO")

                msg.ErrorCode = "user_login_row"
                msg.Message = "LAS USUARIO O CLAVE NO VALIDO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }
        user.Password = ""

        err = row.Scan(&user.Id, &user.Name, &user.LastName, &user.Numero)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL TRAER LA DATA")

                msg.ErrorCode = "user_login_scan"
                msg.Message = "ERROR AL TRAER LA DATA"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        //generamos el token
        token := commons.GenerateJWT(user)
        result := token

        return c.JSON(http.StatusOK, echo.Map{
                "token": result,
                "usuario": user,
        })
}


//
func UpdateUser(c echo.Context) error {
        user := models.User{}
        msg := models.Error{}

        //se lee el json entrante y vuelca en el modelo user
        err := json.NewDecoder(c.Request().Body).Decode(&user)
        if err != nil {
                fmt.Printf("Error al leer el usuario a registrar: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        //codificamos la contraseña antes de introducirla a la BD
        user.Password = commons.Coding(user.Password)

        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "UPDATE usuario set name=$1, lastname=$2, numero=$3 WHERE id=$4;"



        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                msg.ErrorCode = "user_login_prepare"
                msg.Message = "ERROR AL PREPARAR EL REGISTRO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        row := stmt.QueryRow(user.Name, user.LastName, user.Numero, user.Id)
        if row == nil {
                log.Print(err)
                fmt.Println("USUARIO O CLAVE NO VALIDO")

                msg.ErrorCode = "user_login_row"
                msg.Message = "LAS USUARIO O CLAVE NO VALIDO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }
        user.Password = ""

        return c.JSON(200, echo.Map{
                "message": "Se actualizo sus datos",
        })
}

//
func DataUser(c echo.Context) error {
        user := models.User{}
        msg := models.Error{}

        id := c.QueryParam("user")
        fmt.Printf("%s", id)
        fmt.Println("")
        fmt.Println(id)
        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "SELECT u.id, u.name, u.lastname, u.numero, u.email FROM usuario u WHERE u.id=$1;"

        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                msg.ErrorCode = "user_login_prepare"
                msg.Message = "ERROR AL PREPARAR EL REGISTRO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        row := stmt.QueryRow(id)
        if row == nil {
                log.Print(err)
                fmt.Println("USUARIO O CLAVE NO VALIDO")

                msg.ErrorCode = "user_login_row"
                msg.Message = "LAS USUARIO O CLAVE NO VALIDO"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        row.Scan(&user.Id, &user.Name, &user.LastName, &user.Numero, &user.Email)
        user.Password = ""

        return c.JSON(200, user)
}
