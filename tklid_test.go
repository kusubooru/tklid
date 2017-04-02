package tklid

import (
	"strings"
	"testing"
)

const seed = 42

func TestValidate(t *testing.T) {
	id := New(seed)
	if got, want := Validate(id), true; got != want {
		t.Errorf("invalid id: %s", id)
	}
}

func TestValidate_wrongSize(t *testing.T) {
	in := "foo"
	if got, want := Validate(in), false; got != want {
		t.Errorf("Validate(%q) = %t, want %t", in, got, want)
	}
}

func TestValidate_unknownWord(t *testing.T) {
	var unknownWords = make([]string, 0, len(allWords))
	for i := 0; i < len(allWords); i++ {
		unknownWords = append(unknownWords, "foo")
	}
	in := strings.Join(unknownWords, "-")
	if got, want := Validate(in), false; got != want {
		t.Errorf("Validate(%q) = %t, want %t", in, got, want)
	}
}

const (
	fairNumberOfCombos int64 = 10000000000 // 10 billion
)

func TestValidate_unsafeCharacterLength(t *testing.T) {
	id := strings.Repeat("f", safeIDLength+1)
	if got, want := Validate(id), false; got != want {
		t.Errorf("Validate(%q) = %t, want %t", id, got, want)
	}
}

func TestEnoughCombos(t *testing.T) {
	if got, want := combos(), fairNumberOfCombos; got < want {
		t.Errorf("Number of possible combinations = %d which is less than %d", got, want)
	}
}

func combos() int64 {
	combos := int64(1)
	for i := range allWords {
		combos *= int64(len(allWords[i]))
	}
	return combos
}

func TestMaxLength(t *testing.T) {
	longestID := findLongestID()
	if len(longestID) > safeIDLength {
		t.Errorf("longestID: %q has length %d which exceeds safe length of %d", longestID, len(longestID), safeIDLength)
	}
}

func findLongestID() string {
	longestWords := make([]string, 0, len(allWords))
	for i := range allWords {
		longestWords = append(longestWords, findLongestWord(allWords[i]))
	}
	return strings.Join(longestWords, "-")
}

func findLongestWord(words words) string {
	max := 0
	longest := ""
	for _, word := range words {
		if len(word) > max {
			max = len(word)
			longest = word
		}
	}
	return longest
}
