package main

import (
	"fmt"
	"gin_http/cmd/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "poa|ng",
		})
	})
	routes.SetupUserRoutes(r)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Escuchando en el puerto 3000")
}
