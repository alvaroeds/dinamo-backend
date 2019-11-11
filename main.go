package main

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/routes"
        "github.com/labstack/echo"
)

func main(){
        //se inicia la instancia echo
        e := echo.New()

        //se activa el router
        routes.Init(e)

        //se inicia el servidor en el puerto x
        e.Logger.Fatal(
                e.Start(":5051"),
        )
}
