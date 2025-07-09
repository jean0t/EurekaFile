package database

import (
	"fmt"
	"hash"
	"crypto/sha256"
)

func HashPassword(password string) string {
	var hash hash.Hash = sha256.New()
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
