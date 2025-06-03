package Router

import (
	context "github.com/ines-mgg/LetsGoBack/Context"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

// routeGroup represents a group of routes with a common prefix and shared middlewares.
// It contains the prefix for the group, a reference to the parent router,
// and a slice of middlewares that will be applied to all routes in the group.
// The prefix is a string that will be prepended to all routes defined in this group.
// The router field is a pointer to the parent Router, allowing the group to add routes to it.
// The middlewares field is a slice of middleware.Middleware that can be used to apply common functionality
// to all routes in the group, such as logging, authentication, or error handling.
func (g *routeGroup) handle(method string, path string, handler context.HandlerFunc) {
	fullHandler := handler
	for i := len(g.middlewares) - 1; i >= 0; i-- {
		fullHandler = g.middlewares[i](fullHandler)
	}
	g.router.addRoute(method, g.prefix+path, fullHandler)
}

// Use adds middleware to the route group.
// The middlewares will be applied to all routes defined in this group.
func (g *routeGroup) Use(mw ...middleware.Middleware) {
	g.middlewares = append(g.middlewares, mw...)
}

// GET registers a GET route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this route group.
// The path is relative to the group's prefix, allowing for organized route management.
func (g *routeGroup) GET(path string, handler context.HandlerFunc) {
	g.handle("GET", path, handler)
}

// POST registers a POST route with the specified path and handler.
func (g *routeGroup) POST(path string, handler context.HandlerFunc) {
	g.handle("POST", path, handler)
}

// PUT registers a PUT route with the specified path and handler.
func (g *routeGroup) PUT(path string, handler context.HandlerFunc) {
	g.handle("PUT", path, handler)
}

// PATCH registers a PATCH route with the specified path and handler.
func (g *routeGroup) PATCH(path string, handler context.HandlerFunc) {
	g.handle("PATCH", path, handler)
}

// DELETE registers a DELETE route with the specified path and handler.
func (g *routeGroup) DELETE(path string, handler context.HandlerFunc) {
	g.handle("DELETE", path, handler)
}

// Group creates a new route group with a sub-path.
// The sub-path is appended to the current group's prefix, allowing for nested route groups.
func (g *routeGroup) Group(subPath string) *routeGroup {
	return &routeGroup{
		prefix:      g.prefix + subPath,
		router:      g.router,
		middlewares: append([]middleware.Middleware{}, g.middlewares...),
	}
}

// ServeStatic serves static files from the specified directory.
// The path is relative to the group's prefix, allowing for organized static file serving.
// The directory is the file system path where the static files are located.
func (g *routeGroup) ServeStatic(path string, dir string) {
	g.router.ServeStatic(g.prefix+path, dir)
}
