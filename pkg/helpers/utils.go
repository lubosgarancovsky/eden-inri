package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func OpaqueToken(bytes int) (string, error) {
	b := make([]byte, bytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// URL-safe, no padding
	return base64.RawURLEncoding.EncodeToString(b), nil
}
