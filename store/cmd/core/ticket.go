package core

import (
	"encoding/json"
	"errors"
	"os"
)

type Ticket struct {
	BlindedMessage       string `json:"blined_message"`
	Unblinder            string `json:"unblinder"`
	SignedBlindedMessage string `json:"signed_blinded_message"`
	SignedMessage        string `json:"signed_message"`
}

func InsertNewElement(filename string, blindedMessage, unblinder string) error {
	var tickets []Ticket

	// Read existing tickets from file
	fileContent, err := os.ReadFile(filename)
	if err == nil {
		if err := json.Unmarshal(fileContent, &tickets); err != nil {
			return err
		}
	}

	newTicket := Ticket{
		BlindedMessage: blindedMessage,
		Unblinder:      unblinder,
	}
	tickets = append(tickets, newTicket)

	// Save updated tickets back to file
	newFileContent, err := json.Marshal(tickets)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, newFileContent, 0644)
}

func UpdateElementByIndex(filename string, index int, signedBlindedMessage, signedMessage string) error {
	var tickets []Ticket

	// Read existing tickets from file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileContent, &tickets); err != nil {
		return err
	}

	if index < 0 || index >= len(tickets) {
		return errors.New("index out of range")
	}
	tickets[index].SignedBlindedMessage = signedBlindedMessage
	tickets[index].SignedMessage = signedMessage

	// Save updated tickets back to file
	newFileContent, err := json.Marshal(tickets)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, newFileContent, 0644)
}

func GetIncompleteElement(filename string) (*Ticket, int, error) {
	var tickets []Ticket

	// Read existing tickets from file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, -1, err
	}
	if err := json.Unmarshal(fileContent, &tickets); err != nil {
		return nil, -1, err
	}

	for i, ticket := range tickets {
		if ticket.SignedBlindedMessage == "" && ticket.SignedMessage == "" {
			return &ticket, i, nil
		}
	}
	return nil, -1, errors.New("no incomplete element found")
}

func GetTicket(filename string, index int) (*Ticket, error) {
	var tickets []Ticket

	// Read existing tickets from file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileContent, &tickets); err != nil {
		return nil, err
	}

	if index < 0 || index >= len(tickets) {
		return nil, errors.New("index out of range")
	}

	return &tickets[index], nil
}
