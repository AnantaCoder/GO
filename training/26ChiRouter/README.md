# 📚 Chi Router Book API

A robust, in-memory RESTful API for managing books, built using Go and the `go-chi/chi` router. This project demonstrates best practices in Go web development, including thread-safe data access, struct-based HTTP handlers, and automatic Swagger documentation.

## ✨ Features

- **Full CRUD Operations**: Create, Read, Update, and Delete books.
- **Thread-Safe Data Store**: Uses `sync.RWMutex` to ensure safe concurrent access to the in-memory map.
- **Interactive Documentation**: Integrated with Swagger UI (`swaggo/http-swagger`) for live API testing right from the browser.
- **Dependency Injection**: Handlers receive the data store via struct injection rather than relying on global variables.
- **Validation**: Incoming JSON payloads are validated against business rules before being processed.
- **Middleware**: Utilizes Chi's built-in `Logger` and `Recoverer` middlewares for request tracking and crash prevention.

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher

### Running the Server

1. Navigate to the project root:
   ```bash
   cd d:/GO/training/26ChiRouter
   ```
2. Start the server:
   ```bash
   go run main.go
   ```
3. The server will start on port `3000`.

### Viewing the API Documentation
Once the server is running, you can explore the API and test endpoints directly by navigating to:
👉 **[http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)**

## 📂 Project Structure

```text
.
├── main.go               # Application entry point, router setup, and middleware attachment
├── models/
│   └── book.go           # Defines the Book struct and validation logic
├── store/
│   └── book_store.go     # Thread-safe in-memory map simulating a database
├── handlers/
│   └── book.go           # HTTP handlers that bridge the router and the data store
├── middleware/
│   └── auth.go           # Custom middleware implementations (e.g., Auth)
└── docs/                 # Auto-generated Swagger specifications (swagger.json/yaml)
```

## 🛣️ API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/books` | Retrieve a list of all books in the store. |
| `GET` | `/books/{id}` | Retrieve a specific book by its ID. |
| `POST` | `/books` | Create a new book. |
| `PUT` | `/books/{id}` | Update an existing book's details. |
| `DELETE`| `/books/{id}` | Remove a book from the store. |

## 🛠️ Generating Documentation

If you modify the endpoints or the `Book` model, you need to regenerate the Swagger documentation. 

Ensure you have the `swag` CLI tool installed:
```bash
go get -u github.com/swaggo/swag/cmd/swag
```

Then, run the following command from the root directory:
```bash
go run github.com/swaggo/swag/cmd/swag init
```
This updates the files in the `docs/` folder. Restart your server to see the changes in the Swagger UI.
