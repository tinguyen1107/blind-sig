package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/spf13/cobra"
)

var genKey = &cobra.Command{
	Use: "genKey",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := rsa.GenerateKey(rand.Reader, KeySize)
		if err != nil {
			fmt.Println("Failed to generate key:", err)
			return
		}
		err = savePrivateKeyToFile(key, PrivKeyFile)
		if err != nil {
			fmt.Println("Failed to save private key:", err)
			return
		}

		err = savePublicKeyToFile(&key.PublicKey, PubKeyFile)
		if err != nil {
			fmt.Println("Failed to save public key:", err)
			return
		}
	},
}
