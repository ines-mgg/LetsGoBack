package Middleware

import (
	context "lets-go-back/Context"
	"log"
	"net/http"
)

func ErrorRecoveryMiddleware() Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("[ERROR] [%s] Panic recovered: %v - path: %s", c.RequestID(), rec, c.Path)

					c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "internal server error",
					})
				}
			}()

			next(c)
		}
	}
}