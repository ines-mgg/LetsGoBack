package Middleware

import (
	"net/http"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// CORSMiddleware is a middleware that sets CORS headers for HTTP responses.
// It allows cross-origin requests by setting the appropriate headers.
// The headerMap parameter is a map of headers to be set in the response.
// It is typically used to enable cross-origin resource sharing (CORS) in web applications.
// The middleware checks if the request method is OPTIONS, and if so, it responds with a 200 OK status.
// For other methods, it sets the headers and calls the next handler in the chain.
// Usage example:
//
//	corsHeaders := map[string]any{
//	    "Access-Control-Allow-Origin": "*",
//	    "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
//	    "Access-Control-Allow-Headers": "Content-Type, Authorization",
//	}
//
// r.Use(middleware.CORSMiddleware(corsHeaders))
// This middleware should be used before any handlers that require CORS support.
func CORSMiddleware(headerMap map[string]any) Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			for key, value := range headerMap {
				c.Writer.Header().Set(key, value.(string))
			}

			if c.Request.Method == http.MethodOptions {
				c.Writer.WriteHeader(http.StatusOK)
				return
			}

			next(c)
		}
	}
}
