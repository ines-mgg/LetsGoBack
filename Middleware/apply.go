package Middleware

import context "lets-go-back/Context"

func ApplyMiddleware(h context.HandlerFunc, middleware []Middleware) context.HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
