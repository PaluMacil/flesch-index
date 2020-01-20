package flesch_test

import (
	"fmt"
	"github.com/PaluMacil/flesch-index/flesch"
	"path"
	"testing"
)

func TestParseString(t *testing.T) {
	testWords := `
		Twinkle, twinkle, little star. I am not a teapot. Carrots are healthy. My name is Bob! Why?
	`
	report, err := flesch.ParseString(testWords, "testWords")
	if err != nil {
		t.Errorf("parsing: %s", err)
	}
	if len(report.Sentences) != 5 {
		t.Errorf("expected 5 sentences, got %d", len(report.Sentences))
	}

	wc := report.WordCount()
	if wc != 17 {
		t.Errorf("expected word count of 17, got %d", wc)
	}
}

func loadTestDocuments(t *testing.T) []flesch.Document {
	gettysburgDocument, err := flesch.ParseFile(path.Join("..", "GettysburgAddress.txt"))
	if err != nil {
		t.Errorf("Loading GettysburgAddress.txt: %s", err)
	}
	mobyDickDocument, err := flesch.ParseFile(path.Join("..", "MobyDick.txt"))
	if err != nil {
		t.Errorf("Loading MobyDick.txt: %s", err)
	}
	nyTimesDocument, err := flesch.ParseFile(path.Join("..", "NYTimes.txt"))
	if err != nil {
		t.Errorf("Loading NYTimes.txt: %s", err)
	}
	return []flesch.Document{
		gettysburgDocument,
		mobyDickDocument,
		nyTimesDocument,
	}
}

func TestScoreWithinRange(t *testing.T) {
	for _, d := range loadTestDocuments(t) {
		score := d.Score()
		fmt.Println("Doc:", d.Name(), "Score:", score)
		if score < 0 || score > 100 {
			t.Errorf("expected score 1 to 100 for %s, got %v", d.Name(), score)
		}
	}
}

func TestNoZeroLengths(t *testing.T) {
	for _, d := range loadTestDocuments(t) {
		numberOfSentences := len(d.Sentences)
		if numberOfSentences == 0 {
			t.Errorf("expected more than zero sentences in %s", d.Name())
		}

		for sentenceIndex, s := range d.Sentences {
			numberOfWords := len(s.Words)
			if numberOfWords == 0 {
				t.Errorf("expected more than zero words in sentence #%d of %s", sentenceIndex, d.Name())
			}

			for wordIndex, w := range s.Words {
				numberOfSyllables := w.Syllables()
				if numberOfSyllables == 0 {
					t.Errorf("expected more than zero syllables in word #%d ('%s') of sentence #%d of %s",
						wordIndex, w, sentenceIndex, d.Name())
				}
			}
		}
	}
}

func TestSentences(t *testing.T) {
	testWords := `
		Twinkle, twinkle, little star. I am not a teapot. Carrots are healthy. My name is Bob! Why?
	`
	report, err := flesch.ParseString(testWords, "testWords")
	if err != nil {
		t.Errorf("parsing: %s", err)
	}
	if report.Sentences[0].String() != "Twinkle, twinkle, little star." {
		t.Errorf("first sentence '', got %s", report.Sentences[0])
	}
	if report.Sentences[1].String() != "I am not a teapot." {
		t.Errorf("second sentence '', got %s", report.Sentences[1])
	}
	if report.Sentences[2].String() != "Carrots are healthy." {
		t.Errorf("third sentence '', got %s", report.Sentences[2])
	}
	if report.Sentences[3].String() != "My name is Bob!" {
		t.Errorf("fourth sentence '', got %s", report.Sentences[3])
	}
	if report.Sentences[4].String() != "Why?" {
		t.Errorf("fifth sentence '', got %s", report.Sentences[4])
	}
}
