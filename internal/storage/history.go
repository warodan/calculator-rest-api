package storage

import "sync"

type Entry struct {
	FirstNumber  int
	SecondNumber int
	Operation    string
	Result       int
}

type History struct {
	mu     sync.Mutex
	values map[string][]Entry
}

func NewHistory() *History {
	return &History{
		values: make(map[string][]Entry),
	}
}

func (history *History) Add(token string, entry Entry) {
	history.mu.Lock()
	defer history.mu.Unlock()

	history.values[token] = append(history.values[token], entry)
}

// Additional methods

func (history *History) All(token string) []Entry {
	history.mu.Lock()
	defer history.mu.Unlock()

	entries := history.values[token]
	return append([]Entry(nil), entries...)
}

func (history *History) AllTokens() []string {
	history.mu.Lock()
	defer history.mu.Unlock()

	tokens := make([]string, 0, len(history.values))
	for token := range history.values {
		tokens = append(tokens, token)
	}
	return tokens
}

func (history *History) Clear(token string) {
	history.mu.Lock()
	defer history.mu.Unlock()

	delete(history.values, token)
}
