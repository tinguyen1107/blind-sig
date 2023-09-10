package cmd

import (
	"encoding/hex"
	"examples/client/cmd/core"
	"fmt"
	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"
)

// <index> <sig>

func CreateSetSigCmd() *cobra.Command {
	setSig := &cobra.Command{
		Use: "setSig",
		Run: func(cmd *cobra.Command, args []string) {

			index, _ := cmd.Flags().GetInt("index")
			sig, _ := cmd.Flags().GetString("sig")
			if index < 0 {
				fmt.Println("Invalid index")
				return
			}
			if len(sig) <= 0 {
				fmt.Println("Invalid signature")
				return
			}

			//if len(args) >= 2 && args[0] != "" && args[1] != "" {
			//	value, err := strconv.Atoi(args[0])
			//	if err != nil {
			//		fmt.Println("Invalid arg")
			//		return
			//	}
			//	if value%SmallestDevision != 0 {
			//		fmt.Println("Amount must be mutiples of 50.000")
			//		return
			//	}
			//	index = value
			//	sig = args[1]
			//}
			pubKey, err := loadPublicKeyFromFile(PubKeyFile)
			if err != nil {
				fmt.Println("Failed to load PublicKey")
				return
			}
			ticket, err := core.GetTicket(DataFile, index)
			sighex, _ := hex.DecodeString(sig)
			unblinder, _ := hex.DecodeString(ticket.Unblinder)
			unblindedSig := rsablind.Unblind(pubKey, sighex, unblinder)

			err = core.UpdateElementByIndex(DataFile, index, sig, hex.EncodeToString(unblindedSig))
			if err != nil {
				fmt.Println("Failed get not signed ticket:", err)
				return
			}

			fmt.Println("Done")
		},
	}

	setSig.PersistentFlags().Int("index", -1, "Index")
	setSig.PersistentFlags().String("sig", "", "Signature")

	return setSig
}
