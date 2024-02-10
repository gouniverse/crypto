package crypto

import "testing"

func TestPasswordFortify(t *testing.T) {
	result := PasswordFortify("Hello world")

	if result == "" {
		t.Fatal("result MUST NOT be empty")
	}

	if len(result) != 256 {
		t.Fatal("result MUST BE be 256, was:", len(result))
	}

	result2 := PasswordFortify("Hello world")

	if result != result2 {
		t.Fatal("results MUST match", result, result2)
	}
}
