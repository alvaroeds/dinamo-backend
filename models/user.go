package models

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
        "github.com/labstack/gommon/log"
)

type User struct {
        Id              uint   `json:"id"`
        Email           string `json:"email,omitempty"`
        Name            string `json:"name,omitempty"`
        LastName        string `json:"lastname,omitempty"`
        Password        string `json:"password,omitempty"`
        ConfirmPassword string `json:"confirmpassword,omitempty"`
        PhoneNumber     string `json:"phone_number"`
}


func Create_User(user *User)  (string, error){
        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "insert into user_dinamo (email,name,lastname,password,phone_number) values ($1,$2,$3,$4,$5) RETURNING id;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return "", err
        }

        row := stmt.QueryRow(user.Email,user.Name,user.LastName,user.Password,user.PhoneNumber)
        err = row.Scan(&user.Id)
        if err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return "", err
        }
        user.Password = ""
        user.ConfirmPassword = ""

        return "Se registró correctamente", nil
}

func Login_User(user *User) (string, error) {
        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "SELECT u.id, u.name, u.lastname, u.phone_number FROM user_dinamo u WHERE u.email=$1 AND u.password=$2;"

        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                return "", err
        }

        row := stmt.QueryRow(user.Email, user.Password)
        if row == nil {
                log.Print(err)
                fmt.Println("USUARIO O CLAVE NO VALIDO")

                return "USUARIO O CLAVE NO VALIDO", err
        }
        user.Password = ""

        err = row.Scan(&user.Id, &user.Name, &user.LastName, &user.PhoneNumber)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL TRAER LA DATA")

                return "", err
        }

        return "Ingresó correctamente", nil
}

func Update_User(user *User) (string, error)  {
        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "UPDATE user_dinamo set name=$1, lastname=$2, phone_number=$3 WHERE id=$4;"

        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                return "", err
        }

        row := stmt.QueryRow(user.Name, user.LastName, user.PhoneNumber, user.Id)
        if row == nil {
                return "", err
        }

        return "Se actualizo su información", nil
}

func Data_User(id string, user *User) (string, error) {
        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "SELECT u.id, u.name, u.lastname, u.phone_number, u.email FROM user_dinamo u WHERE u.id=$1;"

        stmt, err := db.Prepare(q)
        if err != nil {
                log.Print(err)
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                return "", err
        }

        row := stmt.QueryRow(id)
        if row == nil {
                log.Print(err)

                return "", err
        }

        err = row.Scan(&user.Id, &user.Name, &user.LastName, &user.PhoneNumber, &user.Email)
        if err != nil {
                log.Print(err)

                return "", err
        }
        user.Password = ""

        return "Se trajo toda la data", nil
}
