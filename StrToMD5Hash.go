package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
