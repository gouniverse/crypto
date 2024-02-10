package crypto

import "testing"

func TestStrongifyPassword(t *testing.T) {
	result := StrongifyPassword("Hello world")

	if result == "" {
		t.Fatal("result MUST NOT be empty")
	}

	if len(result) != 200 {
		t.Fatal("result MUST BE be 144, was:", len(result))
	}

	result2 := StrongifyPassword("Hello world")

	if result != result2 {
		t.Fatal("results MUST match", result, result2)
	}
}
