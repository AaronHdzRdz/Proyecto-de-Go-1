package routes

import (
	"gin_http/cmd/controllers"
	"gin_http/cmd/middleware"
	"gin_http/cmd/services"
	"github.com/gin-gonic/gin"
)

// User representa un usuario con ID, nombre y correo electrónico.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User // Lista de usuarios simulada.

// SetupUserRoutes configura las rutas relacionadas con los usuarios.
// Recibe el router de Gin y el servicio de usuarios como parámetros.
func SetupUserRoutes(r *gin.Engine, userService *services.UserService) {
	// Grupo de rutas para administradores.
	admin := r.Group("/admin")
	admin.Use(middleware.APIKeyAuthMiddleware()) // Aplica el middleware de autenticación a todas las rutas del grupo.

	//Controller
	UserController := controllers.NewUserController(userService) // Crea una nueva instancia del controlador de usuarios.

	// Ruta para obtener la lista de usuarios.
	admin.GET("/users", UserController.GetUsers)// Obtiene la lista de usuarios del servicio y la devuelve como respuesta.

	// Ruta para crear un nuevo usuario.
	admin.POST("/users", UserController.CreateUser) // Crea un nuevo usuario y lo agrega a la lista.

	// Ruta para actualizar un usuario existente.
	admin.PUT("/users/:user_id", UserController.UpdatedUser)

	// Ruta para eliminar un usuario existente.
	admin.DELETE("/users/:user_id", UserController.DeleteUser) // Elimina un usuario de la lista.
	// Ruta para obtener un usuario por ID.
}
