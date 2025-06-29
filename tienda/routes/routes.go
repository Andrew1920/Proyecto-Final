package routes

import (
	"net/http"
	"tienda/handlers"

	"github.com/gorilla/mux"
)

// RegisterRoutes define todos los endpoints de la API.
func RegisterRoutes(r *mux.Router, ph *handlers.ProductHandlers, ch *handlers.CartHandlers, uh *handlers.UserHandlers, rh *handlers.ReportHandlers) {
	// Rutas de Usuario
	r.HandleFunc("/register", uh.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", uh.LoginHandler).Methods("POST")

	// Rutas de Productos
	r.HandleFunc("/api/products", ph.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products", ph.CreateProductHandler).Methods("POST")
	r.HandleFunc("/api/products/batch", ph.CreateProductsBatchHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", ph.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", ph.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", ph.DeleteProductHandler).Methods("DELETE")

	// Rutas de Carrito
	r.HandleFunc("/api/cart", ch.CreateCartHandler).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", ch.GetCartHandler).Methods("GET")
	r.HandleFunc("/api/cart/{cartId}/add", ch.AddItemToCartHandler).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", ch.DeleteCartHandler).Methods("DELETE")
	r.HandleFunc("/api/cart/{cartId}/item/{productId}", ch.RemoveItemFromCartHandler).Methods("DELETE")
	r.HandleFunc("/api/cart/{cartId}/checkout", ch.CheckoutHandler).Methods("POST")

	// Ruta de Reportes
	r.HandleFunc("/api/reports/top-selling", rh.TopSellingHandler).Methods("GET")

	// Ruta de bienvenida
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a la API de E-Commerce!"))
	}).Methods("GET")
}
