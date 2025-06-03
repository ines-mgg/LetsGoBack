package Middleware

import (
	"fmt"
	"log"
	"runtime/debug"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// RecoverMiddleware is a middleware that recovers from panics in the application.
// It logs the panic details and returns a 500 Internal Server Error response to the client.
// This middleware is useful for preventing the application from crashing due to unexpected errors.
// It captures the panic, logs the error message along with the stack trace, and generates a unique error ID for tracking.
// The error ID is included in the response message to help identify the specific error instance.
// Usage of this middleware is recommended in production environments to ensure that the application remains stable and provides meaningful error messages to clients.
// It should be placed at the top of the middleware stack to catch any panics that occur in subsequent handlers.
// The middleware can be used in the router setup to wrap the main handler function.
// It is important to note that this middleware should not be used for handling expected errors,
// but rather for unexpected panics that could crash the application.
func RecoverMiddleware() Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			defer func() {
				if err := recover(); err != nil {
					errorID := context.GenerateErrorID()
					log.Printf("[PANIC][%s] %s %s - %v\n%s", errorID, c.Request.Method, c.Request.URL.Path, err, string(debug.Stack()))
					c.ErrorInternalServerError(fmt.Sprintf("An unexpected error occurred. Error ID: %s", errorID))
				}
			}()
			next(c)
		}
	}
}
