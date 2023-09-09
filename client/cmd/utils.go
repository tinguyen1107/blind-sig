package cmd

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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

func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	// Read the private key file
	privKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode PEM data
	block, _ := pem.Decode(privKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Parse the private key
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
