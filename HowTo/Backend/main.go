package main

import (
	"log"

	context "github.com/ines-mgg/LetsGoBack/Context"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
	router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
	// Initialize the router
	log.Println("Starting server...")
	r := router.NewRouter()

	// Set custom NotFound handler
	r.NotFoundHandler = func(c *context.Context) {
		c.ErrorNotFound("Page not found")
	}
	// Set custom MethodNotAllowed handler
	r.MethodNotAllowedHandler = func(c *context.Context) {
		c.ErrorMethodNotAllowed("Method not allowed")
	}

	// Serve static files from the "uploads" directory
	r.ServeStatic("/uploads", "./uploads")

	// Set JWT secret key for authentication
	context.SetJWTSecret("your-secret-key")

	// Example of using CORS middleware
	corsHeaders := map[string]any{
		"Access-Control-Allow-Methods":     "GET, POST, PUT, PATCH, DELETE",
		"Access-Control-Allow-Headers":     "Content-Type, Authorization",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Origin":      "*", // DO NOT USE '*'! Always change to a specific domain
	}
	r.Use(middleware.CORSMiddleware(corsHeaders))

	// Example of using RequestID middleware
	r.Use(middleware.RequestIDMiddleware())

	// Example of using Logger middleware
	r.Use(middleware.LoggerMiddleware("2006-01-02 15:04:05"))

	// Example of using Recover middleware
	r.Use(middleware.RecoverMiddleware())

	// Example of using ErrorRecovery middleware
	r.Use(middleware.ErrorRecoveryMiddleware())

	// Static route
	r.GET("/", func(c *context.Context) {
		c.RespondOK("Hello World, I'm LetsGoBack Framework!")
	})

	// Panic route
	r.GET("/panic", func(c *context.Context) {
		// This will cause a panic and trigger the Recover middleware
		panic("This is a panic test")
	})

	// Example of grouped routes
	AuthRoutes(r)

	// Example of using JWT middleware...
	ProfileRoutes(r)
	// ...with dynamic routes
	ProductRoutes(r)

	r.PrintRoutes()                   // Print registered routes to console
	r.WriteRoutesToJsonFile("routes") // Write routes to JSON file

	// Start the server on port 8080
	log.Fatal(r.Listen(":8080"))

	// Alternatively, you can use a custom port from environment variable
	// port := ":8080"
	// if envPort := os.Getenv("BACKEND_PORT"); envPort != "" {
	// 	port = ":" + envPort
	// }
	// log.Fatal(r.Listen(port))
}
