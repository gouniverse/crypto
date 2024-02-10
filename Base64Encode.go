package crypto

import "encoding/base64"

func Base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}
