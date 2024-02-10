package crypto

import (
	"errors"
	"strconv"
	"strings"
)

func AESFortifiedDecrypt(value string, password string) (string, error) {
	inputBytes, err := Base64Decode(value)

	if err != nil {
		return "", err
	}

	strongPassword := PasswordFortify(password)
	first, err := AESDecrypt(inputBytes, strongPassword[:32])

	if err != nil {
		return "", errors.New("xor. " + err.Error())
	}

	v4, err := Base64Decode(string(first))

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

func AESFortifiedEncrypt(value string, password string) (string, error) {
	strongPassword := PasswordFortify(password)

	v1 := Base64Encode([]byte(value))

	v2 := strconv.Itoa(len(v1)) + "_" + v1

	randomBlock := StrRandom(RequiredBlockLength(len(v2)))

	v3 := v2 + "" + randomBlock[len(v2):]

	v4 := Base64Encode([]byte(v3))

	outputBytes, err := AESEncrypt([]byte(v4), strongPassword[:32])

	if err != nil {
		return "", err
	}

	output := Base64Encode(outputBytes)

	return output, nil
}
