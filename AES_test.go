package crypto

import "testing"

func TestAESEncrypt(t *testing.T) {
	p := PasswordFortify("Hello world")[:32]
	result, err := AESEncrypt([]byte("Hello world"), p)

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(result) < 1 {
		t.Fatal("result MUST NOT be empty")
	}

	original, err := AESDecrypt(result, p)

	if err != nil {
		t.Fatal(err.Error())
	}

	if string(original) != "Hello world" {
		t.Fatal("original MUST match 'Hello world'", original)
	}
}
