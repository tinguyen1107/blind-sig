package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserBalance struct {
	UserID  string `json:"user_id"`
	Balance int    `json:"balance"`
}

type UserBalances []UserBalance

func UpdateBalance(userID string, newBalance int, filename string) error {
	// Read the existing file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal JSON data
	var balances UserBalances
	if err := json.Unmarshal(data, &balances); err != nil {
		return err
	}

	// Update the balance
	found := false
	for i, userBalance := range balances {
		if userBalance.UserID == userID {
			balances[i].Balance = newBalance
			found = true
			break
		}
	}

	if !found {
		balances = append(balances, UserBalance{UserID: userID, Balance: newBalance})
	}

	// Marshal back to JSON
	data, err = json.Marshal(balances)
	if err != nil {
		return err
	}

	// Write back to file
	return ioutil.WriteFile(filename, data, os.ModePerm)
}

func GetBalance(userID string, filename string) (int, error) {
	// Read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	// Unmarshal JSON data
	var balances UserBalances
	if err := json.Unmarshal(data, &balances); err != nil {
		return 0, err
	}

	// Find and return the balance
	for _, userBalance := range balances {
		if userBalance.UserID == userID {
			return userBalance.Balance, nil
		}
	}

	return 0, fmt.Errorf("User ID not found")
}
