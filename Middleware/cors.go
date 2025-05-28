package Middleware

import (
	context "lets-go-back/Context"
	"net/http"
)

func CORSMiddleware(next context.HandlerFunc) context.HandlerFunc {
	return func(c *context.Context) {
		origin := c.Request.Header.Get("Origin")

		if origin == "http://localhost:5500" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		next(c)
	}
}