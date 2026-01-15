package slug

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateSlug() string {
	bytes := make([]byte, 7)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Error reading random bytes: %v", err)
	}
	return hex.EncodeToString(bytes)
}
