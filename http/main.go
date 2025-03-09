package main

import (
    "encoding/json" // Paquete para trabajar con JSON
    "fmt"           // Paquete para formateo de entrada/salida
    "io"            // Paquete para operaciones de entrada/salida
    "net/http"      // Paquete para crear servidores HTTP
    "os"            // Paquete para operaciones del sistema operativo
)

// Definición de la estructura Usuario con etiquetas JSON
type Usuario struct {
    ID     int    `json:"id"`
    Nombre string `json:"nombre"`
    Email  string `json:"email"`
}

// Variable global para almacenar los usuarios
var usuarios []Usuario

// Función para verificar si un usuario existe por su ID
func ExisteUsuario(id int) bool {
    for _, u := range usuarios {
        if u.ID == id {
            return true
        }
    }
    return false
}

// Manejador para la ruta GET /v1/users
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Test-Header", "header")
    json.NewEncoder(w).Encode(usuarios) // Codifica la lista de usuarios a JSON y la envía en la respuesta
}

// Manejador para la ruta POST /v1/users
func handlePostUser(w http.ResponseWriter, r *http.Request) {
    // Lee el cuerpo de la solicitud
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error leyendo el body", http.StatusBadRequest)
        return
    }
    fmt.Println("Body recibido:", string(body))
    
    // Decodifica el JSON del cuerpo en un objeto Usuario
    var usuario Usuario
    err = json.Unmarshal(body, &usuario)
    if err != nil {
        http.Error(w, "Error leyendo el JSON", http.StatusBadRequest)
        return
    }
    
    // Asigna un nuevo ID al usuario en función del último usuario en la lista
    if len(usuarios) > 0 {
        usuario.ID = usuarios[len(usuarios)-1].ID + 1
    } else {
        usuario.ID = 1
    }
    
    // Agrega el nuevo usuario a la lista
    usuarios = append(usuarios, usuario)
    
    // Configura los encabezados de la respuesta
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Test-Header", "header")
    
    // Envía el usuario creado en la respuesta
    json.NewEncoder(w).Encode(usuario)
}

// Manejador para la ruta PUT /v1/users
func handlePutUser(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body) // Lee el cuerpo de la solicitud
    if err != nil {
        http.Error(w, "Error leyendo el body", http.StatusBadRequest)
        return
    }
    var usuario Usuario
    err = json.Unmarshal(body, &usuario) // Decodifica el JSON del cuerpo en un objeto Usuario
    if err != nil {
        http.Error(w, "Error leyendo el JSON", http.StatusBadRequest)
        return
    }
    index := -1
    for i, u := range usuarios {
        if u.ID == usuario.ID {
            index = i
            break
        }
    }

    if index == -1 {
        http.Error(w, "Usuario no encontrado", http.StatusNotFound)
        return
    }
    // Actualiza los datos del usuario
    usuarios[index].Nombre = usuario.Nombre
    usuarios[index].Email = usuario.Email

    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "message": fmt.Sprintf("Usuario %s actualizado", usuario.Nombre),
    }
    json.NewEncoder(w).Encode(response) // Envía una respuesta indicando que el usuario fue actualizado
}

// Manejador para la ruta DELETE /v1/users
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body) // Lee el cuerpo de la solicitud
    if err != nil {
        http.Error(w, "Error leyendo el body", http.StatusBadRequest)
        return
    }
    var usuario Usuario
    err = json.Unmarshal(body, &usuario) // Decodifica el JSON del cuerpo en un objeto Usuario
    if err != nil {
        http.Error(w, "Error leyendo el JSON", http.StatusBadRequest)
        return
    }

    index := -1
    for i, u := range usuarios {
        if u.ID == usuario.ID {
            index = i
            break
        }
    }

    if index == -1 {
        http.Error(w, "Usuario no encontrado", http.StatusNotFound)
        return
    }

    eliminado := usuarios[index]
    usuarios = append(usuarios[:index], usuarios[index+1:]...) // Elimina el usuario de la lista

    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "message": fmt.Sprintf("Usuario %s eliminado", eliminado.Nombre),
    }
    json.NewEncoder(w).Encode(response) // Envía una respuesta indicando que el usuario fue eliminado
}

// Manejador principal para la ruta /v1/users
func Users(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGetUsers(w, r)
    case http.MethodPost:
        handlePostUser(w, r)
    case http.MethodPut:
        handlePutUser(w, r)
    case http.MethodDelete:
        handleDeleteUser(w, r)
    default:
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
    }
}

// Manejador para la ruta /ping
func Ping(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintf(w, "pong") // Responde con "pong" para solicitudes GET
    default:
        http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
    }
}

// Manejador para la ruta raíz /
func Index(w http.ResponseWriter, r *http.Request) {
    content, err := os.ReadFile("public/index.html") // Lee el archivo HTML
    if err != nil {
        fmt.Fprintf(w, "error leyendo el html")
        return
    }
    fmt.Fprintf(w, string(content)) // Envía el contenido del archivo HTML en la respuesta
}

// Función principal
func main() {
    // Inicializa la lista de usuarios con algunos datos
    usuarios = append(usuarios, Usuario{ID: 1, Nombre: "Eduardo", Email: "shernan@alumnos.uaq.mx"})
    usuarios = append(usuarios, Usuario{ID: 2, Nombre: "alan", Email: "alan@alumnos.uaq.mx"})
    
    // Define las rutas y sus manejadores
    http.HandleFunc("/ping", Ping)
    http.HandleFunc("/", Index)
    http.HandleFunc("/v1/users", Users)
    
    // Inicia el servidor en el puerto 3000
    fmt.Println("Servidor corriendo en el puerto 3000")
    http.ListenAndServe(":3000", nil)
}