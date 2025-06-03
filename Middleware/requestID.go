package Middleware

import (
	context "github.com/ines-mgg/LetsGoBack/Context"

	"github.com/google/uuid"
)

// RequestIDMiddleware generates a unique request ID for each incoming request.
// It checks for an existing "X-Request-ID" header in the request.
// If the header is not present, it generates a new UUID and sets it in the response header.
// The request ID is also stored in the context for later use.
// This middleware is useful for tracking requests across distributed systems or logging.
// It ensures that each request can be uniquely identified, which is helpful for debugging and monitoring purposes.
func RequestIDMiddleware() Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			requestID := c.Request.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			c.Writer.Header().Set("X-Request-ID", requestID)
			c.Set("request_id", requestID)

			next(c)
		}
	}
}
