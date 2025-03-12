# Go Gin Image Store

This project is a simple web application built using the Gin framework that allows users to upload and retrieve images. The images are stored in a MongoDB database.

## Project Structure

```
go-gin-image-store
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── config
│   │   └── config.go       # Configuration settings
│   ├── handlers
│   │   ├── image_handler.go # Handlers for image upload and retrieval
│   │   └── routes.go       # Route definitions
│   ├── middleware
│   │   └── auth.go         # Authentication middleware
│   ├── models
│   │   └── image.go        # Image model structure
│   └── storage
│       └── mongodb.go      # MongoDB connection and CRUD operations
├── pkg
│   └── utils
│       └── image.go        # Utility functions for image processing
├── go.mod                   # Module definition and dependencies
├── go.sum                   # Module dependency checksums
├── .env                     # Environment variables
├── .air.toml                # Air configuration for live reload
└── README.md                # Project documentation
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