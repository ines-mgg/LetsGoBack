package main

import (
	"lets-go-back/core"
	"lets-go-back/handlers"
	"log"
	"net/http"
)

func main() {
	r := core.NewRouter()

	// Middlewares
	r.Use(core.CORSMiddleware)
	r.Use(core.RecoverMiddleware)
	r.Use(core.LoggerMiddleware)

	// Routes
	api := r.Group("/api")
	api.GET("/ping", func(c *core.Context) {
		c.JSON(200, map[string]string{"message": "pong"})
	})
	api.GET("/hello", handlers.Hello)
	api.GET("/crash", handlers.Crash)

	users := api.Group("/users")
	users.GET("/:id", func(c *core.Context) {
		id := c.Param("id")
		c.JSON(200, map[string]string{"id": id})
	})
	users.POST("", handlers.CreateUser)

	web := r.Group("/web")

	web.ServeStatic("/static", "./public")

	web.POST("/login", func(c *core.Context) {
		userID := "123"

		token, err := core.GenerateJWT(userID)
		if err != nil {
			c.JSON(500, map[string]string{"error": "cannot generate token"})
			return
		}

		core.SetSession(c.Writer, token)
		c.Set("userID", userID)
		c.JSON(200, map[string]string{"status": "logged in"})
	})

	web.GET("/me", core.JWTAuthMiddleware(func(c *core.Context) {
		userIDValue, exists := c.Get("userID")
		log.Println(userIDValue)
		if !exists {
			c.JSON(401, map[string]string{"error": "Unauthorized"})
			return
		}
		userID := userIDValue.(string)
		c.JSON(200, map[string]string{"user_id": userID})
	}))

	web.POST("/logout", func(c *core.Context) {
		core.ClearSession(c.Writer)
		c.JSON(200, map[string]string{"status": "logged out"})
	})

	// web.POST("/upload", core.UploadLoggerMiddleware(func(c *core.Context){
	// 	uploaded, err := c.GetUploadedFile("file")
	// 	if err != nil {
	// 		c.JSON(400, map[string]string{"error": "no file uploaded"})
	// 		return
	// 	}

	// 	savePath := "./uploads/" + uploaded.Filename
	// 	if err := c.SaveFile(uploaded, savePath); err != nil {
	// 		c.JSON(500, map[string]string{"error": "failed to save file"})
	// 		return
	// 	}

	// 	c.JSON(200, map[string]string{
	// 		"status": "file uploaded",
	// 		"file":   uploaded.Filename,
	// 		"size":   fmt.Sprintf("%d", uploaded.Size),
	// 	})
	// }))

	web.POST("/upload", core.UploadValidatorMiddleware(core.UploadValidationOptions{
		Field:        "file",
		Multiple:     false,
		MaxFileSize:  5 * 1024 * 1024, // 5MB
		AllowedMIMEs: []string{"text/plain", "application/pdf", "application/octet-stream"},
	})(func(c *core.Context) {
		file, _ := c.GetUploadedFile("file")
		c.SaveFile(file, "./uploads/"+file.Filename)
		c.JSON(200, map[string]string{"status": "ok", "filename": file.Filename})
	}))

	r.POST("/upload-multiple", core.UploadValidatorMiddleware(core.UploadValidationOptions{
		Field:        "files",
		Multiple:     true,
		MaxFileSize:  5 << 20, // 5MB
		AllowedMIMEs: []string{"text/plain", "application/octet-stream"},
	})(func(c *core.Context) {
		files, err := c.GetUploadedFiles("files")
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid upload"})
			return
		}

		var uploaded []map[string]string
		var failed []map[string]string

		for _, file := range files {
			savePath := "./uploads/" + file.Filename

			if err := c.SaveFile(file, savePath); err != nil {
				failed = append(failed, map[string]string{
					"filename": file.Filename,
					"error":    "failed to save",
				})
				continue
			}

			uploaded = append(uploaded, map[string]string{
				"filename": file.Filename,
				"status":   "ok",
			})
		}

		response := map[string]any{
			"uploaded": uploaded,
		}
		if len(failed) > 0 {
			response["errors"] = failed
		}

		c.JSON(200, response)
	}))

	r.NotFoundHandler = func(c *core.Context) {
		c.JSON(404, map[string]string{"error": "Route not found"})
	}
	r.MethodNotAllowedHandler = func(c *core.Context) {
		c.JSON(405, map[string]string{"error": "Method not allowed"})
	}

	log.Println("Server listening on :8080")
	r.PrintRoutes()
	http.ListenAndServe(":8080", r)

}
