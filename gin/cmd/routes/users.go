package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func SetupUserRoutes(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		fmt.Println("User-Agent", userAgent)
		c.Header("x-User-Agent", "gin")
		c.JSON((http.StatusOK), users)
	})
	r.POST("/users", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error leyendo el body"})
			return
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error de parseo del body"})
			return
		}
		user.ID = len(users) + 1
		users = append(users, user)

		c.JSON(http.StatusCreated, user)
	})
	r.PUT("/users/:user_id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("user id", id)
	})
	r.DELETE("/users/:user_id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("user id", id)
	})
}
