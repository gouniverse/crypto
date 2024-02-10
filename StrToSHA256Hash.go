package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func StrToSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
