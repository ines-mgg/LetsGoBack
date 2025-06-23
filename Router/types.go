package Router

import (
	context "github.com/ines-mgg/LetsGoBack/Context"
	middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

// dynamicRoute represents a route that can have dynamic parameters.
// It contains the HTTP method, the route pattern, and the handler function.
// The pattern can include parameters prefixed with ":" for single parameters or "*" for catch-all parameters.
// The handler function is a context.HandlerFunc that will be executed when the route is matched.
// The dynamic routes are stored in a slice of dynamicRoute structs within the Router struct.
// This allows the router to handle routes with dynamic segments, such as "/users/:id" or "/files/*filepath".
type dynamicRoute struct {
	method  string
	pattern string
	handler context.HandlerFunc
}

// Router is the main structure that holds all the routes and their handlers.
// It contains a map of HTTP methods to their corresponding routes and handlers.
// The Handlers map is a nested map where the first key is the HTTP method (e.g., "GET", "POST"),
// and the second key is the route path (e.g., "/users/:id").
// The Handlers map allows for quick lookup of handlers based on the HTTP method and route path.
// The Router also maintains a slice of dynamicRoute structs to handle routes with dynamic parameters.
// The Middlewares slice contains middleware functions that can be applied to all routes.
// The NotFoundHandler is a context.HandlerFunc that will be called when no route matches the request.
// The MethodNotAllowedHandler is a context.HandlerFunc that will be called when the method is not allowed for a specific route.
type Router struct {
	Handlers      map[string]map[string]context.HandlerFunc
	DynamicRoutes []dynamicRoute
	Middlewares   []middleware.Middleware

	NotFoundHandler         context.HandlerFunc
	MethodNotAllowedHandler context.HandlerFunc
}

// routeGroup represents a group of routes with a common prefix and shared middlewares.
// It contains the prefix for the group, a reference to the parent router,
// and a slice of middlewares that will be applied to all routes in the group.
// The prefix is a string that will be prepended to all routes defined in this group.
// The router field is a pointer to the parent Router, allowing the group to add routes to it.
// The middlewares field is a slice of middleware.Middleware that can be used to apply common functionality
// to all routes in the group, such as logging, authentication, or error handling.
type routeGroup struct {
	prefix      string
	router      *Router
	middlewares []middleware.Middleware
}
