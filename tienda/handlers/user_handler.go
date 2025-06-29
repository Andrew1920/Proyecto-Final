package handlers

import (
	"encoding/json"
	"net/http"
	"tienda/models"
	"tienda/storage"
	"tienda/utils"
)

// UserHandlers maneja la lógica de usuarios.
type UserHandlers struct {
	store storage.UserStorer
}

// NewUserHandlers es el constructor para los handlers de usuario.
func NewUserHandlers(s storage.UserStorer) *UserHandlers {
	return &UserHandlers{store: s}
}

// RegisterHandler crea nuevas cuentas de usuario.
func (h *UserHandlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}
	// Hashear la contraseña antes de guardarla.
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error al procesar la contraseña", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	createdUser, err := h.store.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict) // Usuario ya existe.
		return
	}
	createdUser.Password = "" // No devolver la contraseña hasheada.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// LoginHandler verifica las credenciales de un usuario.
func (h *UserHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}
	user, err := h.store.GetUserByUsername(credentials.Username)
	if err != nil || !utils.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Inicio de sesión exitoso"})
}
