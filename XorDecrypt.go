package crypto

// XorDecrypt  runs a XOR decryption on the input string
func XorDecrypt(encString string, key string) (output string, err error) {
	inputBytes, err := Base64Decode(encString)

	if err != nil {
		return "", err
	}

	input := string(inputBytes)

	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output, nil
}
