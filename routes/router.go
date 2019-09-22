package routes

import (
        "github.com/alvaroenriqueds/Dinamo/dinamo-backend/controllers"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        "net/http"
)

//func Init() *echo.Echo {
func Init(e *echo.Echo)  {
        //se inicia echo


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

        //Se crean los endpoint
        user_Login(e)



        e.POST("/upload", controllers.Upload)
        e.GET("/validate", controllers.ValidateToken)
}
