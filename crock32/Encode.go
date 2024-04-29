package crock32

import (
	"encoding/base32"
	"math/big"
)

func Encode(num int64) string {
	var bytes = big.NewInt(num).Bytes()

	str := base32.NewEncoding(crockfordAlphabet).
		WithPadding(base32.NoPadding).
		EncodeToString(bytes)

	return str
}
