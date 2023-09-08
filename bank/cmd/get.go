package cmd

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"strconv"

	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
	"github.com/spf13/cobra"
)

const (
	SmallestDevision = 50000
	KeySize          = 2048
	HashSize         = 1536
	PrivKeyFile      = "private_key.pem"
	PubKeyFile       = "public_key.pem"
)

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func LoadKeys(id string) (*rsa.PrivateKey, error) {
	// bytes := make([]byte, length)
	// if _, err := rand.Read(bytes); err != nil {
	// 	panic(err)
	// }
	// return hex.EncodeToString(bytes)
	return nil, nil
}

func saveKeyToFile(key *rsa.PrivateKey, filename string) error {
	privateKeyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		},
	)
	_, err = privateKeyFile.Write(privKeyPEM)
	return err
}

func savePublicKeyToFile(pubKey *rsa.PublicKey, filename string) error {
	publicKeyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}

	pubKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)

	_, err = publicKeyFile.Write(pubKeyPEM)
	return err
}

func loadKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

var generateKeys = &cobra.Command{
	Use:   "gen-key",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := rsa.GenerateKey(rand.Reader, KeySize)
		if err != nil {
			fmt.Println("Failed to generate key:", err)
			return
		}
		err = saveKeyToFile(key, PrivKeyFile)
		if err != nil {
			fmt.Println("Failed to save private key:", err)
			return
		}

		err = savePublicKeyToFile(&key.PublicKey, PubKeyFile)
		if err != nil {
			fmt.Println("Failed to save public key:", err)
			return
		}
	},
}

// getCmd represents the get command
var generateTicket = &cobra.Command{
	Use:   "gen-ticket",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
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
		// Generate a key
		key, _ := loadKeyFromFile(PrivKeyFile)

		outputFile, err := os.Create("blinded_tickets.txt")
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			return
		}
		defer outputFile.Close()

		fmt.Printf("Generate %d tickets to transfer %dvnd\n", amount/SmallestDevision, amount)
		for i := 0; i < amount/SmallestDevision; i++ {
			r := GenerateRandomString(16)
			hashed := fdh.Sum(crypto.SHA256, HashSize, []byte(r))

			// Blind the hashed message
			blinded, _, err := rsablind.Blind(&key.PublicKey, hashed)
			if err != nil {
				fmt.Println("Blind message failed")
				return
			}

			blindedHex := hex.EncodeToString(blinded)
			_, err = outputFile.WriteString(blindedHex + "\n")
			if err != nil {
				fmt.Println("Failed to write to output file:", err)
				return
			}
		}
	},
}

var unblindTicket = &cobra.Command{
	Use:   "unblind-ticket",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
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
		// Generate a key
		key, _ := loadKeyFromFile(PrivKeyFile)

		outputFile, err := os.Create("blinded_tickets.txt")
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			return
		}
		defer outputFile.Close()

		fmt.Printf("Generate %d tickets to transfer %dvnd\n", amount/SmallestDevision, amount)
		for i := 0; i < amount/SmallestDevision; i++ {
			r := GenerateRandomString(16)
			hashed := fdh.Sum(crypto.SHA256, HashSize, []byte(r))

			// Blind the hashed message
			blinded, _, err := rsablind.Blind(&key.PublicKey, hashed)
			if err != nil {
				fmt.Println("Blind message failed")
				return
			}

			blindedHex := hex.EncodeToString(blinded)
			_, err = outputFile.WriteString(blindedHex + "\n")
			if err != nil {
				fmt.Println("Failed to write to output file:", err)
				return
			}
		}
	},
}

func init() {
	// Generate keypair
	// Blind sign to blinded message
	// Check Double spending
	// Spend
	rootCmd.AddCommand(generateTicket, generateKeys)
}
