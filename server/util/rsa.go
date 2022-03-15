package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"eea/config"
	"encoding/pem"

	"github.com/sirupsen/logrus"
)

var (
	privateKey *rsa.PrivateKey
)

func Init() {
	block, _ := pem.Decode([]byte(config.Configs.RSA))
	if block == nil {
		logrus.Fatal("private key error")
	}
	var err error
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logrus.Fatal("init private key failed:" + err.Error())
	}

}

func RSADecrypt(cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}
