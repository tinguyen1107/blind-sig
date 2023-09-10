package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"

	"examples/bank/cmd/core"
)

// Receive userId and ticket
var blindSign = &cobra.Command{
	Use: "blindSign",
	Run: func(cmd *cobra.Command, args []string) {
		var userId = ""
		var blindedTicket = ""
		if len(args) >= 2 && args[0] != "" && args[1] != "" {
			userId = args[0]
			blindedTicket = args[1]
		} else {
			fmt.Printf("Invalid args\n")
			return
		}

		// Check valid balance
		balance, err := core.GetBalance(userId, BalanceFile)
		if err != nil {
			fmt.Printf("Failed to get balance of %s: %s\n", userId, err.Error())
			return
		}
		if balance < SmallestDevision {
			fmt.Printf("Your balance is not enough, must larger than %d\n", SmallestDevision)
			return
		}

		core.UpdateBalance(userId, balance-SmallestDevision, BalanceFile)

		privKey, err := loadPrivateKeyFromFile(PrivKeyFile)
		if err != nil {
			fmt.Println("Failed to load private key:", err)
			fmt.Println("Reverse balance")
			core.UpdateBalance(userId, balance, BalanceFile)
			return
		}
		blindedTicketHex, err := hex.DecodeString(blindedTicket)
		if err != nil {
			fmt.Println("Failed to hex decode ticket:", err)
			fmt.Println("Reverse balance")
			core.UpdateBalance(userId, balance, BalanceFile)
			return
		}
		sig, err := rsablind.BlindSign(privKey, blindedTicketHex)
		sigStr := hex.EncodeToString(sig)

		fmt.Println(sigStr)
	},
}
