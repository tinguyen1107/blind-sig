package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"examples/bank/cmd/core"
)

// Receive userId
var getBalance = &cobra.Command{
	Use: "getBalance",
	Run: func(cmd *cobra.Command, args []string) {
		var userId = ""
		if len(args) >= 1 && args[0] != "" {
			userId = args[0]
		} else {
			fmt.Printf("Invalid args")
			return
		}
		balance, err := core.GetBalance(userId, BalanceFile)
		if err != nil {
			fmt.Printf("Failed to get balance of %s: %s", userId, err.Error())
			return
		}

		fmt.Printf("%s's balance: %d", userId, balance)
	},
}
