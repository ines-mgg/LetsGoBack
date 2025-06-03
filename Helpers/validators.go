package Helpers

import (
	"strings"
)

// IsEmpty checks if a string is empty.
func IsNotEmpty(s string) bool {
	return len(s) > 0
}

// IsValidEmail checks if the provided email is valid.
// It checks the length, presence of '@', and '.' characters.
// A valid email must have exactly one '@' character, not start or end with it,
// and must not have '.' at the start or end or immediately after '@'.
// The email length must be between 3 and 254 characters.
// It returns true if the email is valid, false otherwise.
func IsValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	at := 0
	for i, c := range email {
		if c == '@' {
			at++
			if at > 1 || i == 0 || i == len(email)-1 {
				return false
			}
		} else if c == '.' && (i == 0 || i == len(email)-1 || email[i-1] == '@') {
			return false
		}
	}
	return at == 1
}

// IsValidPassword checks if the provided password is valid.
// A valid password must be between 8 and max characters long (default is 128),
// contain at least one uppercase letter, one lowercase letter, one digit,
// and one special character from the set "!@#$%^&*()-_=+[]{}|;:',.<>?/".
// It returns true if the password is valid, false otherwise.
// If max is less than or equal to 0 or 8, it defaults to 128.
func IsValidPassword(password string, max int) bool {
	if max <= 0 || max <= 8 {
		max = 128 // Default max length
	}
	if len(password) < 8 || len(password) > max {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	specialChars := "!@#$%^&*()-_=+[]{}|;:',.<>?/"
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		} else if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= '0' && c <= '9' {
			hasDigit = true
		} else if strings.ContainsRune(specialChars, c) {
			hasSpecial = true
		}

	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}

// IsValidURL checks if the provided URL is valid.
// A valid URL must start with "http://" or "https://", and must be at least 3 characters long.
// It returns true if the URL is valid, false otherwise.
func IsValidURL(url string) bool {
	if len(url) < 3 || !strings.Contains(url, "://") {
		return false
	}
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}

// IsGreaterThanInt checks if an int value is greater than min.
func IsGreaterThanInt(value int, min int) bool {
	return value > min
}

// IsGreaterThanFloat checks if a float64 value is greater than min.
func IsGreaterThanFloat(value float64, min float64) bool {
	return value > min
}

// IsGreaterThanOrEqualInt checks if an int value is greater than or equal to min.
func IsGreaterThanOrEqualInt(value int, min int) bool {
	return value >= min
}

// IsGreaterThanOrEqualFloat checks if a float64 value is greater than or equal to min.
func IsGreaterThanOrEqualFloat(value float64, min float64) bool {
	return value >= min
}

// IsLessThanInt checks if an int value is less than max.
func IsLessThanInt(value int, max int) bool {
	return value < max
}

// IsLessThanFloat checks if a float64 value is less than max.
func IsLessThanFloat(value float64, max float64) bool {
	return value < max
}

// IsLessThanOrEqualInt checks if an int value is less than or equal to max.
func IsLessThanOrEqualInt(value int, max int) bool {
	return value <= max
}

// IsLessThanOrEqualFloat checks if a float64 value is less than or equal to max.
func IsLessThanOrEqualFloat(value float64, max float64) bool {
	return value <= max
}

// IsInRangeInt checks if an int value is within the range [min, max] (inclusive).
func IsInRangeInt(value int, min int, max int) bool {
	return value >= min && value <= max
}

// IsInRangeFloat checks if a float64 value is within the range [min, max] (inclusive).
func IsInRangeFloat(value float64, min float64, max float64) bool {
	return value >= min && value <= max
}
