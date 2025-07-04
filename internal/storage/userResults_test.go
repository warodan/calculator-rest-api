package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserResults_AddAndAll(t *testing.T) {
	store := NewUserStorage()
	token := "test-token"
	entry := Entry{FirstNumber: 2, SecondNumber: 3, Operation: "sum", Result: 5}

	err := store.Add(token, entry)
	require.NoError(t, err)

	entries, err := store.All(token)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, entry, entries[0])
}

func TestUserResults_Clear(t *testing.T) {
	store := NewUserStorage()
	token := "user1"
	_ = store.Add(token, Entry{FirstNumber: 1, SecondNumber: 2, Operation: "sum", Result: 3})

	err := store.Clear(token)
	require.NoError(t, err)

	entries, err := store.All(token)
	require.NoError(t, err)
	require.Len(t, entries, 0)
}

func TestUserResults_Clear_MissingToken(t *testing.T) {
	store := NewUserStorage()
	err := store.Clear("non-existent")
	require.Error(t, err)
}

func TestUserResults_AllTokens(t *testing.T) {
	store := NewUserStorage()
	_ = store.Add("a", Entry{1, 2, "sum", 3})
	_ = store.Add("b", Entry{2, 3, "sum", 5})

	tokens := store.AllTokens()
	require.ElementsMatch(t, []string{"a", "b"}, tokens)
}
