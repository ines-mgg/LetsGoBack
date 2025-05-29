package main

import (
	context "lets-go-back/Context"
	//middleware "lets-go-back/Middleware"
	router "lets-go-back/Router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	r.GET("/", func(c *context.Context) {
		c.JSON(200, map[string]string{"message": "pong"})
	})

	api := r.Group("/api")
	api.GET("/ping", func(c *context.Context) {
		c.JSON(200, map[string]string{"message": "pong"})
	})

	log.Println("Server listening on :8080")
	r.PrintRoutes()
	http.ListenAndServe(":8080", r)
}
