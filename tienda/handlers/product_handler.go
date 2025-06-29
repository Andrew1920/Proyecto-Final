package handlers

import (
	"encoding/json"
	"net/http"
	"tienda/models"
	"tienda/storage"

	"github.com/gorilla/mux"
)

// ProductHandlers maneja la lógica de productos.
type ProductHandlers struct {
	store storage.ProductStorer
}

// NewProductHandlers es el constructor para los handlers de producto.
func NewProductHandlers(s storage.ProductStorer) *ProductHandlers {
	return &ProductHandlers{store: s}
}

// GetProductsHandler obtiene todos los productos.
func (h *ProductHandlers) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		http.Error(w, "Error interno al obtener productos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductHandler obtiene un producto por su ID.
func (h *ProductHandlers) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product, err := h.store.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateProductHandler crea un nuevo producto.
func (h *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.store.CreateProduct(product)
	if err != nil {
		http.Error(w, "Error interno al crear el producto", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// UpdateProductHandler actualiza un producto existente.
func (h *ProductHandlers) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}
	updatedProduct, err := h.store.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProductHandler elimina un producto.
func (h *ProductHandlers) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := h.store.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// CreateProductsBatchHandler crea múltiples productos a la vez.
func (h *ProductHandlers) CreateProductsBatchHandler(w http.ResponseWriter, r *http.Request) {
	var newProducts []models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProducts); err != nil {
		http.Error(w, "Datos inválidos, el formato JSON del array es incorrecto", http.StatusBadRequest)
		return
	}
	createdProducts, err := h.store.CreateBatchProducts(newProducts)
	if err != nil {
		http.Error(w, "Error interno del servidor al crear productos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProducts)
}
