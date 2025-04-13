// Package core provides the implementation of a lightweight HTTP router
// with support for dynamic routes, middleware, and route grouping.
// This file, router.go, defines the core routing logic, including the Router struct,
// route registration, middleware handling, and HTTP request dispatching.
// It also includes support for dynamic route matching and route grouping.
package core

import (
	"fmt"
	"net/http"
	"strings"
)

// Types for router.go
type HandlerFunc func(*Context)

type Router struct {
	handlers      map[string]map[string]HandlerFunc
	dynamicRoutes []dynamicRoute
	middlewares   []Middleware

	NotFoundHandler         HandlerFunc
	MethodNotAllowedHandler HandlerFunc
}

type RouteGroup struct {
	prefix      string
	router      *Router
	middlewares []Middleware
}

type dynamicRoute struct {
	method  string
	pattern string
	handler HandlerFunc
}

// Basics functions
func matchPattern(pattern, path string) (map[string]string, bool) {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	params := make(map[string]string)

	for i := 0; i < len(patternParts); i++ {
		if i >= len(pathParts) {
			return nil, false
		}

		pp := patternParts[i]
		pv := pathParts[i]

		if strings.HasPrefix(pp, ":") {
			params[pp[1:]] = pv
		} else if strings.HasPrefix(pp, "*") {
			params[pp[1:]] = strings.Join(pathParts[i:], "/")
			return params, true
		} else if pp != pv {
			return nil, false
		}
	}

	if len(pathParts) != len(patternParts) {
		return nil, false
	}

	return params, true
}

func matchRoute(route, path string) (bool, map[string]string) {
	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], ":") {
			key := strings.TrimPrefix(routeParts[i], ":")
			params[key] = pathParts[i]
		} else if routeParts[i] != pathParts[i] {
			return false, nil
		}
	}

	return true, params
}

func (r *Router) PrintRoutes() {
	fmt.Println("Registered routes:")
	for method, paths := range r.handlers {
		for path := range paths {
			fmt.Printf("%s\t%s\n", method, path)
		}
	}

	for _, dr := range r.dynamicRoutes {
		fmt.Printf("%s\t%s\n", dr.method, dr.pattern)
	}
}

// Router functions
func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(method string, path string, handler HandlerFunc) {
	if strings.Contains(path, ":") {
		r.dynamicRoutes = append(r.dynamicRoutes, dynamicRoute{
			method:  method,
			pattern: path,
			handler: handler,
		})
		return
	}

	if r.handlers[method] == nil {
		r.handlers[method] = make(map[string]HandlerFunc)
	}
	r.handlers[method][path] = handler
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	r.dynamicRoutes = append(r.dynamicRoutes, dynamicRoute{
		method:  method,
		pattern: path,
		handler: handler,
	})
}

func (r *Router) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

func (r *Router) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		prefix: prefix,
		router: r,
	}
}

func (r *Router) ServeStatic(prefix string, dir string) {
	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	r.Handle("GET", prefix+"/*filepath", func(c *Context) {
		fs.ServeHTTP(c.Writer, c.Request)
	})
}

// RouteGroup functions
func (g *RouteGroup) handle(method string, path string, handler HandlerFunc) {
	fullHandler := handler
	for i := len(g.middlewares) - 1; i >= 0; i-- {
		fullHandler = g.middlewares[i](fullHandler)
	}
	g.router.addRoute(method, g.prefix+path, fullHandler)
}

func (g *RouteGroup) Use(mw ...Middleware) {
	g.middlewares = append(g.middlewares, mw...)
}

func (g *RouteGroup) GET(path string, handler HandlerFunc) {
	g.handle("GET", path, handler)
}

func (g *RouteGroup) POST(path string, handler HandlerFunc) {
	g.handle("POST", path, handler)
}

func (g *RouteGroup) Group(subPath string) *RouteGroup {
	return &RouteGroup{
		prefix:      g.prefix + subPath,
		router:      g.router,
		middlewares: append([]Middleware{}, g.middlewares...),
	}
}

func (g *RouteGroup) ServeStatic(path string, dir string) {
	g.router.ServeStatic(g.prefix+path, dir)
}

// Main function !
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	if handler, ok := r.handlers[method][path]; ok {
		ctx := NewContext(w, req)
		for i := len(r.middlewares) - 1; i >= 0; i-- {
			handler = r.middlewares[i](handler)
		}

		handler(ctx)
		return
	}

	for _, route := range r.dynamicRoutes {
		if route.method != method {
			continue
		}

		params, ok := matchPattern(route.pattern, path)
		if ok {
			ctx := NewContext(w, req)
			ctx.Params = params

			// middlewares
			handler := route.handler
			for i := len(r.middlewares) - 1; i >= 0; i-- {
				handler = r.middlewares[i](handler)
			}

			handler(ctx)
			return
		}
	}

	for m, routes := range r.handlers {
		if m != method {
			if _, ok := routes[path]; ok && r.MethodNotAllowedHandler != nil {
				ctx := NewContext(w, req)
				r.MethodNotAllowedHandler(ctx)
				return
			}
		}
	}

	if r.NotFoundHandler != nil {
		ctx := NewContext(w, req)
		r.NotFoundHandler(ctx)
		return
	}

	http.NotFound(w, req)
}
