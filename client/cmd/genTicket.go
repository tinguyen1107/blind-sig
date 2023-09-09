package cmd

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"

	"examples/client/cmd/core"
)

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

var genTicket = &cobra.Command{
	Use: "genTicket",
	Run: func(cmd *cobra.Command, args []string) {
		var amount = 0

		if len(args) >= 1 && args[0] != "" {
			value, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid arg")
				return
			}
			if value%SmallestDevision != 0 {
				fmt.Println("Amount must be mutiples of 50.000")
				return
			}
			amount = value
		}
		// Generate a pubKey
		pubKey, err := loadPublicKeyFromFile(PubKeyFile)
		if err != nil {
			fmt.Println("Failed to load PublicKey")
			return
		}

		fmt.Printf("Generate %d tickets to transfer %dvnd\n", amount/SmallestDevision, amount)
		for i := 0; i < amount/SmallestDevision; i++ {
			r := generateRandomString(16)
			hashed := fdh.Sum(crypto.SHA256, HashSize, []byte(r))

			// Blind the hashed message
			blinded, unblinder, err := rsablind.Blind(pubKey, hashed)
			if err != nil {
				fmt.Println("Blind message failed")
				return
			}

			err = core.InsertNewElement(
				DataFile,
				hex.EncodeToString(hashed),
				hex.EncodeToString(blinded),
				hex.EncodeToString(unblinder),
			)

			if err != nil {
				fmt.Println("Failed to write to data file:", err)
				return
			}
		}
	},
}
