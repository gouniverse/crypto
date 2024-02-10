package crypto

import "math/rand"

// StrRandom returns a random string of specified length
func StrRandom(length int) string {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersLength := len(characters)
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(characters[rand.Intn(charactersLength-1)])
	}
	return randomString
}
