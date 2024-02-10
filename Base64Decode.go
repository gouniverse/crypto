package crypto

import "encoding/base64"

func Base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
