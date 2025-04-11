package services

import (
	"fmt"
)

// User representa un usuario con un ID, nombre y correo electrónico.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserController es un controlador que maneja las solicitudes relacionadas con los usuarios.
// Contiene un servicio de usuario para manejar la lógica de negocio.

// UserService es un servicio que maneja la lógica de negocio relacionada con los usuarios.
type UserService struct {
	users []User // Lista de usuarios almacenados.
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService() *UserService {
	return &UserService{
		users: []User{}, // Inicializa la lista de usuarios como vacía.
	}
}

// GetUsers devuelve la lista de todos los usuarios.
func (s *UserService) GetUsers() []User {
	return s.users
}

// CreateUser agrega un nuevo usuario a la lista y le asigna un ID único.
func (s *UserService) CreateUser(user User) User {
	user.ID = len(s.users) + 1      // Asigna un ID basado en la longitud actual de la lista.
	s.users = append(s.users, user) // Agrega el usuario a la lista.
	return user
}

/// UpdateUser actualiza la información de un usuario existente.
// Busca el usuario por ID y actualiza su nombre y correo electrónico.
func (s *UserService) UpdateUser(userID int, updatedUser User) (User, error) {
	for i, user := range s.users {
		if user.ID == userID {
			s.users[i].Name = updatedUser.Name
			s.users[i].Email = updatedUser.Email
			return s.users[i], nil // Devuelve el usuario actualizado.
		}
	}
	return User{}, fmt.Errorf("usuario no encontrado") // Devuelve un error si no se encuentra el usuario.
	
}

// DeleteUser elimina un usuario de la lista por su ID.
// Si el usuario no se encuentra, devuelve un error.
func (s *UserService) DeleteUser(userID int) error {
	for i, user := range s.users {
		if user.ID == userID {
			s.users = append(s.users[:i], s.users[i+1:]...) // Elimina el usuario de la lista.
			return nil // Devuelve nil si la eliminación fue exitosa.
		}
	}
	return fmt.Errorf("usuario no encontrado") // Devuelve un error si no se encuentra el usuario.
}