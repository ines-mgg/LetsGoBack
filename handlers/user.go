package handlers

import (
	"lets-go-back/core"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Hello(c *core.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "Hello from your framework!"})
}

func GetUser(c *core.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, map[string]string{"user_id": id})
}

func CreateUser(c *core.Context) {
	var req CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Name == "" {
		c.ErrorBadRequest("Name is required")
		return
	}

	if req.Age < 0 {
		c.ErrorBadRequest("Age must be >= 0")
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created",
		"user":    req,
	})
}

func Crash(c *core.Context) {
    panic("Something went terribly wrong 😱")
}
