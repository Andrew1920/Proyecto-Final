// Variable global para el ID del carrito, persistido en el navegador.
let currentCartId = localStorage.getItem('cartId');

// Al cargar la página, asegura que exista un carrito y carga los productos.
document.addEventListener('DOMContentLoaded', () => {
    if (!currentCartId) {
        createNewCart();
    }
    fetchProducts();
});

// Crea un nuevo carrito en la API y guarda su ID.
async function createNewCart() {
    try {
        const response = await fetch('http://localhost:8080/api/cart', { method: 'POST' });
        if (!response.ok) throw new Error('No se pudo crear el carrito');
        const cart = await response.json();
        currentCartId = cart.id;
        localStorage.setItem('cartId', currentCartId);
        return currentCartId;
    } catch (error) {
        console.error('Error al crear nuevo carrito:', error);
        return null;
    }
}

// Obtiene y muestra todos los productos en tarjetas.
async function fetchProducts() {
    const productListContainer = document.getElementById('product-list');
    const apiUrl = 'http://localhost:8080/api/products';
    try {
        const response = await fetch(apiUrl);
        if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);
        const products = await response.json();
        productListContainer.innerHTML = ''; // Limpia el mensaje "Cargando...".
        if (!products || products.length === 0) {
            productListContainer.innerHTML = '<p>No hay productos disponibles.</p>';
            return;
        }
        // Crea una tarjeta por cada producto.
        products.forEach(product => {
            const productCard = document.createElement('div');
            productCard.className = 'product-card';
            productCard.id = `product-${product.id}`;
            const formattedPrice = product.price.toLocaleString('es-EC', { style: 'currency', currency: 'USD' });
            productCard.innerHTML = `
                <h2>${product.name}</h2>
                <p class="description">${product.description}</p>
                <div class="price-stock-container">
                    <span class="price">${formattedPrice}</span>
                    <span class="stock">Stock: ${product.stock}</span>
                </div>
                <div class="card-buttons">
                    <button class="add-to-cart-btn">Añadir al Carrito</button>
                    <button class="delete-btn">Eliminar</button>
                </div>
            `;
            // Asigna los eventos a los botones.
            productCard.querySelector('.add-to-cart-btn').addEventListener('click', () => { addProductToCart(product.id); });
            productCard.querySelector('.delete-btn').addEventListener('click', () => { deleteProduct(product.id); });
            productListContainer.appendChild(productCard);
        });
    } catch (error) {
        console.error('Error al obtener los productos:', error);
        productListContainer.innerHTML = '<p>No se pudieron cargar los productos.</p>';
    }
}

// Lógica para añadir un producto al carrito, con auto-reparación de ID.
async function addProductToCart(productId) {
    if (!currentCartId) { await createNewCart(); } // Asegura tener un ID de carrito.
    if (!currentCartId) { alert("Error crítico: No se pudo obtener un ID de carrito."); return; }
    
    const apiUrl = `http://localhost:8080/api/cart/${currentCartId}/add`;
    const productToAdd = { productId: productId, quantity: 1 };
    try {
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(productToAdd)
        });
        if (!response.ok) {
            // Lógica de auto-reparación si el carrito no existe en el servidor.
            if (response.status === 404) {
                localStorage.removeItem('cartId');
                const newId = await createNewCart();
                if (newId) { addProductToCart(productId); } // Reintenta la operación.
                return;
            }
            throw new Error(`Error del servidor: ${response.status}`);
        }
        alert(`¡Producto añadido al carrito con éxito!`);
    } catch (error) {
        console.error('Error al añadir al carrito:', error);
        alert('Hubo un error al añadir el producto al carrito.');
    }
}

// Lógica para eliminar un producto de la tienda.
async function deleteProduct(productId) {
    if (!confirm('¿Estás seguro de que quieres eliminar este producto?')) return;
    const apiUrl = `http://localhost:8080/api/products/${productId}`;
    try {
        const response = await fetch(apiUrl, { method: 'DELETE' });
        if (response.ok && response.status === 204) {
            alert('Producto eliminado con éxito.');
            document.getElementById(`product-${productId}`)?.remove(); // Elimina la tarjeta de la vista.
        } else {
            throw new Error('La eliminación falló en el servidor.');
        }
    } catch (error) {
        console.error('Error al eliminar el producto:', error);
        alert('Hubo un error al eliminar el producto.');
    }
}