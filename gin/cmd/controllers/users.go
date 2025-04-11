package controllers

import (
	"encoding/json"
	"gin_http/cmd/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController maneja las solicitudes relacionadas con los usuarios.
type UserController struct {
	userService *services.UserService // Servicio de usuario para manejar la lógica de negocio.
}

// NewUserController crea una nueva instancia de UserController.
// Recibe un servicio de usuario como dependencia.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService, // Inicializa el controlador con el servicio de usuario.
	}
}

// GetUsers maneja la solicitud para obtener la lista de usuarios.
// Responde con un JSON que contiene todos los usuarios.
func (s *UserController) GetUsers(c *gin.Context) {
	users := s.userService.GetUsers() // Obtiene la lista de usuarios del servicio.
	c.JSON(http.StatusOK, users)      // Devuelve la lista de usuarios en formato JSON.
}

// CreateUser maneja la solicitud para crear un nuevo usuario.
// Lee el cuerpo de la solicitud y responde con el contenido recibido.
func (s *UserController) CreateUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body) // Lee el cuerpo de la solicitud.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error leyendo el body"}) // Devuelve un error si no se puede leer el cuerpo.
		return
	}

	var user services.User            // Declara una variable para almacenar el usuario.
	err = json.Unmarshal(body, &user) // Deserializa el cuerpo en la variable user.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error deserializando el body"}) // Devuelve un error si no se puede deserializar el cuerpo.
		return
	}
	user = s.userService.CreateUser(user) // Crea un nuevo usuario utilizando el servicio.
	c.JSON(http.StatusOK, user)           // Devuelve el usuario creado como respuesta.
}

// UpdatedUser maneja la solicitud para actualizar un usuario existente.
// Obtiene el ID del usuario de los parámetros de la URL y responde con un mensaje de confirmación.
func (s *UserController) UpdatedUser(c *gin.Context) {
	userID := c.Param("user_id")                           // Obtiene el ID del usuario de los parámetros de la URL.
	var UpdatedUser services.User                          // Declara una variable para almacenar el usuario actualizado.
	if err := c.ShouldBindJSON(&UpdatedUser); err != nil { // Intenta deserializar el cuerpo de la solicitud en la variable UpdatedUser.
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON Invalido"}) // Devuelve un mensaje si no se puede deserializar.
		return
	}

	idInt, err := strconv.Atoi(userID) // Convierte el ID del usuario a un entero.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID Invalido"}) // Devuelve un mensaje si el ID no es válido.
		return
	}

	user, err := s.userService.UpdateUser(idInt, UpdatedUser) // Intenta actualizar el usuario utilizando el servicio.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Usuario no encontrado"}) // Devuelve un mensaje si el usuario no se encuentra.
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser maneja la solicitud para eliminar un usuario.
// Obtiene el ID del usuario de los parámetros de la URL y responde con un mensaje de confirmación.
func (s *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("user_id")       // Cambiar "id" por "user_id" para que coincida con la ruta.
	idInt, err := strconv.Atoi(userID) // Convierte el ID del usuario a un entero.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID Invalido"}) // Devuelve un mensaje si el ID no es válido.
		return
	}

	err = s.userService.DeleteUser(idInt) // Intenta eliminar el usuario utilizando el servicio.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Usuario no encontrado"}) // Devuelve un mensaje si el usuario no se encuentra.
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario eliminado",
		"id":      userID, // Devuelve un mensaje de confirmación con el ID del usuario eliminado.
	})
}
