package crypto

// StrongifyPassword performs multiple calculations
// on top of the password and changes it to a derivative
// long hash. This is done so that even simple and not-long
// passwords  can become longer and stronger (200 characters).
func StrongifyPassword(password string) string {
	p1 := StrToMD5Hash(password) + StrToMD5Hash(password) + StrToMD5Hash(password) + StrToMD5Hash(password)

	p1 = StrToSHA256Hash(p1)
	p2 := StrToSHA1Hash(p1) + StrToSHA1Hash(p1) + StrToSHA1Hash(p1)
	p3 := StrToSHA1Hash(p2) + StrToMD5Hash(p2) + StrToSHA1Hash(p2)
	p4 := StrToSHA256Hash(p3)
	p5 := StrToSHA256Hash(p4) + StrToSHA256Hash(StrToMD5Hash(p4)) + StrToSHA256Hash(StrToSHA1Hash(p4)) + StrToSHA256Hash(p4)
	return p5
}
