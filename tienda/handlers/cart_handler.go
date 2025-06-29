package handlers

import (
	"encoding/json"
	"net/http"
	"tienda/models"
	"tienda/storage"

	"github.com/gorilla/mux"
)

// CartHandlers necesita dependencias de carritos, productos y órdenes.
type CartHandlers struct {
	cartStore    storage.CartStorer
	productStore storage.ProductStorer
	orderStore   storage.OrderStorer
}

// NewCartHandlers es el constructor que inyecta todas las dependencias.
func NewCartHandlers(cs storage.CartStorer, ps storage.ProductStorer, os storage.OrderStorer) *CartHandlers {
	return &CartHandlers{cartStore: cs, productStore: ps, orderStore: os}
}

// CreateCartHandler crea un nuevo carrito de compras vacío.
func (h *CartHandlers) CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	newCart := models.Cart{Items: []models.CartItem{}, Total: 0}
	createdCart, err := h.cartStore.CreateCart(newCart)
	if err != nil {
		http.Error(w, "Error al crear el carrito", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCart)
}

// GetCartHandler obtiene el contenido de un carrito.
func (h *CartHandlers) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]
	cart, err := h.cartStore.GetCartByID(cartId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// AddItemToCartHandler añade un producto a un carrito.
func (h *CartHandlers) AddItemToCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]
	var req struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	if req.Quantity <= 0 {
		http.Error(w, "La cantidad debe ser positiva", http.StatusBadRequest)
		return
	}
	product, err := h.productStore.GetProductByID(req.ProductID)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}
	cart, err := h.cartStore.GetCartByID(cartId)
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	// Lógica para añadir ítem o actualizar cantidad.
	found := false
	for i, item := range cart.Items {
		if item.ProductID == req.ProductID {
			cart.Items[i].Quantity += req.Quantity
			found = true
			break
		}
	}
	if !found {
		newItem := models.CartItem{ProductID: product.ID, Quantity: req.Quantity, Price: product.Price}
		cart.Items = append(cart.Items, newItem)
	}
	// Recalcular total.
	var total float64
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	cart.Total = total
	updatedCart, err := h.cartStore.UpdateCart(cartId, cart)
	if err != nil {
		http.Error(w, "Error al actualizar el carrito", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCart)
}

// RemoveItemFromCartHandler elimina un producto del carrito.
func (h *CartHandlers) RemoveItemFromCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId, productId := vars["cartId"], vars["productId"]
	cart, err := h.cartStore.GetCartByID(cartId)
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	// Lógica para quitar el ítem del slice.
	itemFound := false
	newItems := []models.CartItem{}
	for _, item := range cart.Items {
		if item.ProductID == productId {
			itemFound = true
		} else {
			newItems = append(newItems, item)
		}
	}
	if !itemFound {
		http.Error(w, "Producto no encontrado en el carrito", http.StatusNotFound)
		return
	}
	cart.Items = newItems
	// Recalcular total.
	var total float64
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	cart.Total = total
	updatedCart, err := h.cartStore.UpdateCart(cartId, cart)
	if err != nil {
		http.Error(w, "Error al actualizar el carrito", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCart)
}

// CheckoutHandler finaliza la compra.
func (h *CartHandlers) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]
	cart, err := h.cartStore.GetCartByID(cartId)
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}
	// Guarda el carrito en el historial de órdenes.
	if err := h.orderStore.CreateOrderFromCart(cart); err != nil {
		http.Error(w, "Error al procesar la orden", http.StatusInternalServerError)
		return
	}
	// Elimina el carrito activo.
	if err := h.cartStore.DeleteCart(cartId); err != nil {
		// No es un error crítico, se puede loguear para mantenimiento.
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "¡Compra realizada con éxito!"})
}

// DeleteCartHandler vacía un carrito sin completar la compra.
func (h *CartHandlers) DeleteCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartId := vars["cartId"]
	if err := h.cartStore.DeleteCart(cartId); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
