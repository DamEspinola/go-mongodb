# Go Gin Image Store

This project is a simple web application built using the Gin framework that allows users to upload and retrieve images. The images are stored in a MongoDB database.

## Project Structure

```go-gin-image-store/
├── cmd/
│   └── server/
│       └── main.go             # Punto de entrada de la aplicación
├── internal/
│   ├── config/
│   │   └── config.go           # Configuración y variables de entorno
│   ├── delivery/
│   │   └── http/
│   │       ├── http.go         # Configuración de rutas
│   │       └── image_handler.go # Handlers para subir y servir imágenes
│   ├── domain/
│   │   └── models/
│   │       └── image.go        # Modelo de imagen
│   ├── repository/
│   │   ├── image_repository.go # Interfaz del repositorio
│   │   └── mongodb/
│   │       └── mongo_repository.go # Implementación MongoDB
│   ├── storage/
│   │   └── database.go         # Conexión a MongoDB
│   └── usercase/               # (Nota: debería ser "usecase")
│       └── image_service.go    # Lógica de negocio de imágenes
├── .vscode/
│   └── launch.json            # Configuración de depuración
├── go.mod                     # Definición de módulo y dependencias
├── go.sum                     # Checksums de dependencias
└── .env                       # Variables de entorno
```
## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-gin-image-store
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the environment variables:**
   Create a `.env` file in the root directory and add your MongoDB connection string and any other necessary environment variables.

4. **Run the application:**
   ```
   go run cmd/server/main.go
   ```
   
5. **Development with Live Reload:**
   Install Air for live reloading during development:
   ```
   go install github.com/cosmtrek/air@latest
   ```
   Then run the application using Air:
   ```
   air
   ```

## Usage

- **Upload an Image:**
  Send a POST request to `/upload` with the image file.

- **Retrieve an Image:**
  Send a GET request to `/images/:id` to retrieve an image by its ID.

## License

This project is licensed under the MIT License.