package models

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
)

type Calendar struct {
        IdFolder int32 `json:"id_folder,omitempty"`
        IdDay uint `json:"id_day"`
        StartTime string `json:"start_time"`
        EndTime string `json:"end_time"`
}

func Create_Calendar(cal *Calendar, id int32) (string, error) {
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

        _, err = stmt.Query(id, cal.IdDay, cal.StartTime, cal.EndTime)
        if err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return "", err
        }

        return "Se creo correctamente", nil
}


func Data_Calendar(id int32) (c []Calendar, err error) {
        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "select c.id_folder, c.id_day, c.start_time, c.end_time from calendar c where c.id_folder = $1;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return nil, err
        }

        rows, err := stmt.Query(id)
        if err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return nil, err
        }

        row := Calendar{}
        for rows.Next() {
                err := rows.Scan(
                        &row.IdFolder,
                        &row.IdDay,
                        &row.StartTime,
                        &row.EndTime,
                )
                if err != nil {
                        fmt.Println("3")
                        return nil, err
                }

                c = append(c, row)
        }
        return c, nil
}
