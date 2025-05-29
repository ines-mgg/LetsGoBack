package Context

import (
	"fmt"
	"log"
	"net/http"
)

func (c *Context) AbortWithError(status int, err error) {
	errorID := GenerateErrorID()
	log.Printf("[ERROR][%s] %s %s - %v", errorID, c.Request.Method, c.Request.URL.Path, err)

	c.AbortWithStatusJSON(status, fmt.Sprintf("Something went wrong. Error ID: %s", errorID))
}

func (c *Context) ErrorBadRequest(msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, msg)
}
