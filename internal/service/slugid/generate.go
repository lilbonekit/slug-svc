package slugid

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func Generate(n int) string {
	if n <= 0 {
		n = 6
	}
	raw := make([]byte, n)
	_, _ = rand.Read(raw)

	s := base64.RawURLEncoding.EncodeToString(raw)
	s = strings.TrimRight(s, "-_")
	if len(s) < n {
		return s
	}
	return s[:n]
}
