package storage

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Entry struct {
	FirstNumber  int
	SecondNumber int
	Operation    string
	Result       int
}

type UserResults struct {
	mu     sync.Mutex
	values map[string][]Entry
}

func NewUserStorage() *UserResults {
	return &UserResults{
		values: make(map[string][]Entry),
	}
}

func validateToken(token string) error {
	if token == "" {
		return fmt.Errorf("token is empty")
	}

	if _, err := uuid.Parse(token); err != nil {
		return fmt.Errorf("token is not valid UUID: %w", err)
	}

	return nil
}

func (userResults *UserResults) Add(token string, entry Entry) error {
	if err := validateToken(token); err != nil {
		return err
	}

	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	userResults.values[token] = append(userResults.values[token], entry)

	return nil
}

// Additional methods

func (userResults *UserResults) All(token string) ([]Entry, error) {
	if err := validateToken(token); err != nil {
		return nil, err
	}

	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	entries := userResults.values[token]
	return append(make([]Entry, 0, len(entries)), entries...), nil
}

func (userResults *UserResults) AllTokens() []string {
	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	tokens := make([]string, 0, len(userResults.values))
	for token := range userResults.values {
		tokens = append(tokens, token)
	}
	return tokens
}

func (userResults *UserResults) Clear(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}

	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	if _, ok := userResults.values[token]; !ok {
		return fmt.Errorf("token %q not found", token)
	}

	delete(userResults.values, token)
	return nil
}
