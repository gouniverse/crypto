package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func StrToSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
