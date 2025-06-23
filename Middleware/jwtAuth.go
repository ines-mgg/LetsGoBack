package Middleware

import (
	context "github.com/ines-mgg/LetsGoBack/Context"
)

// JWTAuthMiddleware is a middleware that validates JWT tokens in the Authorization header.
// It checks if the token is present, validates it, and extracts the claims.
// If the token is valid, it stores the claims in the context for further use.
// If the token is missing or invalid, it responds with an unauthorized error.
// The `dataKeyName` parameter specifies the key under which the claims will be stored in the context.
// This middleware is useful for protecting routes that require authentication and authorization.
// It ensures that only requests with valid JWT tokens can access the protected resources.
// Usage example:
//
//	r.GET("/protected", middleware.JWTAuthMiddleware("userClaims"), func(c *context.Context) {
//	    userClaims := c.Get("userClaims")
//	    // Handle the request with the user claims
//	    c.RespondOK("Welcome, " + userClaims.Username)
//	})
func JWTAuthMiddleware(dataKeyName string) Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.ErrorUnauthorized("Authorization header is missing")
				return
			}
			claims, err := context.ValidateJWT(token)
			if err != nil {
				c.ErrorUnauthorized("Invalid or expired token")
				return
			}
			// Store the claims in the context for further use
			c.Set(dataKeyName, claims)
			next(c)
		}
	}
}
