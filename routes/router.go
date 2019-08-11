package routes

import (
        "../controllers"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        "net/http"
)

//func Init() *echo.Echo {
func Init()  {
        //se inicia echo
        e := echo.New()

        // Middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
                // AllowOrigins: []string{"https://potencie.com", "https://potencie.com"},
                AllowOrigins: []string{"*"},
                AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
        }))

        //se levanta el front de pruebas
        e.Static("/", "../public")

        //Se crean los endpoints
        e.POST("/login", controllers.LoginUser)
        e.POST("/register", controllers.CreateUser)
        e.GET("/validate", controllers.ValidateToken)
        e.POST("/upload", controllers.Upload)

        //se inicia el servidor en el puerto x
        //e.Start(":2020")
        e.Logger.Fatal(
                e.Start(":2020"),
                )

        //return  e
}
