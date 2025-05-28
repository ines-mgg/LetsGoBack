package Middleware

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
