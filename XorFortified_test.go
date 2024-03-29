package crypto

import "testing"

func TestXorFortified(t *testing.T) {
	original := "Hello world"
	password := "Star Wars"
	enc := XorFortifiedEncrypt(original, password)

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
