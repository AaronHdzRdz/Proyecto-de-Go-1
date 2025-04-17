package main

import (
	"fmt"
	"gin_http/cmd/database"
	"gin_http/cmd/models"
	"gin_http/cmd/routes"
	"gin_http/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dbConn:= database.NewDataBase()
	err:= dbConn.Db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error al migrar la base de datos: ", err)
	}
	userService := services.NewUserService(dbConn)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status": "ok",
		})
	})
	routes.SetupUserRoutes(r, userService)
	r.Run(":3000")
	fmt.Println("Escuchando el puerto 3000")
}