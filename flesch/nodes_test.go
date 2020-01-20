package flesch_test

import (
	"github.com/PaluMacil/flesch-index/flesch"
	"testing"
)

func TestTypeOfRune(t *testing.T) {
	vowels := []byte{[]byte("a")[0], []byte("e")[0], []byte("i")[0], []byte("o")[0], []byte("u")[0], []byte("A")[0],
		[]byte("E")[0], []byte("I")[0], []byte("O")[0], []byte("U")[0]}
	for _, vowel := range vowels {
		token := flesch.Tokenize(vowel)
		if flesch.TypeOfRune(token) != flesch.RuneTypeVowel {
			t.Errorf("For %v, didn't get vowel", token.Value)
		}
	}

	consonants := []byte{[]byte("b")[0], []byte("B")[0], []byte("L")[0], []byte("f")[0], []byte("T")[0], []byte("Q")[0]}
	for _, consonant := range consonants {
		token := flesch.Tokenize(consonant)
		if flesch.TypeOfRune(token) != flesch.RuneTypeConsonant {
			t.Errorf("For %v, didn't get consonant", token.Value)
		}
	}

	whitespace := []byte{[]byte("\n")[0], []byte("\t")[0], []byte(" ")[0], []byte("\r")[0]}
	for _, ws := range whitespace {
		token := flesch.Tokenize(ws)
		if flesch.TypeOfRune(token) != flesch.RuneTypeWhiteSpace {
			t.Errorf("For %v, didn't get whitespace", token.Value)
		}
	}

	// TODO: numbers, other, and sentence enders

}
