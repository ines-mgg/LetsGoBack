package Router

import (
	context "lets-go-back/Context"
	middleware "lets-go-back/Middleware"
)

type DynamicRoute struct {
	Method  string
	Pattern string
	Handler context.HandlerFunc
}

type Router struct {
	Handlers      map[string]map[string]context.HandlerFunc
	DynamicRoutes []DynamicRoute
	Middlewares   []middleware.Middleware

	NotFoundHandler         context.HandlerFunc
	MethodNotAllowedHandler context.HandlerFunc
}

type RouteGroup struct {
	Prefix      string
	Router      *Router
	Middlewares []middleware.Middleware
}
