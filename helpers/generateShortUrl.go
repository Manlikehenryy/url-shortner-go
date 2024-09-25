package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Function to generate a short URL (hash of the original URL)
func GenerateShortURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL + fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(hash[:])[:8] // Return only the first 8 characters
}