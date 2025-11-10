package crypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func SHA256Hex(data []byte) string {
	return hex.EncodeToString(SHA256(data))
}
