package crypto

import "testing"

func TestStrongifyPassword(t *testing.T) {
	result := StrongifyPassword("Hello world")

	if result == "" {
		t.Fatal("result MUST NOT be empty")
	}

	if len(result) != 256 {
		t.Fatal("result MUST BE be 256, was:", len(result))
	}

	result2 := StrongifyPassword("Hello world")

	if result != result2 {
		t.Fatal("results MUST match", result, result2)
	}
}
