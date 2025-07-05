package auth

import (
	"crypto/sha256"
)

func CodeHash(code string) []byte {
	sha := sha256.New()
	sha.Write([]byte(code))
	return sha.Sum(nil)
}
