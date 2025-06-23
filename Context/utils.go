package Context

import (
	"crypto/rand"
	"fmt"
)

// GenerateErrorID generates a random error ID for logging or tracking purposes.
// It returns a hexadecimal string representation of the random bytes.
// This function can be used to create unique identifiers for errors or requests,
// which can help in debugging and tracing issues in the application.
// The generated ID is a 4-byte random value, which is sufficient for most use cases.
// The ID is generated using the crypto/rand package to ensure cryptographic security.
func GenerateErrorID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
