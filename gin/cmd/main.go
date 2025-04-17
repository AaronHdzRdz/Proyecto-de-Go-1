package main

import (
	"fmt"
	"gin_http/cmd/database"
	"gin_http/cmd/middleware"
	"gin_http/cmd/routes"
	"gin_http/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// main.go es el punto de entrada de la aplicación.
// Este archivo configura el servidor HTTP utilizando el framework Gin,
// registra middlewares, inicializa servicios y define las rutas de la aplicación.
// main es la función principal que inicializa y ejecuta el servidor HTTP.
// Configura el middleware de registro, inicializa el servicio de usuarios,
// define una ruta de prueba "/ping" y configura las rutas relacionadas con usuarios.
// Finalmente, inicia el servidor en el puerto 3000.
func main() {
	r := gin.Default()

	dbConn := database.NewDataBase()
	defer dbConn.Db.Close()
	

	r.Use(middleware.LoggerMiddleware())

	userService := services.NewUserService(dbConn)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.SetupUserRoutes(r, userService)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Escuchando en el puerto 3000")
}
