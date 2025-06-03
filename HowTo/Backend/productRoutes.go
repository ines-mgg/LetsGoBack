package main

import (
	"strconv"

	context "github.com/ines-mgg/LetsGoBack/Context"
	helpers "github.com/ines-mgg/LetsGoBack/Helpers"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
	router "github.com/ines-mgg/LetsGoBack/Router"

	"github.com/golang-jwt/jwt/v5"
)

// Helper function to return the user role from JWT claims
func getUserRole(c *context.Context) {
	val, ok := c.Get("jwtClaims") // Retrieve data from JWT claims
	if !ok {
		c.ErrorUnauthorized("Unauthorized access")
		return
	}
	claims, ok := val.(jwt.MapClaims)
	if !ok {
		c.ErrorUnauthorized("Invalid claims type")
		return
	}
	// Check if the user has permission to create products
	role, ok := claims["role"].(string)
	if !ok || role != RoleAdmin {
		c.ErrorUnauthorized("Unauthorized access")
		return
	}
}

func ProductRoutes(r *router.Router) {
	// Public routes
	r.GET("/products", func(c *context.Context) {
		c.RespondOK(Products)
	})
	// Dynamic route
	r.GET("/products/:id", func(c *context.Context) {
		id := c.Params["id"]
		for _, product := range Products {
			if id == strconv.Itoa(product.ID) {
				c.RespondOK(product)
				return
			}
		}
		c.ErrorNotFound("Product not found")
	})
	// Protected routes with JWT middleware
	adminProductRoutes := r.Group("/products")
	// We declare the JWT claims key as "jwtClaims", which means that the JWT claims will be stored in the context with this key.
	// This allows us to access the claims in the handlers using c.Get("jwtClaims").
	// The JWTAuthMiddleware will validate the JWT token and store the claims in the context.
	// You can change the key name to whatever you prefer, but make sure to use the same key when retrieving the claims in your handlers.
	// Identifying the key as "jwtClaims" is a common practice, but you can use any name that makes sense for your application.
	// Ideally, you should choose a key name that clearly indicates the purpose of the data it holds, such as "jwtClaims" for JWT claims or "userClaims" for user-related claims.
	// If you use the same key in multiple places, it will help maintain consistency and clarity in your code and the sore will be reset to the JWT claims.
	adminProductRoutes.Use(middleware.JWTAuthMiddleware("jwtClaims"))
	adminProductRoutes.POST("", func(c *context.Context) {
		getUserRole(c) // Check if the user has permission to create products
		var req ProductRequest
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		// Validate required fields
		if !helpers.IsNotEmpty(req.Name) {
			c.ErrorBadRequest("Name is required")
			return
		}
		if !helpers.IsNotEmpty(req.Description) {
			c.ErrorBadRequest("Description is required")
			return
		}
		if !helpers.IsGreaterThanFloat(req.Price, 0) {
			c.ErrorBadRequest("Price must be greater than 0")
			return
		}
		if !helpers.IsGreaterThanInt(req.Stock, 0) {
			c.ErrorBadRequest("Stock must be greater than or equal to 0")
			return
		}
		newProduct := Product{
			ID:          len(Products) + 1,
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Stock:       req.Stock,
		}
		Products = append(Products, newProduct)
		c.RespondCreated(map[string]string{"message": "Product created successfully"})
	})
	adminProductRoutes.PUT("/:id", func(c *context.Context) {
		getUserRole(c)
		id := c.Params["id"]
		var req ProductRequest
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		// Validate required fields
		if !helpers.IsNotEmpty(req.Name) {
			c.ErrorBadRequest("Name is required")
			return
		}
		if !helpers.IsNotEmpty(req.Description) {
			c.ErrorBadRequest("Description is required")
			return
		}
		if !helpers.IsGreaterThanFloat(req.Price, 0) {
			c.ErrorBadRequest("Price must be greater than 0")
			return
		}
		if !helpers.IsGreaterThanInt(req.Stock, 0) {
			c.ErrorBadRequest("Stock must be greater than or equal to 0")
			return
		}
		for i, product := range Products {
			if id == strconv.Itoa(product.ID) {
				// Update product details
				Products[i].Name = req.Name
				Products[i].Description = req.Description
				Products[i].Price = req.Price
				Products[i].Stock = req.Stock
				c.RespondOK(map[string]string{"message": "Product updated successfully"})
				return
			}
		}
		c.ErrorNotFound("Product not found")
	})
	// Add a specific middleware | Here we send only multiple files
	adminProductRoutes.PUT("/:id/pictures", middleware.UploadValidatorMiddleware(middleware.UploadValidationOptions{
		Field:        "files",
		Multiple:     true,
		MaxFileSize:  5 * 1024 * 1024, // 5MB
		AllowedMIMEs: []string{"image/jpeg", "image/png"},
		MaxMemory:    10 * 1024 * 1024, // 10MB
	})(
		func(c *context.Context) {
			getUserRole(c)
			id := c.Params["id"]
			files, ok := c.Get("uploadedFiles")
			if ok {
				uploadedFiles, ok := files.([]*context.UploadedFile)
				if ok {
					for _, uploadedFile := range uploadedFiles {
						for i, product := range Products {
							if id == strconv.Itoa(product.ID) {
								err := c.SaveFile(uploadedFile, "./uploads/productsPictures/"+uploadedFile.Filename)
								if err != nil {
									c.ErrorInternalServerError("Failed to save product picture")
									return
								}
								// Update product pictures
								Products[i].Pictures = append(Products[i].Pictures, uploadedFile.Filename)
							}
						}
					}
				}
				c.RespondOK(map[string]string{"message": "Product pictures updated successfully"})
				return
			}
			c.ErrorBadRequest("No files uploaded")
		}))
	adminProductRoutes.DELETE("/:id", func(c *context.Context) {
		getUserRole(c)
		id := c.Params["id"]
		for i, product := range Products {
			if id == strconv.Itoa(product.ID) {
				// Check if product has pictures and delete them
				if len(product.Pictures) > 0 {
					for _, picture := range product.Pictures {
						err := c.DeleteFile("./uploads/productsPictures/" + picture)
						if err != nil {
							c.ErrorInternalServerError("Failed to delete product picture")
							return
						}
					}
				}
				// Remove product from the list
				Products = append(Products[:i], Products[i+1:]...)
				c.RespondNoContent(map[string]string{"message": "Product deleted successfully"})
				return
			}
		}
		c.ErrorNotFound("Product not found")
	})
}
