// Lógica para la página del carrito.
document.addEventListener('DOMContentLoaded', () => {
    loadCart();
});

// Carga y dibuja el contenido del carrito.
async function loadCart() {
    const cartContainer = document.getElementById('cart-container');
    const cartTotalElement = document.getElementById('cart-total');
    const checkoutContainer = document.getElementById('checkout-container');
    const cartId = localStorage.getItem('cartId');

    if (!cartId) {
        cartContainer.innerHTML = '<p>Tu carrito está vacío.</p>';
        return;
    }
    const apiUrl = `http://localhost:8080/api/cart/${cartId}`;
    try {
        const response = await fetch(apiUrl);
        if (response.status === 404) { // El carrito ya no existe.
            cartContainer.innerHTML = '<p>Tu carrito está vacío.</p>';
            checkoutContainer.innerHTML = '';
            cartTotalElement.innerHTML = '';
            localStorage.removeItem('cartId');
            return;
        }
        if (!response.ok) throw new Error('No se pudo cargar el carrito.');
        
        const cart = await response.json();
        if (!cart.items || cart.items.length === 0) {
            cartContainer.innerHTML = '<p>Tu carrito está vacío.</p>';
            return;
        }

        // Construye la tabla del carrito dinámicamente.
        let tableHtml = `<table><thead><tr><th>Producto</th><th>Cantidad</th><th>Precio Unitario</th><th>Subtotal</th><th>Acción</th></tr></thead><tbody>`;
        const productsResponse = await fetch('http://localhost:8080/api/products');
        const products = await productsResponse.json();
        const productMap = new Map(products.map(p => [p.id, p.name]));
        cart.items.forEach(item => {
            const productName = productMap.get(item.productId) || 'Producto no encontrado';
            const price = item.price.toLocaleString('es-EC', { style: 'currency', currency: 'USD' });
            const subtotal = (item.price * item.quantity).toLocaleString('es-EC', { style: 'currency', currency: 'USD' });
            tableHtml += `
                <tr id="item-${item.productId}">
                    <td>${productName}</td>
                    <td>${item.quantity}</td>
                    <td>${price}</td>
                    <td>${subtotal}</td>
                    <td><button class="delete-item-btn" data-product-id="${item.productId}">Eliminar</button></td>
                </tr>`;
        });
        tableHtml += `</tbody></table>`;
        cartContainer.innerHTML = tableHtml;

        // Muestra el total y el botón de comprar.
        const total = cart.total.toLocaleString('es-EC', { style: 'currency', currency: 'USD' });
        cartTotalElement.innerHTML = `<strong>Total: ${total}</strong>`;
        checkoutContainer.innerHTML = `<button id="checkout-btn" class="submit-btn">Realizar Compra</button>`;

        // Asigna eventos a los botones de eliminar y comprar.
        document.querySelectorAll('.delete-item-btn').forEach(button => {
            button.addEventListener('click', () => { removeItemFromCart(cartId, button.dataset.productId); });
        });
        document.getElementById('checkout-btn').addEventListener('click', () => { checkout(cartId); });
    } catch (error) {
        console.error('Error:', error);
        cartContainer.innerHTML = '<p>Error al cargar el carrito.</p>';
    }
}

// Llama a la API para quitar un ítem del carrito.
async function removeItemFromCart(cartId, productId) {
    if (!confirm('¿Quitar este producto del carrito?')) return;
    const apiUrl = `http://localhost:8080/api/cart/${cartId}/item/${productId}`;
    try {
        const response = await fetch(apiUrl, { method: 'DELETE' });
        if (!response.ok) throw new Error('No se pudo eliminar el producto');
        loadCart(); // Recarga el carrito para mostrar los cambios.
    } catch (error) {
        console.error('Error al eliminar:', error);
        alert('Error al eliminar el producto del carrito.');
    }
}

// Llama a la API para finalizar la compra.
async function checkout(cartId) {
    if (!confirm('¿Finalizar la compra? Esto vaciará tu carrito y registrará la venta.')) return;
    const apiUrl = `http://localhost:8080/api/cart/${cartId}/checkout`;
    try {
        const response = await fetch(apiUrl, { method: 'POST' });
        if (!response.ok) throw new Error('No se pudo procesar la compra');
        alert('¡Gracias por tu compra!');
        localStorage.removeItem('cartId'); // Limpia el carrito del navegador.
        window.location.href = 'index.html'; // Redirige al inicio.
    } catch (error) {
        console.error('Error en la compra:', error);
        alert('Hubo un error al procesar tu compra.');
    }
}