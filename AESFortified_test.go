package crypto

import "testing"

func TestAESFortified(t *testing.T) {
	original := "Hello world"
	password := "Star Wars"
	enc, err := AESFortifiedEncrypt(original, password)

	if err != nil {
		t.Fatal(err.Error(), enc)
	}

	if enc == "" {
		t.Fatal("enc MUST NOT be empty")
	}

	dec, err := AESFortifiedDecrypt(enc, password)

	if err != nil {
		t.Fatal(err.Error(), enc)
	}

	if dec != original {
		t.Fatal("dec MUST match:", original, " found: ", dec)
	}
}
