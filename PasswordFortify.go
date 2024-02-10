package crypto

// PasswordFortify performs multiple calculations
// on top of the password and changes it to a derivative
// long hash. This is done so that even simple and not-long
// passwords  can become longer and stronger (256 characters).
// If password is longer than 255 it is left intact.
func PasswordFortify(password string) string {
	if len(password) > 255 {
		return password
	}

	pass1 := StrToSHA256Hash(password) + StrToSHA256Hash(password) + StrToSHA256Hash(password) + StrToSHA256Hash(password)
	pass2 := StrToSHA256Hash(pass1)
	pass3 := StrToSHA256Hash(pass2) + StrToSHA1Hash(pass2) + StrToSHA1Hash(pass2)
	pass4 := StrToSHA256Hash(pass3) + StrToMD5Hash(pass3) + StrToMD5Hash(pass3)
	pass5 := StrToSHA256Hash(pass4) + StrToSHA1Hash(pass4) + StrToSHA1Hash(pass4)
	pass6 := StrToSHA256Hash(pass5) + StrToSHA256Hash(StrToMD5Hash(pass5)) + StrToSHA256Hash(StrToSHA1Hash(pass5)) + StrToSHA256Hash(pass5)
	return pass6
}
