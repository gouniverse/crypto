package crypto

// xorEncrypt  runs a XOR encryption on the input string
func XorEncrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return Base64Encode([]byte(output))
}
