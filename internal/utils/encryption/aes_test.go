package encryption_test

import (
	"golang-starter/internal/utils/encryption"
	"testing"
)

func TestEncryptDecryptAesCBF(t *testing.T) {
	cipherText, _ := encryption.AesCFBEncryption(
		"saya makan beras",
		"12345678123456781234567812345678",
	)

	if len(cipherText) == 0 {
		t.Errorf("Maybe there was an incorrect when encrypt this data")
	}

	plainText, _ := encryption.AesCFBDecryption(
		cipherText,
		"12345678123456781234567812345678",
	)

	if plainText != "saya makan beras" {
		t.Errorf("Error when decrypting text")
	}
}

func TestAddKeyLen(t *testing.T) {
	newText := encryption.AddKeyLen("greg")

	if len(newText) != 32 {
		t.Errorf("Len is not 32")
	}
	if newText != "sayamakanmiesayamakanmiesayamakanmie" {
		t.Errorf("String is not same as expected")
	}
}
