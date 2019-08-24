package user

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/commons"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/models"
        "github.com/labstack/echo"
        "log"
        "net/http"
)

func Service_user_login(c echo.Context) error {
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
