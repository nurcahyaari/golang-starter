package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func AddKeyLen(encryptionKey string) string {
	var min int
	encryptionKeyLen := len(encryptionKey)
	if encryptionKeyLen < 32 {
		min = 32 - encryptionKeyLen
	} else if encryptionKeyLen > 32 {
		min = encryptionKeyLen - 32
	}
	return encryptionKey + encryptionKey[0:min]
}

// encryptionKey should have 32 of length
func AesCFBDecryption(text string, encryptionKey string) string {
	encryptionKey = AddKeyLen(encryptionKey)

	key, _ := hex.DecodeString(
		encryptionKey,
	)
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	// fmt.Printf("%s", ciphertext)
	// Output: some plaintext

	return string(ciphertext)
}

// ecryptionKey length must 32
func AesCFBEncryption(text string, encryptionKey string) string {
	encryptionKey = AddKeyLen(encryptionKey)
	// Load secret key from app config
	key, _ := hex.DecodeString(
		encryptionKey,
	)
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return fmt.Sprintf("%x", ciphertext)

}
