package cmd

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func savePrivateKeyToFile(key *rsa.PrivateKey, filename string) error {
	privateKeyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		},
	)
	_, err = privateKeyFile.Write(privKeyPEM)
	return err
}

func savePublicKeyToFile(pubKey *rsa.PublicKey, filename string) error {
	publicKeyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}

	pubKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)

	_, err = publicKeyFile.Write(pubKeyPEM)
	return err
}
