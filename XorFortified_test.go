package crypto

import "testing"

func TestML1(t *testing.T) {
	original := "Hello world"
	password := "Hello world"
	enc := XorFortifiedEncrypt("Hello world", password)

	if enc == "" {
		t.Fatal("enc MUST NOT be empty")
	}

	dec, err := XorFortifiedDecrypt(enc, password)

	if err != nil {
		t.Fatal(err.Error(), enc)
	}

	if dec != original {
		t.Fatal("dec MUST match:", original, " found: ", dec)
	}
}
