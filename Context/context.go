package Context

import "net/http"

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

func (c *Context) Get(key string) (any, bool) {
	val, ok := c.Data[key]
	return val, ok
}

func (c *Context) Set(key string, value any) {
	c.Data[key] = value
}

func (c *Context) GetStatus() int {
	return c.Status
}

func (c *Context) SetStatus(status int) {
	c.Data["status"] = status
}

func (c *Context) GetMethod() string {
	return c.Method
}

func (c *Context) GetPath() string {
	return c.Path
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

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
