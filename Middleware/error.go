package Middleware

import (
	"log"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// ErrorRecoveryMiddleware is a middleware that recovers from panics in the application.
// It logs the panic details and returns a 500 Internal Server Error response to the client.
// This middleware is useful for preventing the application from crashing due to unexpected errors.
// It captures the panic, logs the error message along with the stack trace, and generates a unique error ID for tracking.
// The error ID is included in the response message to help identify the specific error instance.
// Usage of this middleware is recommended in production environments to ensure that the application remains stable and provides meaningful error messages to clients.
// It should be placed at the top of the middleware stack to catch any panics that occur in subsequent handlers.
func ErrorRecoveryMiddleware() Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("[ERROR] [%s] Panic recovered: %v - path: %s", c.RequestID(), rec, c.Path)
					c.ErrorInternalServerError("An unexpected error occurred")
				}
			}()

			next(c)
		}
	}
}
