package Middleware

// LogLevel determines the log level based on the HTTP status code.
// It returns a string representing the log level:
// - "[ERROR]" for status codes 500 and above
// - "[WARN]" for status codes 400 to 499
// - "[INFO]" for status codes below 400
// This function is typically used in logging middleware to categorize log messages based on the response status.
// It helps in filtering and prioritizing log messages based on their severity.
// Example usage:
//   status := 404
//   logLevel := LogLevel(status)
//   fmt.Println(logLevel) // Output: [WARN]
func LogLevel(status int) string {
	switch {
	case status >= 500:
		return "[ERROR]"
	case status >= 400:
		return "[WARN]"
	default:
		return "[INFO]"
	}
}
