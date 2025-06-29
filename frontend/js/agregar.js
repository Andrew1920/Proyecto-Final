// Lógica para el formulario de añadir producto.
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('add-product-form');
    const statusMessage = document.getElementById('status-message');
    form.addEventListener('submit', async (event) => {
        event.preventDefault(); // Evita que la página se recargue.
        // Recolecta los datos del formulario.
        const productData = {
            name: form.name.value,
            description: form.description.value,
            price: parseFloat(form.price.value),
            stock: parseInt(form.stock.value, 10)
        };
        const apiUrl = 'http://localhost:8080/api/products';
        try {
            // Envía los datos a la API.
            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(productData),
            });
            if (!response.ok) { throw new Error(`Error HTTP: ${response.status}`); }
            // Muestra mensaje de éxito y limpia el formulario.
            statusMessage.textContent = '¡Producto añadido con éxito!';
            statusMessage.className = 'success';
            form.reset();
        } catch (error) {
            console.error('Error al añadir el producto:', error);
            statusMessage.textContent = `Error al añadir el producto.`;
            statusMessage.className = 'error';
        }
    });
});