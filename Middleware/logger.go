package Middleware

import (
	"log"
	"net/http"
	"time"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// LoggerMiddleware is a middleware that logs the request and response details.
// It logs the request method, path, status code, duration, and request ID.
// The log format is customizable through the timeFormat parameter.
// The middleware captures the start time of the request, calls the next handler,
// and then logs the details after the handler has completed.
// The log entry includes the log level based on the status code, the formatted time,
// the HTTP method, the request path, the status code, the duration of the request,
// and the request ID.
// The log level is determined by the status code, with different levels for success (2xx),
// client error (4xx), and server error (5xx) responses.
// The timeFormat parameter allows customization of the timestamp format in the log entry.
// The middleware is useful for monitoring and debugging purposes, providing insights into the performance and behavior of the application.
// It can be used in the router setup to wrap the main handler function.
// The middleware should be placed after the RequestIDMiddleware to ensure that the request ID is available for logging.
// The middleware can be used in the router setup to wrap the main handler function.
// It is important to note that this middleware does not modify the request or response,
// but only logs the details of the request and response after the handler has executed.
func LoggerMiddleware(timeFormat string) Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			start := time.Now()

			next(c)

			duration := time.Since(start)
			status := c.GetStatus()
			method := c.Method
			path := c.Path

			reqID, _ := c.Get("request_id")
			reqIDStr, _ := reqID.(string)

			if status == 0 {
				status = http.StatusOK
			}

			if reqID == "" {
				reqID = "n/a"
			}

			log.Printf("%s [%s] [%s] %s - %d (%s) - reqID: %s",
				LogLevel(status),
				time.Now().Format(timeFormat),
				method,
				path,
				status,
				duration.Truncate(time.Millisecond),
				reqIDStr,
			)
		}
	}
}
