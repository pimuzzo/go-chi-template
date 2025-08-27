# go-chi-template

A simple and idiomatic REST API boilerplate in Go using [Chi](https://github.com/go-chi/chi), designed for clarity, modularity and extensibility.  
Includes Swagger documentation, basic validation and a clean separation of concerns.

## ✨ Features

- 🧱 Modular project structure (`cmd/`, `internal/`)
- 🔀 Routing and middleware with `chi`
- 📄 Swagger UI documentation (`/docs`)
- ✅ Request validation with `go-playground/validator`
- 🧪 Unit-test friendly structure (in progress)
- 🐳 Minimal Docker support
- 🧰 Designed as a reusable starter template

## 📦 Technologies

- Go 1.22+
- chi v5
- logrus
- swaggo/swag for Swagger
- go-playground/validator

## 🚀 Getting Started

```bash
# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
swag init -g cmd/server/main.go

# Run the server
go run ./cmd/server
