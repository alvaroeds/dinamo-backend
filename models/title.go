package models

import (
	"fmt"
	"github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
        "time"
)

type Title struct {
	Id         int32  `json:"id"`
	IdFolder   int32  `json:"id_folder"`
	DateClient string  `json:"date_client"`
	Name       string `json:"name"`
}

func Create_Title(title *Title) (string, error) {
	//se abre conexion con la base de datos
	db := configuration.GetConnectionPsql()
	defer db.Close()

	//se inserta el usuario
	q := "insert into title (id_folder, date_client, name, created_at) values ($1, $2, $3, $4) RETURNING id;"

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al preparar el registro: %s", err)
		return "", err
	}

	ca := time.Now()
        fmt.Println(title)
	fmt.Println(ca)
	row := stmt.QueryRow(title.IdFolder, title.DateClient, title.Name, ca)
	err = row.Scan(&title.Id)
	if err != nil {
		fmt.Printf("Error al scanear el registro: %s", err)
		return "", err
	}

	return "Se creo el titulo", nil
}
