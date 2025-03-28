package routes

import (
	"net/http"

	"littleeinsteinchildcare/backend/internal/handlers"
	// Commented out until implemented
	// "littleeinsteinchildcare/backend/internal/repositories"
	// "littleeinsteinchildcare/backend/internal/services"
)

// SetupRouter configures and returns the main router
func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Initialize repositories, services, and handlers
	// Commented out until implemented
	// userRepo := repositories.NewUserRepository()
	// userService := services.NewUserService(userRepo)

	// Create a handler without actual dependencies for now
	userHandler := handlers.NewUserHandler(nil) // Passing nil for now

	// Register routes
	RegisterUserRoutes(router, userHandler)

	// API information endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the Little Einstein Childcare API", "version": "1.0"}`))
	})

	return router
}
