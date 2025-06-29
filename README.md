# Proyecto E-Commerce: API RESTful con Go y Frontend con JS

Este repositorio contiene el código fuente de una aplicación web de comercio electrónico funcional, construida con una arquitectura desacoplada que separa el backend del frontend.

- **Autor:** Andrew Steeven Chavez Hernandez
- **Fecha de Creación:** 29 de Junio de 2025

---

## 🎯 Objetivo del Programa

El objetivo principal de este proyecto es demostrar la construcción de un sistema web moderno y escalable desde cero. Sirve como un caso práctico para aplicar conceptos clave de desarrollo de software, incluyendo:

- **Diseño de APIs RESTful:** Crear un conjunto de servicios web robustos, lógicos y estandarizados.
- **Arquitectura Desacoplada:** Separar las responsabilidades del servidor (lógica de negocio) y del cliente (interfaz de usuario).
- **Desarrollo Full-Stack:** Implementar tanto el backend como el frontend, y asegurar su correcta comunicación.
- **Manejo de Datos:** Simular la persistencia y manipulación de datos en un entorno de servidor.

---

## ✨ Funcionalidades Principales

El sistema se divide en dos componentes principales: el Backend (la API) y el Frontend (la aplicación web).

### **Backend (API en Go)**

El servidor, construido en Go, actúa como el cerebro de la aplicación y expone una serie de endpoints para gestionar todos los recursos.

-   **Gestión Completa de Productos (CRUD):**
    -   `GET /api/products`: Obtiene la lista de todos los productos.
    -   `POST /api/products`: Crea un nuevo producto.
    -   `DELETE /api/products/{id}`: Elimina un producto específico.
    -   `PUT /api/products/{id}`: Actualiza un producto existente (no implementado en el frontend, pero la API está lista).
-   **Gestión del Carrito de Compras:**
    -   `POST /api/cart`: Crea un nuevo carrito de compras para un usuario.
    -   `POST /api/cart/{cartId}/add`: Añade un producto a un carrito específico.
    -   `DELETE /api/cart/{cartId}/item/{productId}`: Elimina un ítem del carrito.
    -   `POST /api/cart/{cartId}/checkout`: Procesa la compra, convierte el carrito en una orden y lo vacía.
-   **Sistema de Autenticación de Usuarios:**
    -   `POST /register`: Registra un nuevo usuario con contraseña encriptada.
    -   `POST /login`: Valida las credenciales de un usuario.
-   **Módulo de Reportes:**
    -   `GET /api/reports/top-selling`: Genera un reporte con los productos más vendidos en base a las compras finalizadas.

### **Frontend (Aplicación Web con HTML, CSS y JavaScript)**

La interfaz de usuario, con la que interactúa el cliente, es una aplicación web limpia y responsiva.

-   **Navegación Intuitiva:** La aplicación cuenta con una barra de navegación para acceder fácilmente a todas las secciones:
    -   **Inicio:** Página de bienvenida.
    -   **Ver Productos:** Catálogo principal donde se listan todos los productos.
    -   **Añadir Producto:** Formulario para crear y agregar nuevos productos al sistema.
    -   **Ver Carrito:** Página que muestra los productos añadidos, el total, y permite finalizar la compra.
    -   **Reportes:** Visualización del reporte de los productos más vendidos.
-   **Interactividad con el Catálogo:** Desde la página de productos, un usuario puede:
    -   Añadir cualquier producto al carrito con un solo clic.
    -   Eliminar un producto del sistema (simulando una vista de administrador).
-   **Gestión Persistente del Carrito:** El ID del carrito se guarda en el `localStorage` del navegador, permitiendo que el usuario no pierda su carrito si cierra la pestaña.
-   **Feedback al Usuario:** Se muestran mensajes de estado para notificar al usuario cuando una acción (como añadir un producto o finalizar una compra) ha sido exitosa o ha fallado.

---

## 🛠️ Tecnologías Utilizadas

-   **Backend:**
    -   Lenguaje: **Go**
    -   Enrutamiento HTTP: **Gorilla Mux** (`github.com/gorilla/mux`)
    -   Generación de UUIDs: `github.com/google/uuid`
    -   Encriptación de contraseñas: `golang.org/x/crypto/bcrypt`
-   **Frontend:**
    -   Estructura: **HTML5**
    -   Estilos: **CSS3**
    -   Lógica y comunicación con API: **JavaScript (ES6+)**
-   **Formato de Datos:** **JSON** para la comunicación entre el frontend y el backend.

---

## 🚀 Instrucciones de Ejecución

Para ejecutar este proyecto, necesitas tener **Go (versión 1.18 o superior)** instalado en tu sistema.

**1. Iniciar el Servidor del Backend:**

Abre una terminal, navega a la carpeta del backend y ejecuta el siguiente comando. La API se iniciará en `http://localhost:8080`.

```bash
# Navegar a la carpeta del backend
cd tienda

# Descargar las dependencias e iniciar el servidor
go run main.go
```

**2. Iniciar el Servidor del Frontend:**

Abre una **segunda terminal** (sin cerrar la primera), navega a la carpeta del frontend y ejecuta el servidor de archivos local.

```bash
# Navegar a la carpeta del frontend
cd frontend

# Iniciar el servidor de archivos simple
go run server.go
```

**3. Abrir la Aplicación:**

Una vez que ambos servidores estén en ejecución, abre tu navegador web y ve a la siguiente dirección:
-Link de la api
**`http://localhost:8080`**
-link de la aplicacion web
**`http://localhost:8001`**
¡Listo! Ya puedes interactuar con la aplicación web completa.
