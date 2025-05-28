package Middleware

import (
	context "lets-go-back/Context"
	"log"
	"time"
)

func LoggerMiddleware(next context.HandlerFunc) context.HandlerFunc {
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
			status = 200 // fallback
		}

		if reqID == "" {
			reqID = "n/a"
		}

		log.Printf("%s [%s] [%s] %s - %d (%s) - reqID: %s",
			LogLevel(status),
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			status,
			duration.Truncate(time.Millisecond),
			reqIDStr,
		)
	}
}