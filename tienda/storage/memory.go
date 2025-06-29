package storage

import (
	"fmt"
	"sync"
	"tienda/models"

	"github.com/google/uuid"
)

// MemoryStore implementa todas las interfaces de almacenamiento en memoria.
type MemoryStore struct {
	productsData    map[string]models.Product
	cartsData       map[string]models.Cart
	usersData       map[string]models.User
	completedOrders []models.Cart
	mutex           sync.Mutex // Previene errores de concurrencia al modificar los mapas.
}

// NewMemoryStore es el constructor para crear nuestro almacén.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		productsData:    make(map[string]models.Product),
		cartsData:       make(map[string]models.Cart),
		usersData:       make(map[string]models.User),
		completedOrders: []models.Cart{},
	}
}

// --- MÉTODOS PARA PRODUCTOS ---
// El patrón Lock/Unlock se repite en todos los métodos para seguridad en la concurrencia.
func (s *MemoryStore) GetProducts() ([]models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	list := make([]models.Product, 0, len(s.productsData))
	for _, prod := range s.productsData {
		list = append(list, prod)
	}
	return list, nil
}
func (s *MemoryStore) GetProductByID(id string) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	p, ok := s.productsData[id]
	if !ok {
		return models.Product{}, fmt.Errorf("producto con id %s no encontrado", id)
	}
	return p, nil
}
func (s *MemoryStore) CreateProduct(p models.Product) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	p.ID = uuid.NewString()
	s.productsData[p.ID] = p
	return p, nil
}
func (s *MemoryStore) UpdateProduct(id string, p models.Product) (models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.productsData[id]; !ok {
		return models.Product{}, fmt.Errorf("producto no encontrado para actualizar")
	}
	p.ID = id
	s.productsData[id] = p
	return p, nil
}
func (s *MemoryStore) DeleteProduct(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.productsData[id]; !ok {
		return fmt.Errorf("producto no encontrado para eliminar")
	}
	delete(s.productsData, id)
	return nil
}
func (s *MemoryStore) CreateBatchProducts(products []models.Product) ([]models.Product, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	created := make([]models.Product, 0)
	for _, p := range products {
		p.ID = uuid.NewString()
		s.productsData[p.ID] = p
		created = append(created, p)
	}
	return created, nil
}

// --- MÉTODOS PARA CARRITOS ---
func (s *MemoryStore) GetCartByID(id string) (models.Cart, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	c, ok := s.cartsData[id]
	if !ok {
		return models.Cart{}, fmt.Errorf("carrito con id %s no encontrado", id)
	}
	return c, nil
}
func (s *MemoryStore) CreateCart(c models.Cart) (models.Cart, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	c.ID = uuid.NewString()
	s.cartsData[c.ID] = c
	return c, nil
}
func (s *MemoryStore) UpdateCart(id string, c models.Cart) (models.Cart, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.cartsData[id]; !ok {
		return models.Cart{}, fmt.Errorf("carrito no encontrado para actualizar")
	}
	c.ID = id
	s.cartsData[id] = c
	return c, nil
}
func (s *MemoryStore) DeleteCart(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.cartsData[id]; !ok {
		return fmt.Errorf("carrito no encontrado para eliminar")
	}
	delete(s.cartsData, id)
	return nil
}

// --- MÉTODOS PARA USUARIOS ---
func (s *MemoryStore) CreateUser(u models.User) (models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exists := s.usersData[u.Username]; exists {
		return models.User{}, fmt.Errorf("el usuario '%s' ya existe", u.Username)
	}
	u.ID = uuid.NewString()
	s.usersData[u.Username] = u
	return u, nil
}
func (s *MemoryStore) GetUserByUsername(username string) (models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user, ok := s.usersData[username]
	if !ok {
		return models.User{}, fmt.Errorf("usuario '%s' no encontrado", username)
	}
	return user, nil
}

// --- MÉTODOS PARA ÓRDENES ---
func (s *MemoryStore) CreateOrderFromCart(c models.Cart) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.completedOrders = append(s.completedOrders, c)
	return nil
}
func (s *MemoryStore) GetAllOrders() ([]models.Cart, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	ordersCopy := make([]models.Cart, len(s.completedOrders))
	copy(ordersCopy, s.completedOrders)
	return ordersCopy, nil
}
