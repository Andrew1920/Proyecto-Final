package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"tienda/models"
	"tienda/storage"
)

// ReportHandlers depende del almacén de órdenes y productos.
type ReportHandlers struct {
	orderStore   storage.OrderStorer
	productStore storage.ProductStorer
}

// NewReportHandlers es el constructor para los handlers de reporte.
func NewReportHandlers(os storage.OrderStorer, ps storage.ProductStorer) *ReportHandlers {
	return &ReportHandlers{orderStore: os, productStore: ps}
}

// TopSellingHandler genera el reporte de los productos más vendidos.
func (h *ReportHandlers) TopSellingHandler(w http.ResponseWriter, r *http.Request) {
	// Lee del historial de órdenes completadas.
	orders, err := h.orderStore.GetAllOrders()
	if err != nil {
		http.Error(w, "Error al obtener órdenes", http.StatusInternalServerError)
		return
	}
	// Agrega las cantidades de cada producto vendido.
	productCounts := make(map[string]int)
	for _, order := range orders {
		for _, item := range order.Items {
			productCounts[item.ProductID] += item.Quantity
		}
	}
	type ReportItem struct {
		Product  models.Product `json:"product"`
		Quantity int            `json:"quantity_sold"`
	}
	reportData := make([]ReportItem, 0, len(productCounts))
	// Enriquece el reporte con los datos de cada producto.
	for productID, quantity := range productCounts {
		if product, err := h.productStore.GetProductByID(productID); err == nil {
			reportData = append(reportData, ReportItem{
				Product:  product,
				Quantity: quantity,
			})
		}
	}
	// Ordena el reporte de más a menos vendido.
	sort.Slice(reportData, func(i, j int) bool {
		return reportData[i].Quantity > reportData[j].Quantity
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reportData)
}
