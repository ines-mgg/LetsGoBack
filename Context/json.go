package Context

import "encoding/json"

func (c *Context) JSON(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if c.GetStatus() == 0 {
		c.SetStatus(status)
		c.Writer.WriteHeader(status)
	}

	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) BindJSON(obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	return decoder.Decode(obj)
}

func (c *Context) AbortWithStatusJSON(status int, message string) {
	c.JSON(status, map[string]string{
		"error": message,
	})
}
