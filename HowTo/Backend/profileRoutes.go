package main

import (
	context "github.com/ines-mgg/LetsGoBack/Context"
	helpers "github.com/ines-mgg/LetsGoBack/Helpers"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
	router "github.com/ines-mgg/LetsGoBack/Router"

	"github.com/golang-jwt/jwt/v5"
)

// Helper function to return the user ID from JWT claims
func getUserId(c *context.Context) int {
	val, ok := c.Get("jwtClaims") // Retrieve data from JWT claims
	if !ok {
		c.ErrorUnauthorized("Unauthorized access")
		return 0
	}
	claims, ok := val.(jwt.MapClaims)
	if !ok {
		c.ErrorUnauthorized("Invalid claims type")
		return 0
	}
	userIDFloat, ok := claims["id"].(float64) // JWT claims typically use float64 for numeric values
	if !ok {
		c.ErrorUnauthorized("Invalid user ID in claims")
		return 0
	}
	userID := int(userIDFloat)
	return userID
}

func ProfileRoutes(r *router.Router) {
	// Grouped routes
	profile := r.Group("/profile")
	profile.Use(middleware.JWTAuthMiddleware("jwtClaims")) // Use JWT middleware to protect these routes
	profile.GET("", func(c *context.Context) {
		userID := getUserId(c)
		if userID == 0 {
			return // Unauthorized access handled in getUserId
		}
		for _, user := range Users {
			if user.ID == userID {
				c.RespondOK(user)
				return
			}
		}
		c.ErrorNotFound("User not found")
	})
	profile.PUT("", func(c *context.Context) {
		var req ProfileRequest
		userID := getUserId(c)
		if userID == 0 {
			return // Unauthorized access handled in getUserId
		}
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		// Validate required fields
		if !helpers.IsNotEmpty(req.Username) || !helpers.IsNotEmpty(req.Email) {
			c.ErrorBadRequest("Username and email are required")
			return
		}
		if !helpers.IsValidEmail(req.Email) {
			c.ErrorBadRequest("Invalid email format")
			return
		}
		for i, user := range Users {
			if user.ID == userID {
				// Update user details
				Users[i].Username = req.Username
				Users[i].Email = req.Email
				c.RespondOK(map[string]string{"message": "Profile updated successfully"})
				return
			}
		}
		c.ErrorNotFound("User not found")
	})
	profile.PUT("/change-password", func(c *context.Context) {
		var req ChangePasswordRequest
		userID := getUserId(c)
		if userID == 0 {
			return // Unauthorized access handled in getUserId
		}
		if err := c.BindJSON(&req); err != nil {
			c.ErrorBadRequest("Invalid request body")
			return
		}
		// Validate required fields
		if !helpers.IsNotEmpty(req.OldPassword) || !helpers.IsNotEmpty(req.NewPassword) {
			c.ErrorBadRequest("Old password and new password are required")
			return
		}
		if !helpers.IsValidPassword(req.NewPassword, 128) {
			c.ErrorBadRequest("New password must be between 8 and 128 characters, and contain uppercase, lowercase, digit, and special character")
			return
		}
		for i, user := range Users {
			if user.ID == userID {
				// Check if old password matches
				if user.Password != req.OldPassword {
					c.ErrorUnauthorized("Old password is incorrect")
					return
				}
				// Update password
				Users[i].Password = req.NewPassword
				c.RespondOK(map[string]string{"message": "Password changed successfully"})
				return
			}
		}
		c.ErrorNotFound("User not found")
	})
	// Add a specific middleware | Here we send only one file
	profile.PUT("/profile-picture",
		middleware.UploadValidatorMiddleware(middleware.UploadValidationOptions{
			Field:        "file",
			Multiple:     false,
			MaxFileSize:  5 * 1024 * 1024, // 5MB
			AllowedMIMEs: []string{"image/jpeg", "image/png"},
			MaxMemory:    10 * 1024 * 1024, // 10MB
		})(
			func(c *context.Context) {
				userID := getUserId(c)
				if userID == 0 {
					return // Unauthorized access handled in getUserId
				}
				// Upload profile picture
				file, ok := c.Get("uploadedFile")
				if ok {
					uploadedFile, ok := file.(*context.UploadedFile)
					if ok {
						for i, user := range Users {
							if user.ID == userID {
								// Delete old profile picture file if it exists
								if user.ProfilPic != "" {
									err := c.DeleteFile("./uploads/profilePictures/" + user.ProfilPic)
									if err != nil {
										c.ErrorInternalServerError("Failed to delete old profile picture")
										return
									}
								}
								// Save the uploaded file to the uploads directory
								err := c.SaveFile(uploadedFile, "./uploads/profilePictures/"+uploadedFile.Filename)
								if err != nil {
									c.ErrorInternalServerError("Failed to save profile picture")
									return
								}
								// Update user profile picture
								Users[i].ProfilPic = uploadedFile.Filename
								c.RespondOK(map[string]string{"message": "Profile picture updated successfully"})
								return
							}
						}
					}
					c.ErrorNotFound("User not found")
					return
				}
				c.ErrorBadRequest("No file uploaded")
			},
		),
	)
	profile.DELETE("", func(c *context.Context) {
		userID := getUserId(c)
		if userID == 0 {
			return // Unauthorized access handled in getUserId
		}
		for i, user := range Users {
			if user.ID == userID {
				// Remove profile picture file if it exists
				if user.ProfilPic != "" {
					err := c.DeleteFile("./uploads/profilePictures/" + user.ProfilPic)
					if err != nil {
						c.ErrorInternalServerError("Failed to delete profile picture")
						return
					}
				}
				// Delete user
				Users = append(Users[:i], Users[i+1:]...)
				c.RespondNoContent(map[string]string{"message": "User deleted successfully"})
				return
			}
		}
		c.ErrorNotFound("User not found")
	})
}
