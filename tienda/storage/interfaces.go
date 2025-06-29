package storage

import "tienda/models"

// Storer agrupa todas las interfaces de almacenamiento para una fácil inyección.
type Storer interface {
	ProductStorer
	CartStorer
	UserStorer
	OrderStorer
}

// ProductStorer define el contrato para el almacenamiento de productos.
type ProductStorer interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	CreateProduct(p models.Product) (models.Product, error)
	UpdateProduct(id string, p models.Product) (models.Product, error)
	DeleteProduct(id string) error
	CreateBatchProducts(products []models.Product) ([]models.Product, error)
}

// CartStorer define el contrato para el almacenamiento de carritos.
type CartStorer interface {
	GetCartByID(id string) (models.Cart, error)
	CreateCart(c models.Cart) (models.Cart, error)
	UpdateCart(id string, c models.Cart) (models.Cart, error)
	DeleteCart(id string) error
}

// UserStorer define el contrato para el almacenamiento de usuarios.
type UserStorer interface {
	CreateUser(u models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

// OrderStorer define el contrato para las órdenes completadas.
type OrderStorer interface {
	CreateOrderFromCart(c models.Cart) error
	GetAllOrders() ([]models.Cart, error)
}
