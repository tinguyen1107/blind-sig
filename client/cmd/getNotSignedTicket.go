package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"examples/client/cmd/core"
)

var getNotSignedTicket = &cobra.Command{
	Use: "getNotSignedTicket",
	Run: func(cmd *cobra.Command, args []string) {
		ticket, id, err := core.GetIncompleteElement(DataFile)
		if err != nil {
			fmt.Println("Failed get not signed ticket:", err)
			return
		}
		jsonStr, err := json.MarshalIndent(ticket, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Not Signed Ticket (%d): \n%s\n", id, string(jsonStr))
	},
}
