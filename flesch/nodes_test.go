package flesch_test

import (
	"github.com/PaluMacil/flesch-index/flesch"
	"testing"
)

func TestTypeOfRune(t *testing.T) {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for _, vowel := range vowels {
		r := []rune(vowel)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeVowel {
			t.Errorf("For %v, didn't get vowel", string(r))
		}
	}

	consonants := []string{"b", "B", "L", "f", "T", "Q"}
	for _, consonant := range consonants {
		r := []rune(consonant)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeConsonant {
			t.Errorf("For %v, didn't get consonant", string(r))
		}
	}

	whitespace := []string{"\n", "\t", " ", "\r"}
	for _, ws := range whitespace {
		r := []rune(ws)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeWhiteSpace {
			t.Errorf("For %v, didn't get whitespace", string(r))
		}
	}

	wordStops := []string{"”", "\"", ",", ")"}
	for _, ws := range wordStops {
		r := []rune(ws)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeWordStop {
			t.Errorf("For %s, didn't get word stop", string(r))
		}
	}

	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, number := range numbers {
		r := []rune(number)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeNumber {
			t.Errorf("For %s, didn't get number", string(r))
		}
	}

	sentenceStop := []string{".", ";", "!", "?"}
	for _, stop := range sentenceStop {
		r := []rune(stop)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeSentenceStop {
			t.Errorf("For %s, didn't get sentence stop", string(r))
		}
	}

	others := []string{"語", "—"}
	for _, stop := range others {
		r := []rune(stop)[0]
		if flesch.TypeOfRune(r) != flesch.RuneTypeOther {
			t.Errorf("For %s, didn't get other", string(r))
		}
	}
}

type SyllableTestResult struct {
	Word     string
	Expected int
}

func TestSyllablesFromString(t *testing.T) {
	testCases := []SyllableTestResult{
		{"car", 1},
		{"test", 1},
		{"beer", 1},
		{"care", 1},
		{"carrot", 2},
		{"consecrate", 3},
		{"abraham", 3},
	}
	for _, test := range testCases {
		result := flesch.SyllablesFromString(test.Word)
		if test.Expected != result {
			t.Errorf("%s: expected %d syllables, got %d", test.Word, test.Expected, result)
		}
	}
}
