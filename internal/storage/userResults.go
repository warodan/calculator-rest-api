package storage

import "sync"

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

func (userResults *UserResults) Add(token string, entry Entry) {
	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	userResults.values[token] = append(userResults.values[token], entry)
}

// Additional methods

func (userResults *UserResults) All(token string) []Entry {
	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	entries := userResults.values[token]
	return append([]Entry(nil), entries...)
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

func (userResults *UserResults) Clear(token string) {
	userResults.mu.Lock()
	defer userResults.mu.Unlock()

	delete(userResults.values, token)
}
