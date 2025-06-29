package models

// CartItem representa un artículo dentro de un carrito.
type CartItem struct {
	ProductID string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Precio del producto al momento de añadirlo.
}

// Cart representa el carrito de compras.
type Cart struct {
	ID    string     `json:"id"`
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}
