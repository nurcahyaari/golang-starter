package rsa

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// ReadPublicKey will return public key
func ReadPublicKey(fileLoc string) (*rsa.PublicKey, error) {
	publicKeyFile, err := os.Open(fileLoc) //openssl rsa -in app.rsa -pubout > app.rsa.pub
	if err != nil {
		return nil, err
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	publicKeyFile.Close()

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

func ReadPrivateKey(fileLoc string) (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(fileLoc)
	if err != nil {
		return nil, err
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		return nil, err
	}
	return privateKeyImported, nil
}
