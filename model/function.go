package model

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

//GenerateToken function generates a new token
func GenerateToken() string {
	b := make([]byte, 16)
	rand.Read(b)

	hasher := md5.New()
	hasher.Write(b)

	return hex.EncodeToString(hasher.Sum(nil))
}