package flesch_test

import (
	"github.com/PaluMacil/flesch-index/flesch"
	"testing"
)

func TestParseBytes(t *testing.T) {
	testWords := `
		Twinkle, twinkle, little star. I am not a teapot. Carrots are healthy. My name is Bob! Why?
	`
	report, err := flesch.ParseBytes([]byte(testWords))
	if err != nil {
		t.Errorf("parsing: %s", err)
	}
	if len(report.Sentences) != 5 {
		t.Errorf("expected 5 sentences, got %d", len(report.Sentences))
	}
}

func TestSentences(t *testing.T) {
	testWords := `
		Twinkle, twinkle, little star. I am not a teapot. Carrots are healthy. My name is Bob! Why?
	`
	report, err := flesch.ParseBytes([]byte(testWords))
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
