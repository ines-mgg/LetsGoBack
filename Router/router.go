package Router

import (
	"net/http"
	"strings"

	context "github.com/ines-mgg/LetsGoBack/Context"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

// NewRouter creates a new Router instance.
// It initializes the Handlers map and the DynamicRoutes slice.
func NewRouter() *Router {
	return &Router{
		Handlers: make(map[string]map[string]context.HandlerFunc),
	}
}

// addRoute adds a new route to the router.
// It checks if the path contains a dynamic segment (indicated by ":").
// If it does, the route is added to the DynamicRoutes slice.
func (r *Router) addRoute(method string, path string, handler context.HandlerFunc) {
	if strings.Contains(path, ":") {
		r.DynamicRoutes = append(r.DynamicRoutes, dynamicRoute{
			method,
			path,
			handler,
		})
		return
	}

	if r.Handlers[method] == nil {
		r.Handlers[method] = make(map[string]context.HandlerFunc)
	}
	r.Handlers[method][path] = handler
}

// handle is a helper function to add a dynamic route.
func (r *Router) handle(method, path string, handler context.HandlerFunc) {
	r.DynamicRoutes = append(r.DynamicRoutes, dynamicRoute{
		method,
		path,
		handler,
	})
}

// Use adds a middleware to the router.
func (r *Router) Use(m middleware.Middleware) {
	r.Middlewares = append(r.Middlewares, m)
}

// GET registers a GET route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) GET(path string, handler context.HandlerFunc) {
	r.addRoute("GET", path, handler)
}

// POST registers a POST route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) POST(path string, handler context.HandlerFunc) {
	r.addRoute("POST", path, handler)
}

// PUT registers a PUT route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) PUT(path string, handler context.HandlerFunc) {
	r.addRoute("PUT", path, handler)
}

// PATCH registers a PATCH route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) PATCH(path string, handler context.HandlerFunc) {
	r.addRoute("PATCH", path, handler)
}

// DELETE registers a DELETE route with the specified path and handler.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) DELETE(path string, handler context.HandlerFunc) {
	r.addRoute("DELETE", path, handler)
}

// Group creates a new route group with a common prefix.
func (r *Router) Group(prefix string) *routeGroup {
	return &routeGroup{
		prefix: prefix,
		router: r,
	}
}

// ServeStatic serves static files from the specified directory.
// The prefix is the URL path that will be used to access the static files.
// The directory is the file system path where the static files are located.
// It uses http.FileServer to serve the files and http.StripPrefix to remove the prefix from the file paths.
// The handler will be wrapped with the middlewares defined for this router.
func (r *Router) ServeStatic(prefix string, dir string) {
	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	r.handle("GET", prefix+"/*filepath", func(c *context.Context) {
		fs.ServeHTTP(c.Writer, c.Request)
	})
}
