// Lógica para la página de reportes.
document.addEventListener('DOMContentLoaded', () => {
    generateReport();
});

// Obtiene y muestra el reporte de productos más vendidos.
async function generateReport() {
    const reportContainer = document.getElementById('report-container');
    const apiUrl = 'http://localhost:8080/api/reports/top-selling';
    try {
        const response = await fetch(apiUrl);
        if (!response.ok) throw new Error('No se pudo generar el reporte.');
        const reportData = await response.json();
        if (!reportData || reportData.length === 0) {
            reportContainer.innerHTML = '<p>No hay datos de ventas para generar un reporte. ¡Realiza una compra primero!</p>';
            return;
        }
        // Construye la tabla del reporte dinámicamente.
        let tableHtml = `<table><thead><tr><th>Producto</th><th>Descripción</th><th>Cantidad Vendida</th></tr></thead><tbody>`;
        reportData.forEach(item => {
            tableHtml += `
                <tr>
                    <td>${item.product.name}</td>
                    <td>${item.product.description}</td>
                    <td><strong>${item.quantity_sold}</strong></td>
                </tr>`;
        });
        tableHtml += `</tbody></table>`;
        reportContainer.innerHTML = tableHtml;
    } catch (error) {
        console.error('Error:', error);
        reportContainer.innerHTML = '<p>Error al generar el reporte.</p>';
    }
}