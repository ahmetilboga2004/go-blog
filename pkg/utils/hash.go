package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	saltedPassword := fmt.Sprintf("%s%s", salt, password)
	hash := sha256.Sum256([]byte(saltedPassword))

	return hex.EncodeToString(hash[:])
}
