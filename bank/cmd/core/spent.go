package core

import (
	"encoding/json"
	"os"
)

type Tickets []string

func IsTicketExist(ticket string, filename string) (bool, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// Unmarshal JSON data
	var tickets Tickets
	if err := json.Unmarshal(data, &tickets); err != nil {
		return false, err
	}

	// Check if the ticket exists
	for _, t := range tickets {
		if t == ticket {
			return true, nil
		}
	}

	return false, nil
}

func AddTicket(ticket string, filename string) error {
	// Read the existing file
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal JSON data
	var tickets Tickets
	if err := json.Unmarshal(data, &tickets); err != nil {
		return err
	}

	// Add the ticket
	tickets = append(tickets, ticket)

	// Marshal back to JSON
	data, err = json.Marshal(tickets)
	if err != nil {
		return err
	}

	// Write back to file
	return os.WriteFile(filename, data, os.ModePerm)
}
