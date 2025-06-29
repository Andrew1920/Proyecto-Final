package models

// User define la estructura de un usuario.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // No se expone en respuestas JSON.
}
