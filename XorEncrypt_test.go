package crypto

import "testing"

func TestXorEncrypt(t *testing.T) {
	p := PasswordFortify("Hello world")
	result := XorEncrypt("Hello world", p)

	if result == "" {
		t.Fatal("result MUST NOT be empty")
	}

	// if len(result) != 200 {
	// 	t.Fatal("result MUST BE be 144, was:", len(result))
	// }

	original, err := XorDecrypt(result, p)

	if err != nil {
		t.Fatal(err.Error())
	}

	if original != "Hello world" {
		t.Fatal("original MUST match 'Hello world'", original)
	}
}
