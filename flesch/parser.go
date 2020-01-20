package flesch

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type DocumentReport struct {
	Sentences []Sentence
}

func ParseFile(filename string) (DocumentReport, error) {
	rawData, err := ioutil.ReadFile(filename)
	if err != nil {
		return DocumentReport{}, fmt.Errorf("reading %s: %w", filename, err)
	}
	return ParseBytes(rawData)
}

func ParseBytes(bytes []byte) (DocumentReport, error) {
	var report DocumentReport

	runes := []rune(string(bytes))
	var currentRuneIndex int
	for {
		sentence, err := GetSentence(runes, currentRuneIndex)
		if err != nil {
			break
		}
		report.Sentences = append(report.Sentences, sentence)
		currentRuneIndex = sentence.End + 1
	}

	return report, nil
}

var NoMoreSentences = errors.New("no more sentences")

func GetSentence(allRunes []rune, start int) (Sentence, error) {
	i := start
	sentence := Sentence{allRunes: allRunes}
	var sentenceStarted bool
	for {
		if i >= len(allRunes) {
			return sentence, NoMoreSentences
		}
		r := allRunes[i]
		// if the sentence hasn't started yet...
		if !sentenceStarted {
			// if the rune is a vowel or consonant, the sentence will have started here
			if TypeOfRune(r) == RuneTypeVowel || TypeOfRune(r) == RuneTypeConsonant {
				sentenceStarted = true
				sentence.Start = i
			}
		} else {
			if TypeOfRune(r) == RuneTypeSentenceStop {
				sentence.End = i
				return sentence, nil
			}
		}

		i++
	}
}
