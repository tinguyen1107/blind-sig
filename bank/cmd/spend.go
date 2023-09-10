package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"

	"examples/bank/cmd/core"
)

// Receive userId and ticket path
var spend = &cobra.Command{
	Use: "spend",
	Run: func(cmd *cobra.Command, args []string) {
		var userId = ""
		var ticket = ""
		var sig = ""
		if len(args) >= 3 && args[0] != "" && args[1] != "" && args[2] != "" {
			userId = args[0]
			ticket = args[1]
			sig = args[2]
		} else {
			fmt.Printf("Invalid args")
			return
		}

		pubKey, err := loadPublicKeyFromFile(PubKeyFile)
		if err != nil {
			fmt.Println("Failed to load PublicKey")
			return
		}

		hashHex, _ := hex.DecodeString(ticket)
		sigHex, _ := hex.DecodeString(sig)

		if err := rsablind.VerifyBlindSignature(pubKey, hashHex, sigHex); err != nil {
			panic("failed to verify signature")
		}

		isSpent, err := core.IsTicketExist(ticket, SpentFile)
		if err != nil {
			fmt.Printf("Check spent error:", err)
			return
		}
		if isSpent {
			fmt.Printf("Ticket already used")
			return
		} else {
			core.AddTicket(ticket, SpentFile)
		}

		balance, err := core.GetBalance(userId, BalanceFile)
		if err != nil {
			fmt.Printf("Failed to get balance of %s: %s", userId, err.Error())
			return
		}

		core.UpdateBalance(userId, balance-SmallestDevision, BalanceFile)
	},
}
