package main

import (
	"time"

	context "github.com/ines-mgg/LetsGoBack/Context"
	helpers "github.com/ines-mgg/LetsGoBack/Helpers"
	router "github.com/ines-mgg/LetsGoBack/Router"

	"github.com/golang-jwt/jwt/v5"
)

func AuthRoutes(r *router.Router) {
	// Grouped routes
	auth := r.Group("/auth")
	auth.POST("/register", func(c *context.Context) {
		var req RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		if !helpers.IsNotEmpty(req.Username) || !helpers.IsNotEmpty(req.Email) || !helpers.IsNotEmpty(req.Password) {
			c.ErrorBadRequest("Missing required fields")
			return
		}
		if !helpers.IsValidEmail(req.Email) {
			c.ErrorBadRequest("Invalid email format")
			return
		}
		if !helpers.IsValidPassword(req.Password, 128) {
			c.ErrorBadRequest("Password must be between 8 and 128 characters, and contain uppercase, lowercase, digit, and special character")
			return
		}
		// Check if user already exists
		for _, user := range Users {
			if user.Username == req.Username || user.Email == req.Email {
				c.ErrorBadRequest("User already exists")
				return
			}
		}
		newUser := User{
			ID:        len(Users) + 1,
			Username:  req.Username,
			ProfilPic: "", // Empty by default
			Email:     req.Email,
			Password:  req.Password,
			Role:      RoleUser,
		}
		Users = append(Users, newUser)
		c.RespondCreated(map[string]string{"message": "User registered successfully"})
	})
	auth.POST("/login", func(c *context.Context) {
		var req LoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		// Validate required fields
		if !helpers.IsNotEmpty(req.Username) || !helpers.IsNotEmpty(req.Password) {
			c.ErrorBadRequest("Username and password are required")
			return
		}
		for _, user := range Users {
			if user.Username == req.Username && user.Password == req.Password {
				// Generate JWT token
				token, err := context.GenerateJWT(map[string]any{
					"id":   user.ID,
					"role": user.Role,
					"exp":  jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expiration time
					"iat":  jwt.NewNumericDate(time.Now()),                     // Issued at time
				})
				if err != nil {
					c.ErrorInternalServerError("Failed to generate token")
					return
				}
				c.RespondOK(map[string]any{"message": "Login successful", "token": token})
				return
			}
			c.ErrorNotFound("User not found")
		}
		c.ErrorUnauthorized("Invalid credentials")
	})
}
