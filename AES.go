package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AESEncrypt(input []byte, key string) (bytes []byte, err error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return bytes, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return bytes, err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return bytes, err
	}

	return gcm.Seal(nonce, nonce, input, nil), nil
}

func AESDecrypt(input []byte, key string) (bytes []byte, err error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return bytes, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return bytes, err
	}

	if len(input) < gcm.NonceSize() {
		return bytes, errors.New("malformed input")
	}

	output, err := gcm.Open(nil,
		input[:gcm.NonceSize()],
		input[gcm.NonceSize():],
		nil,
	)

	return output, err
}
