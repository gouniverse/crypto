package crock32

import (
	"encoding/base32"
	"math/big"
)

func Decode(str string) int64 {
	bytes, err := base32.NewEncoding(crockfordAlphabet).
		WithPadding(base32.NoPadding).
		DecodeString(str)

	if err != nil {
		return 0
	}

	return big.NewInt(0).
		SetBytes(bytes).
		Int64()
}
