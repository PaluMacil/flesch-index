package flesch

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type DocumentReport struct {
	Sentences    []Sentence
	currentToken int
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

	tokens := make([]Token, len(bytes))
	for i, b := range bytes {
		tokens[i] = Tokenize(b)
	}
	var currentTokenIndex int
	for {
		sentence, err := GetSentence(tokens, currentTokenIndex)
		if err != nil {
			break
		}
		report.Sentences = append(report.Sentences, sentence)
		currentTokenIndex = sentence.End + 1
	}

	return report, nil
}

var NoMoreSentences = errors.New("no more sentences")

func GetSentence(allTokens []Token, start int) (Sentence, error) {
	i := start
	sentence := Sentence{allTokens: allTokens}
	var sentenceStarted bool
	for {
		if i >= len(allTokens) {
			return sentence, NoMoreSentences
		}
		token := allTokens[i]
		// if the sentence hasn't started yet...
		if !sentenceStarted {
			// if the token is a vowel or consonant, the sentence will have started here
			if token.Type == TokenTypeVowel || token.Type == TokenTypeConsonant {
				sentenceStarted = true
				sentence.Start = i
			}
		} else {
			if token.Type == TokenTypeSentenceStop {
				sentence.End = i
				return sentence, nil
			}
		}

		i++
	}
}
