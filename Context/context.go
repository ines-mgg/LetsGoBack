package Context

import "net/http"

// NewContext creates and returns a new Context instance.
// It initializes the Context with the provided http.ResponseWriter and http.Request,
// sets up empty maps for Params and Data, and assigns the request's URL path and method.
// This function is typically used to encapsulate HTTP request and response data
// for further processing within the application.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Params:  make(map[string]string),
		Path:    r.URL.Path,
		Method:  r.Method,
		Data:    make(map[string]any),
	}
}

// Get retrieves the value associated with the given key from the Context's data map.
// It returns the value (of type any) and a boolean indicating whether the key was found.
func (c *Context) Get(key string) (any, bool) {
	val, ok := c.Data[key]
	return val, ok
}

// Set stores a key-value pair in the Context's data map.
// The key is a string, and the value can be of any type.
func (c *Context) Set(key string, value any) {
	c.Data[key] = value
}

// GetStatus retrieves the HTTP status code from the Context.
// If the status is not set, it returns 200 (OK) as the default status code.
// It also sets the "status" key in the Data map to the retrieved status code.
func (c *Context) GetStatus() int {
	return c.Status
}

// SetStatus sets the HTTP status code in the Context.
// It updates the "status" key in the Data map with the provided status code.
func (c *Context) SetStatus(status int) {
	c.Data["status"] = status
}

// GetMethod retrieves the HTTP method (e.g., GET, POST) from the Context.
// It returns the method as a string.
// This method is useful for determining the type of HTTP request being processed.
func (c *Context) GetMethod() string {
	return c.Method
}

// GetPath retrieves the URL path from the Context.
// It returns the path as a string.
func (c *Context) GetPath() string {
	return c.Path
}

// Param retrieves the value of a specific parameter from the Context's Params map.
// It takes a key as an argument and returns the corresponding value as a string.
// If the key does not exist in the Params map, it returns an empty string.
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// RequestID retrieves the "request_id" value from the context as a string.
// If the "request_id" is not set or is not a string, it returns an empty string.
func (c *Context) RequestID() string {
	val, ok := c.Get("request_id")
	if !ok {
		return ""
	}
	if id, ok := val.(string); ok {
		return id
	}
	return ""
}
