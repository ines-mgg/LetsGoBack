package Middleware

import (
	context "lets-go-back/Context"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next context.HandlerFunc) context.HandlerFunc {
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
