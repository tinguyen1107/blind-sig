package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"examples/bank/cmd/core"
)

var checkSpent = &cobra.Command{
	Use: "checkSpent",
	Run: func(cmd *cobra.Command, args []string) {
		var ticket = ""
		if len(args) >= 1 && args[0] != "" {
			ticket = args[0]
		} else {
			fmt.Printf("Invalid args")
			return
		}
		isSpent, err := core.IsTicketExist(ticket, SpentFile)
		if err != nil {
			fmt.Printf("Check spent error:", err)
			return
		}
		if isSpent {
			fmt.Printf("Ticket already used")
		} else {
			fmt.Printf("Ticket is new")
		}
	},
}
