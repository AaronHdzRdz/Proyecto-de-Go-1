package services

import (
	"fmt"
	"gin_http/cmd/database"
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
	db *database.Database // Conexión a la base de datos.
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(connection *database.Database) *UserService {
	return &UserService{
		users: []User{}, // Inicializa la lista de usuarios vacía.
		db:   connection, // Conexión a la base de datos.
	}
}

// GetUsers devuelve la lista de todos los usuarios.
func (s *UserService) GetUsers() ([]User, error) {
	rows, err := s.db.Db.Query("SELECT id, name, email FROM users")
	var users []User
	if err != nil {
		fmt.Printf("Error al obtener usuarios: %v\n", err)
		return []User{}, err // Devuelve un error si la consulta falla.
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Printf("Error al escanear usuario: %v\n", err)
			continue
		}
		users = append(users, user) // Agrega el usuario a la lista.
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("Error al iterar sobre usuarios: %v\n", err)
		return users, err // Devuelve un error si ocurre durante la iteración.	
	}
	return users, nil
}

// CreateUser agrega un nuevo usuario a la lista y le asigna un ID único.
func (s *UserService) CreateUser(user User) (User, error) {
	insertStmt := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	var user_id int // Variable para almacenar el ID del nuevo usuario.
	err := s.db.Db.QueryRow(insertStmt, user.Name, user.Email).Scan(&user_id) // Ejecuta la consulta de inserción y escanea el ID devuelto.
	if err != nil {
		fmt.Printf("Error al insertar usuario: %v\n", err) // Maneja el error si ocurre.
		return User{}, err // Devuelve un usuario vacío en caso de error.
	}
	user.ID = user_id // Asigna el ID devuelto al usuario.
	return user, nil
}

// UpdateUser actualiza la información de un usuario existente.
// Busca el usuario por ID y actualiza su nombre y correo electrónico.
func (s *UserService) UpdateUser(userID int, updatedUser User) (User, error) {
	updateStmt := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	result, err := s.db.Db.Exec(updateStmt, updatedUser.Name, updatedUser.Email, userID)
	if err != nil {
		fmt.Printf("Error al actualizar usuario: %v\n", err)
		return User{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Error al obtener filas afectadas: %v\n", err)
		return User{}, err
	}

	if rowsAffected == 0 {
		return User{}, fmt.Errorf("usuario con ID %d no encontrado", userID)
	}

	updatedUser.ID = userID
	return updatedUser, nil
}

// DeleteUser elimina un usuario de la lista por su ID.
// Si el usuario no se encuentra, devuelve un error.
func (s *UserService) DeleteUser(userID int) error {
	deleteStmt := "DELETE FROM users WHERE id = $1"
	result, err := s.db.Db.Exec(deleteStmt, userID)
	if err != nil {
		fmt.Printf("Error al eliminar usuario: %v\n", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Error al obtener filas afectadas: %v\n", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", userID)
	}

	return nil
}