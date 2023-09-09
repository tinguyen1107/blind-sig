package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"
)

var verifySig = &cobra.Command{
	Use: "verifySig",
	Run: func(cmd *cobra.Command, args []string) {
		hash := ""
		sig := ""
		if len(args) >= 2 && args[0] != "" && args[1] != "" {
			hash = args[0]
			sig = args[1]
		}

		pubKey, err := loadPublicKeyFromFile(PubKeyFile)
		if err != nil {
			fmt.Println("Failed to load PublicKey")
			return
		}

		hashHex, _ := hex.DecodeString(hash)
		sigHex, _ := hex.DecodeString(sig)

		if err := rsablind.VerifyBlindSignature(pubKey, hashHex, sigHex); err != nil {
			panic("failed to verify signature")
		} else {
			fmt.Println("ALL IS WELL")
		}
	},
}
