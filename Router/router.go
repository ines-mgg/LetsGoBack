package Router

import (
	context "lets-go-back/Context"
	middleware "lets-go-back/Middleware"
	"net/http"
	"strings"
)

func NewRouter() *Router {
	return &Router{
		Handlers: make(map[string]map[string]context.HandlerFunc),
	}
}

func (r *Router) AddRoute(method string, path string, handler context.HandlerFunc) {
	if strings.Contains(path, ":") {
		r.DynamicRoutes = append(r.DynamicRoutes, DynamicRoute{
			Method:  method,
			Pattern: path,
			Handler: handler,
		})
		return
	}

	if r.Handlers[method] == nil {
		r.Handlers[method] = make(map[string]context.HandlerFunc)
	}
	r.Handlers[method][path] = handler
}

func (r *Router) Handle(method, path string, handler context.HandlerFunc) {
	r.DynamicRoutes = append(r.DynamicRoutes, DynamicRoute{
		Method:  method,
		Pattern: path,
		Handler: handler,
	})
}

func (r *Router) Use(m middleware.Middleware) {
	r.Middlewares = append(r.Middlewares, m)
}

func (r *Router) GET(path string, handler context.HandlerFunc) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler context.HandlerFunc) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		Prefix: prefix,
		Router: r,
	}
}

func (r *Router) ServeStatic(prefix string, dir string) {
	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	r.Handle("GET", prefix+"/*filepath", func(c *context.Context) {
		fs.ServeHTTP(c.Writer, c.Request)
	})
}