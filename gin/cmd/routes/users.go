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
		userID := c.Param("user_id")
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error de parseo del body"})
			return
		}
		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == userID {
				users[i].Name = updatedUser.Name
				users[i].Email = updatedUser.Email
				c.JSON(http.StatusOK, users[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	})
	r.DELETE("/users/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == userID {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	})
}
