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

// Receive userId
// var getBalance = &cobra.Command{
// 	Use: "getBalance",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		var userId = ""
// 		if len(args) >= 1 && args[0] != "" {
// 			userId = args[0]
// 		} else {
// 			fmt.Printf("Invalid args")
// 			return
// 		}
// 		balance, err := core.GetBalance(userId, BalanceFile)
// 		if err != nil {
// 			fmt.Printf("Failed to get balance of %s: %s", userId, err.Error())
// 			return
// 		}
//
// 		fmt.Printf("%s's balance: %d", userId, balance)
// 	},
// }

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// func LoadKeys(id string) (*rsa.PrivateKey, error) {
// 	// bytes := make([]byte, length)
// 	// if _, err := rand.Read(bytes); err != nil {
// 	// 	panic(err)
// 	// }
// 	// return hex.EncodeToString(bytes)
// 	return nil, nil
// }
//
// func loadKeyFromFile(filename string) (*rsa.PrivateKey, error) {
// 	keyBytes, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	block, _ := pem.Decode(keyBytes)
// 	if block == nil {
// 		return nil, fmt.Errorf("failed to decode PEM block")
// 	}
//
// 	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return key, nil
// }

// getCmd represents the get command
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
			r := GenerateRandomString(16)
			hashed := fdh.Sum(crypto.SHA256, HashSize, []byte(r))

			// Blind the hashed message
			blinded, unblinder, err := rsablind.Blind(pubKey, hashed)
			if err != nil {
				fmt.Println("Blind message failed")
				return
			}

			err = core.InsertNewElement(
				DataFile,
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

// var unblindTicket = &cobra.Command{
// 	Use:   "unblind-ticket",
// 	Short: "This command will get the desired Gopher",
// 	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		var amount = 0
//
// 		if len(args) >= 1 && args[0] != "" {
// 			value, err := strconv.Atoi(args[0])
// 			if err != nil {
// 				fmt.Println("Invalid arg")
// 				return
// 			}
// 			if value%SmallestDevision != 0 {
// 				fmt.Println("Amount must be mutiples of 50.000")
// 				return
// 			}
// 			amount = value
// 		}
// 		// Generate a key
// 		key, _ := loadKeyFromFile(PrivKeyFile)
//
// 		outputFile, err := os.Create("blinded_tickets.txt")
// 		if err != nil {
// 			fmt.Println("Failed to create output file:", err)
// 			return
// 		}
// 		defer outputFile.Close()
//
// 		fmt.Printf("Generate %d tickets to transfer %dvnd\n", amount/SmallestDevision, amount)
// 		for i := 0; i < amount/SmallestDevision; i++ {
// 			r := GenerateRandomString(16)
// 			hashed := fdh.Sum(crypto.SHA256, HashSize, []byte(r))
//
// 			// Blind the hashed message
// 			blinded, _, err := rsablind.Blind(&key.PublicKey, hashed)
// 			if err != nil {
// 				fmt.Println("Blind message failed")
// 				return
// 			}
//
// 			blindedHex := hex.EncodeToString(blinded)
// 			_, err = outputFile.WriteString(blindedHex + "\n")
// 			if err != nil {
// 				fmt.Println("Failed to write to output file:", err)
// 				return
// 			}
// 		}
// 	},
// }
