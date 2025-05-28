package Context

import (
	"crypto/rand"
	"fmt"
)

func GenerateErrorID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
