package crypto

import (
	"errors"
	"strconv"
	"strings"
)

func XorFortifiedDecrypt(value string, password string) (string, error) {
	strongPassword := PasswordFortify(password)
	first, err := XorDecrypt(value, strongPassword)

	if err != nil {
		return "", errors.New("xor. " + err.Error())
	}

	if !IsBase64(first) {
		return "", errors.New("password incorrect")
	}

	v4, err := Base64Decode(first)

	if err != nil {
		return "", errors.New("base64.1. " + err.Error())
	}

	a := strings.Split(string(v4), "_")

	if len(a) < 2 {
		return "", errors.New("password incorrect")
	}

	upTo, err := strconv.Atoi(a[0])

	if err != nil {
		return "", errors.New("atoi. " + err.Error())
	}

	v1 := a[1][0:upTo]

	v2, err := Base64Decode(v1)
	if err != nil {
		return "", errors.New("base64.2. " + err.Error())
	}

	return string(v2), nil
}

func XorFortifiedEncrypt(value string, password string) string {
	strongPassword := PasswordFortify(password)

	v1 := Base64Encode([]byte(value))

	v2 := strconv.Itoa(len(v1)) + "_" + v1

	randomBlock := StrRandom(RequiredBlockLength(len(v2)))

	v3 := v2 + "" + randomBlock[len(v2):]

	v4 := Base64Encode([]byte(v3))

	last := XorEncrypt(v4, strongPassword)

	return last
}
