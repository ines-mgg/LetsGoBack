package Router

import (
	context "lets-go-back/Context"
	"net/http"
)

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
		if route.Method != method {
			continue
		}
		params, ok := MatchPattern(route.Pattern, path)
		if ok {
			ctx := context.NewContext(w, req)
			ctx.Params = params

			// middlewares
			handler := route.Handler
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
