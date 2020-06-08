package utils

import (
	"math/rand"
	"time"
)

// GenerateKey returns a 16 characters long random string.
func GenerateKey() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	tokenBytes := make([]rune, 16)
	for i := range tokenBytes {
		tokenBytes[i] = letters[rand.Intn(len(letters))]
	}

	return string(tokenBytes)
}
