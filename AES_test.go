package crypto

import "testing"

func TestAESEncrypt(t *testing.T) {
	p := StrongifyPassword("Hello world")[:32]
	result, err := AESEncrypt("Hello world", p)

	if err != nil {
		t.Fatal(err.Error())
	}

	if result == "" {
		t.Fatal("result MUST NOT be empty")
	}

	// if len(result) != 200 {
	// 	t.Fatal("result MUST BE be 144, was:", len(result))
	// }

	original, err := AESDecrypt(result, p)

	if err != nil {
		t.Fatal(err.Error())
	}

	if original != "Hello world" {
		t.Fatal("original MUST match 'Hello world'", original)
	}
}
