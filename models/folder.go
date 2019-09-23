package models

import (
        "fmt"
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/configuration"
)

type Folder struct {
	Id       int32      `json:"id"`
	IdUser   int32      `json:"id_user"`
	IdRoom   *int32      `json:"id_room"`
	Name     string     `json:"name"`
	Tags     string     `json:"tags,omitempty"`
	Calendar []Calendar `json:"calendar,omitempty"`
}

func Create_Folder(folder *Folder) (string, error) {
	//se abre conexion con la base de datos
	db := configuration.GetConnectionPsql()
	defer db.Close()

	//se inserta el usuario
	q := "insert into folder (id_user, id_room, name, tags) values ($1, $2, $3, $4) RETURNING id;"

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("Error al preparar el registro: %s", err)
		return "", err
	}

	row := stmt.QueryRow(folder.IdUser, folder.IdRoom, folder.Name, folder.Tags)
	err = row.Scan(&folder.Id)
	if err != nil {
		fmt.Printf("Error al scanear el registro: %s", err)
		return "", err
	}
        fmt.Println(len(folder.Calendar))
	if folder.Calendar != nil && len(folder.Calendar) > 0 {
		for i := 0; i < len(folder.Calendar); i++ {
		        fmt.Println(folder.Calendar[i])
			resp, err :=Create_Calendar(&folder.Calendar[i], folder.Id)
                        if err != nil {
                                fmt.Println("NO TODO SALIO BIEN")
                        }
			fmt.Println(resp)
		}
	}

	return "Se creo tú espacio", nil
}


func Update_Folder(folder *Folder) (string, error) {
        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "UPDATE folder set id_room=$1, name=$2, tags=$3 WHERE id=$4;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al preparar el registro: %s", err)
                return "", err
        }

        _, err = stmt.Query(folder.IdRoom, folder.Name, folder.Tags, folder.Id)
        if  err != nil {
                fmt.Printf("Error al scanear el registro: %s", err)
                return "", err
        }

        if folder.Calendar != nil && len(folder.Calendar) > 0 {
                fmt.Println("INGRESE AL CALENDAR UPDATE")
                for i := 0; i <= len(folder.Calendar); i++ {
                        //aca insertamos el calendario diractamente
                }
        }

        return "Se actualizó tu espacio", nil
}


func Data_Folder(id string, folder *[]Folder) (string, error) {
        //abrimos conexcion con la BBDD
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se verifica si el usuario existe
        q := "select f.id, f.id_user, f.id_room, f.name, f.tags from folder f where f.id_user=$1"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Println("ERROR AL PREPARAR EL REGISTRO")

                return "", err
        }

        rows, err := stmt.Query(id)
        if err != nil {
                fmt.Println("USUARIO O CLAVE NO VALIDO")

                return "", err
        }
        row := Folder{}
        for rows.Next() {
                err := rows.Scan(
                        &row.Id,
                        &row.IdUser,
                        &row.IdRoom,
                        &row.Name,
                        &row.Tags,
                )
                if err != nil {
                        fmt.Println("3")
                        return "", err
                }

                row.Calendar, err = Data_Calendar(row.Id)
                *folder = append(*folder, row)
        }

        return "Se trajo toda la data", nil
}
