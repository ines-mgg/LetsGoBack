package Router

import (
	"net/http"
	"strings"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// matchPattern checks if the given path matches the specified pattern.
// It supports dynamic segments prefixed with ":" and wildcard segments prefixed with "*".
// It returns a map of parameters extracted from the path if it matches the pattern,
// or nil and false if it does not match.
// The function splits both the pattern and path into parts, iterating through them to check for matches.
// If a part of the pattern starts with ":", it is treated as a dynamic segment and its value is stored in the params map.
// If a part of the pattern starts with "*", it captures the rest of the path as a wildcard segment.
// If the pattern and path match exactly, it returns the params map.
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

// ServeHTTP is the main entry point for handling HTTP requests.
// It checks if a handler exists for the request method and path.
// If a static handler is found, it creates a new context and applies middlewares in reverse order.
// If a dynamic route matches, it extracts parameters and applies middlewares.
// If no handler is found, it checks for method not allowed and not found handlers.
// If no handlers are defined, it falls back to the default http.NotFound handler.
// It uses the context package to create a new context for each request,
// allowing access to request and response data, as well as any parameters extracted from dynamic routes.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	if handler, ok := r.Handlers[method][path]; ok {
		ctx := context.NewContext(w, req)
		for i := len(r.Middlewares) - 1; i >= 0; i-- {
			handler = r.Middlewares[i](handler)
		}

		handler(ctx)
		return
	}

	for _, route := range r.DynamicRoutes {
		if route.method != method {
			continue
		}
		params, ok := matchPattern(route.pattern, path)
		if ok {
			ctx := context.NewContext(w, req)
			ctx.Params = params

			// middlewares
			handler := route.handler
			for i := len(r.Middlewares) - 1; i >= 0; i-- {
				handler = r.Middlewares[i](handler)
			}

			handler(ctx)
			return
		}
	}

	for m, routes := range r.Handlers {
		if m != method {
			if _, ok := routes[path]; ok && r.MethodNotAllowedHandler != nil {
				ctx := context.NewContext(w, req)
				r.MethodNotAllowedHandler(ctx)
				return
			}
		}
	}

	if r.NotFoundHandler != nil {
		ctx := context.NewContext(w, req)
		r.NotFoundHandler(ctx)
		return
	}

	http.NotFound(w, req)
}

// Listen starts the HTTP server on the given address using the router as handler.
// It uses http.ListenAndServe to bind the router to the specified address.
// This function is typically called in the main function of the application to start serving requests.
// It returns an error if the server fails to start, allowing the caller to handle it appropriately.
func (r *Router) Listen(addr string) error {
	return http.ListenAndServe(addr, r)
}
