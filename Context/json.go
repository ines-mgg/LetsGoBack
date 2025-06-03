package Context

import "encoding/json"

// BindJSON binds the request body to the provided object.
// It uses the json package to decode the JSON data from the request body into the specified object.
// If there is an error during decoding, it returns the error.
// The object must be a pointer to a struct or a map that matches the JSON structure.
func (c *Context) BindJSON(obj any) error {
	decoder := json.NewDecoder(c.Request.Body)
	return decoder.Decode(obj)
}

// json sends a JSON response with the specified status code and message.
// It sets the Content-Type header to "application/json" and writes the status code to the response.
// The message can be any type that can be marshaled to JSON.
// It uses the json package to encode the message into JSON format and writes it to the response writer.
// This function is typically used to send structured data back to the client in a JSON format.
// The status code indicates the HTTP status of the response, such as 200 for success or 404 for not found.
func (c *Context) json(status int, message any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.SetStatus(status)
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(message)
}

// abortWithStatusJSON sends a JSON response with an error message and the specified status code.
// It sets the Content-Type header to "application/json" and writes the status code to the response.
// The message is a string that describes the error or issue encountered.
// This function is typically used to handle errors in a consistent manner, providing a structured JSON response
// to the client when an error occurs.
func (c *Context) abortWithStatusJSON(status int, message string) {
	c.json(status, map[string]string{
		"error": message,
	})
}
