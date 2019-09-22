package models

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
        "time"
)

type Calendar struct {
        IdFolder int32 `json:"id_folder,omitempty"`
        IdDay uint `json:"id_day"`
        StartTime time.Time `json:"start_time"`
        EndTime time.Time `json:"end_time"`
}

func Create_Calendar(cal *Calendar) (string, error) {
        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "insert into calendar (id_folder, id_day, start_time, end_time) values ($1, $2, $3, $4);"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return "", err
        }

        row := stmt.QueryRow(cal.IdFolder, cal.IdDay, cal.StartTime, cal.EndTime)
        if row == nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return "", err
        }

        return "Se creo correctamente", nil
}
