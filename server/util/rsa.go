package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"eea/config"
	"encoding/pem"
	"errors"
)

func RSADecrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(config.RSAPrivateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}
