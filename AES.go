package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AESEncrypt(input string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	output := Base64Encode(gcm.Seal(nonce, nonce, []byte(input), nil))

	return output, nil
}

func AESDecrypt(input string, key string) (string, error) {
	ciphertext, err := Base64Decode(input)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}

	output, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)

	return string(output), err
}
