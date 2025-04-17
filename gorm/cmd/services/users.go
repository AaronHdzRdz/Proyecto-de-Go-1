package services

import (
	"fmt"
	"gin_http/cmd/database"
	"gin_http/cmd/models"
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
    var user []User
    result := s.db.Db.Where("deleted_at IS NULL").Find(&user)
    if result.Error != nil {
        return []User{}, result.Error
    }
    if result.RowsAffected == 0 {
        return []User{}, nil
    }
    var userResult []User
    for _, user := range user {
        userResult = append(userResult, User{
            ID:    int(user.ID),
            Name:  user.Name,
            Email: user.Email,
        })
    }
    return userResult, nil
}

// CreateUser agrega un nuevo usuario a la lista y le asigna un ID único.
func (s *UserService) CreateUser(user User) (User, error) {
	userDb:= models.User{
		Name:  user.Name,
		Email: user.Email,
	}
	result:= s.db.Db.Create(&userDb)
	if result.Error != nil {
		fmt.Println("Error al crear usuario: ", result.Error)
		return User{}, nil
	}
	user.ID = int(userDb.ID)
	return user, nil
}

// UpdateUser actualiza la información de un usuario existente.
// Busca el usuario por ID y actualiza su nombre y correo electrónico.
func (s *UserService) UpdateUser(userID int, updatedUser User) (User, error) {
    var user models.User
    // Busca el usuario por ID
    resultRead := s.db.Db.First(&user, userID)
    if resultRead.Error != nil {
        fmt.Println("Error al buscar usuario: ", resultRead.Error)
        return User{}, resultRead.Error
    }

    // Actualiza el usuario con los nuevos datos
    resultUpdate := s.db.Db.Model(&user).Updates(models.User{
        Name:  updatedUser.Name,
        Email: updatedUser.Email,
    })
    if resultUpdate.Error != nil {
        fmt.Println("Error al actualizar usuario: ", resultUpdate.Error)
        return User{}, resultUpdate.Error
    }

    // Devuelve el usuario actualizado con el ID correcto
    return User{
        ID:    int(user.ID),
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

// DeleteUser elimina un usuario de la lista por su ID.
// Si el usuario no se encuentra, devuelve un error.
func (s *UserService) DeleteUser(userID int) error {
	var user models.User
	resultRead := s.db.Db.Where("deleted_at IS NULL").Find(&user, userID)
	if resultRead.Error != nil {
		fmt.Println("Error al buscar usuario: ", resultRead.Error)
		return resultRead.Error
	}
	resultDelete := s.db.Db.Delete(&user, userID)
	if resultDelete.Error != nil {
		fmt.Println("Error al eliminar usuario: ", resultDelete.Error)
		return resultDelete.Error
	}
	return nil
}