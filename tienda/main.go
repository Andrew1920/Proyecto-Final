package main

import (
	"log"
	"net/http"
	"tienda/handlers"
	"tienda/routes"
	"tienda/storage"

	"github.com/gorilla/mux"
)

// CORSMiddleware permite la comunicación entre el frontend y el backend.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite cualquier origen
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		// Maneja la petición "pre-vuelo" de CORS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// 1. Inicializa la capa de almacenamiento (nuestra base de datos en memoria).
	store := storage.NewMemoryStore()

	// 2. Crea las instancias de los manejadores, inyectando el 'store' como dependencia.
	productHandlers := handlers.NewProductHandlers(store)
	userHandlers := handlers.NewUserHandlers(store)
	cartHandlers := handlers.NewCartHandlers(store, store, store)
	reportHandlers := handlers.NewReportHandlers(store, store)

	// 3. Crea el enrutador principal.
	r := mux.NewRouter()

	// 4. Aplica el middleware de CORS a todas las peticiones.
	r.Use(CORSMiddleware)

	// 5. Registra todas las rutas de la API.
	routes.RegisterRoutes(r, productHandlers, cartHandlers, userHandlers, reportHandlers)

	// 6. Inicia el servidor de la API.
	log.Println("🚀 Servidor API iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error al iniciar el servidor API: ", err)
	}
}
