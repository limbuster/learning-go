package main

import (
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, want string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is a new test"
	dictionary := Dictionary{word: definition}
	t.Run("Known word", func(t *testing.T) {
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a new test"
		dictionary := Dictionary{word: "this is just a test"}
		dictionary.Update(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Unknown word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
