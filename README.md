# Go API Project Structure

This repository follows a clean, layered architecture for a Go REST API application. This structure is designed to promote separation of concerns, testability, and maintainability.

## Project Structure
```
Project Root/
├── cmd/                    (Command line applications)
│   └── api/                (API server executable)
│       └── main.go         (Application entry point - initializes and starts the server)
├── internal/               (Private application code - not importable by external packages)
│   ├── api/                (API-specific code)
│   │   ├── routes/         (HTTP route definitions)
│   │   │   ├── user_routes.go     (User endpoints: GET /users, POST /users, etc.)
│   │   │   ├── product_routes.go  (Product endpoints: GET /products, etc.)
│   │   │   └── router.go          (Central router configuration - combines all routes)
│   │   └── middleware/     (HTTP middleware functions)
│   │       ├── auth.go            (Authentication/Authorization checks)
│   │       ├── cors.go            (Cross-Origin settings for frontend access)
│   │       └── logging.go         (Request logging and tracing)
│   ├── handlers/           (HTTP request handlers - processes HTTP requests)
│   │   ├── user_handler.go        (Functions handling user-related requests)
│   │   └── product_handler.go     (Functions handling product-related requests)
│   ├── models/             (Data structures representing domain objects)
│   │   ├── user.go                (User entity definition with fields and validation)
│   │   └── product.go             (Product entity definition)
│   ├── repositories/       (Database access layer)
│   │   ├── user_repo.go           (User CRUD operations in the database)
│   │   └── product_repo.go        (Product CRUD operations in the database)
│   ├── services/           (Business logic layer)
│   │   ├── user_service.go        (User-related business rules and operations)
│   │   └── product_service.go     (Product-related business rules and operations)
│   └── config/             (Application configuration processing)
│       └── config.go              (Code to load and validate app configuration)
├── pkg/                    (Reusable packages that could be used by external applications)
├── configs/                (Configuration files)
│   ├── app.env                    (Environment-specific variables)
│   └── app.yaml                   (Application settings in YAML format)
└── docs/                   (API documentation)
    └── swagger.yaml               (OpenAPI/Swagger API specification)
```


# Understanding the Layers

### Routes vs Handlers

- **Routes** (`internal/api/routes/`): Define URL patterns and HTTP methods your API responds to. They map each endpoint to a specific handler function. Think of routes as the "address" of your API endpoints.

- **Handlers** (`internal/handlers/`): Contain the implementation for processing HTTP requests. They extract data from requests, validate it, call the appropriate services, and format the HTTP response. Handlers bridge the HTTP world with your application's business logic.

Example comparison:
```go
// In routes/user_routes.go
router.GET("/users/:id", userHandler.GetUserByID)  // Defines the URL pattern

// In handlers/user_handler.go
func (h *UserHandler) GetUserByID(c *gin.Context) {
    id := c.Param("id")
    user, err := h.userService.GetUserByID(id) // Invoking the service layer
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}
```

### Services

- **Services** (`internal/services/`): Form the heart of your application containing core business logic. They:
    - Implement domain-specific rules and workflows
    - Are independent of HTTP or other delivery mechanisms
    - Coordinate between different repositories when needed
    - Handle complex operations that span multiple data sources
    - Enforce business constraints and validation rules

Example:
```go
// In services/user_service.go
func (s *UserService) GetUserByID(id string) (*User, error) {
    user, err := s.userRepo.FindByID(id)
    if err != nil {
        return nil, err
    }
    return user, nil
}
```