package slugid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

func Generate(n int) (string, error) {
	if n <= 0 {
		n = 6
	}

	raw := make([]byte, n)
	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}

	s := base64.RawURLEncoding.EncodeToString(raw)
	s = strings.TrimRight(s, "-_")

	if len(s) < n {
		return s, nil
	}
	return s[:n], nil
}
