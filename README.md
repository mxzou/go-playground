# Go Recipe API - From First Principles

## Introduction to Go

Go (or Golang) is a statically typed, compiled programming language designed at Google. It combines the efficiency of a compiled language with the ease of use of a dynamic language. This project demonstrates Go's core principles and features through a practical, production-ready REST API for managing recipes.

### Key Go Principles

1. **Simplicity**: Go emphasizes simplicity and readability. The language has a small set of keywords and a clean syntax.
2. **Concurrency**: Go has built-in support for concurrent programming through goroutines and channels.
3. **Strong Typing**: Go is statically typed, which helps catch errors at compile time rather than runtime.
4. **Memory Safety**: Go includes garbage collection, making memory management easier and safer.
5. **Standard Library**: Go has a rich standard library that provides essential functionality without external dependencies.

## Application Architecture

This application follows a clean, layered architecture that separates concerns and promotes maintainability:

### Layers

1. **Models**: Define the data structures and business rules.
2. **Repositories**: Handle data persistence and retrieval.
3. **Services**: Implement business logic and orchestrate operations.
4. **Handlers**: Manage HTTP requests and responses.
5. **Middleware**: Provide cross-cutting concerns like authentication and logging.

### Design Patterns

- **Dependency Injection**: Components receive their dependencies rather than creating them.
- **Repository Pattern**: Abstracts data access behind interfaces.
- **Middleware Pattern**: Processes requests before they reach handlers.

## Features

- User authentication with JWT
- CRUD operations for recipes
- Recipe search by ingredients, tags, and title
- Recipe ratings and reviews
- Pagination support

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Running the Application

```bash
# Clone the repository
git clone <repository-url>
cd recipe-api

# Run the application
go run main.go
```

The server will start on http://localhost:8080

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and get JWT token

### Recipes

- `GET /api/recipes` - Get all recipes
- `GET /api/recipes/{id}` - Get recipe by ID
- `POST /api/recipes` - Create a new recipe
- `PUT /api/recipes/{id}` - Update a recipe
- `DELETE /api/recipes/{id}` - Delete a recipe

### Search

- `GET /api/search/ingredient?q={ingredient}` - Search recipes by ingredient
- `GET /api/search/tag?q={tag}` - Search recipes by tag
- `GET /api/search/title?q={title}` - Search recipes by title

### Ratings

- `GET /api/recipes/{id}/ratings` - Get ratings for a recipe
- `POST /api/recipes/{id}/ratings` - Add a rating to a recipe

## Code Structure

```
.
├── handlers/       # HTTP request handlers
├── middleware/     # HTTP middleware components
├── models/         # Data structures and business rules
├── repositories/   # Data access layer
├── services/       # Business logic layer
└── main.go         # Application entry point
```

## Testing

The application includes both unit tests and end-to-end tests to ensure reliability and correctness.

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover
```

## Design Decisions

### In-Memory Storage

For simplicity, this application uses in-memory storage for data persistence. In a production environment, you would replace the repository implementations with database-backed versions.

### JWT Authentication

JSON Web Tokens (JWT) are used for authentication because they are stateless and can scale horizontally without shared session storage.

### Middleware Approach

The application uses middleware for cross-cutting concerns like logging, recovery from panics, and authentication. This keeps the handlers focused on their primary responsibilities.

## Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)