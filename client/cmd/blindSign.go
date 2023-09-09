package cmd

import (
	"github.com/spf13/cobra"
)

// Receive userId and ticket
var blindSign = &cobra.Command{
	Use: "blindSign",
	Run: func(cmd *cobra.Command, args []string) {
		// var userId = ""
		// var blindedTicket = ""
		// if len(args) >= 2 && args[0] != "" && args[2] != "" {
		// 	userId = args[0]
		// 	blindedTicket = args[1]
		// } else {
		// 	fmt.Printf("Invalid args\n")
		// 	return
		// }
		//
		// // Check valid balance
		// // balance, err := core.GetBalance(userId, BalanceFile)
		// if err != nil {
		// 	fmt.Printf("Failed to get balance of %s: %s\n", userId, err.Error())
		// 	return
		// }
		// if balance < SmallestDevision {
		// 	fmt.Printf("Your balance is not enough, must larger than %d\n", SmallestDevision)
		// 	return
		// }
		//
		// privKey, err := loadPrivateKeyFromFile(PrivKeyFile)
		// if err != nil {
		// 	fmt.Println("Failed to load private key:", err)
		// 	return
		// }
		// sig, err := rsablind.BlindSign(privKey, []byte(blindedTicket))
		//
		// fmt.Println(sig)
	},
}
