package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// ReadPublicKey will return public key
func ReadPublicKeyFromEnv(rsaPublic string) (*rsa.PublicKey, error) {
	data, _ := pem.Decode([]byte(rsaPublic))
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("cannot reflect the interface")
	}

	return publicKey, nil
}

func ReadPrivateKeyFromEnv(rsaPrivate string) (*rsa.PrivateKey, error) {
	data, _ := pem.Decode([]byte(rsaPrivate))
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKeyImported, nil
}
