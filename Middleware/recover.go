package Middleware

import (
	"fmt"
	context "lets-go-back/Context"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next context.HandlerFunc) context.HandlerFunc {
	return func(c *context.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorID := context.GenerateErrorID()
				log.Printf("[PANIC][%s] %s %s - %v\n%s", errorID, c.Request.Method, c.Request.URL.Path, err, string(debug.Stack()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error. Error ID: %s", errorID))
			}
		}()
		next(c)
	}
}