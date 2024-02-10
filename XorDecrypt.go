package crypto

// xorDecrypt  runs a XOR encryption on the input string
func XorDecrypt(encstring string, key string) (output string, err error) {
	inputBytes, err := Base64Decode(encstring)

	if err != nil {
		return "", err
	}

	input := string(inputBytes)

	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output, nil
}
