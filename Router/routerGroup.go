package Router

import (
	context "lets-go-back/Context"
	middleware "lets-go-back/Middleware"
)

func (g *RouteGroup) handle(method string, path string, handler context.HandlerFunc) {
	fullHandler := handler
	for i := len(g.Middlewares) - 1; i >= 0; i-- {
		fullHandler = g.Middlewares[i](fullHandler)
	}
	g.Router.AddRoute(method, g.Prefix+path, fullHandler)
}

func (g *RouteGroup) Use(mw ...middleware.Middleware) {
	g.Middlewares = append(g.Middlewares, mw...)
}

func (g *RouteGroup) GET(path string, handler context.HandlerFunc) {
	g.handle("GET", path, handler)
}

func (g *RouteGroup) POST(path string, handler context.HandlerFunc) {
	g.handle("POST", path, handler)
}

func (g *RouteGroup) Group(subPath string) *RouteGroup {
	return &RouteGroup{
		Prefix:      g.Prefix + subPath,
		Router:      g.Router,
		Middlewares: append([]middleware.Middleware{}, g.Middlewares...),
	}
}

func (g *RouteGroup) ServeStatic(path string, dir string) {
	g.Router.ServeStatic(g.Prefix+path, dir)
}
