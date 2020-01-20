package flesch

import "strings"

// Sentence is encountered whenever you find a word that
// ends in a specific punctuation symbol: a period, question
// mark, or exclamation point.
type Sentence struct {
	allRunes []rune
	Start    int
	End      int
	Words    []Word
}

func (s Sentence) Runes() []rune {
	return s.allRunes[s.Start : s.End+1]
}

func (s Sentence) String() string {
	var b strings.Builder
	for _, r := range s.Runes() {
		b.WriteRune(r)
	}
	return b.String()
}

// Word is contiguous sequence of alphabetic characters.
// Whitespace defines word boundaries.
type Word struct {
	allRunes []rune
	Start    int
	End      int
}

func (w Word) Runes() []rune {

	return w.allRunes[w.Start : w.End+1]
}

func (w Word) String() string {
	var b strings.Builder
	for _, r := range w.Runes() {
		b.WriteRune(r)
	}
	return b.String()
}

// Syllables are considered to have been encountered whenever
// you detect a vowel at the start of a word or a vowel
// following a consonant in a word. A lone ‘e’ at the end
// of a word does not count as a syllable.
func (w Word) Syllables() int {
	var syllables int
	word := w.String()

	// vowel at start of word
	if TypeOfRune(rune(word[0])) == RuneTypeVowel {
		syllables++
	}

	// ...or a vowel following a consonant in a word
	for i, r := range word {
		// Not relevant for first character
		if i == 0 {
			continue
		}
		if TypeOfRune(r) == RuneTypeVowel && TypeOfRune(rune(word[i-1])) == RuneTypeConsonant {
			// do not count if last character and a LONE 'e'
			if i == /* last: */ len(word)-1 && /* is 'e': */ r == 'e' && /* lone: */ rune(word[i-1]) != 'e' {
				continue
			}
			// otherwise, count it
			syllables++
		}
	}

	return 0
}

type RuneType int

func TypeOfRune(r rune) RuneType {
	runeType, ok := runeLookup[r]
	if !ok {
		return RuneTypeOther
	}

	return runeType
}

var runeLookup = map[rune]RuneType{
	// Numbers
	'0': RuneTypeNumber,
	'1': RuneTypeNumber,
	'2': RuneTypeNumber,
	'3': RuneTypeNumber,
	'4': RuneTypeNumber,
	'5': RuneTypeNumber,
	'6': RuneTypeNumber,
	'7': RuneTypeNumber,
	'8': RuneTypeNumber,
	'9': RuneTypeNumber,

	// Lowercase
	'a': RuneTypeVowel,
	'b': RuneTypeConsonant,
	'c': RuneTypeConsonant,
	'd': RuneTypeConsonant,
	'e': RuneTypeVowel,
	'f': RuneTypeConsonant,
	'g': RuneTypeConsonant,
	'h': RuneTypeConsonant,
	'i': RuneTypeVowel,
	'j': RuneTypeConsonant,
	'k': RuneTypeConsonant,
	'l': RuneTypeConsonant,
	'm': RuneTypeConsonant,
	'n': RuneTypeConsonant,
	'o': RuneTypeVowel,
	'p': RuneTypeConsonant,
	'q': RuneTypeConsonant,
	'r': RuneTypeConsonant,
	's': RuneTypeConsonant,
	't': RuneTypeConsonant,
	'u': RuneTypeVowel,
	'v': RuneTypeConsonant,
	'w': RuneTypeConsonant,
	'x': RuneTypeConsonant,
	'y': RuneTypeConsonant,
	'z': RuneTypeConsonant,

	// Uppercase
	'A': RuneTypeVowel,
	'B': RuneTypeConsonant,
	'C': RuneTypeConsonant,
	'D': RuneTypeConsonant,
	'E': RuneTypeVowel,
	'F': RuneTypeConsonant,
	'G': RuneTypeConsonant,
	'H': RuneTypeConsonant,
	'I': RuneTypeVowel,
	'J': RuneTypeConsonant,
	'K': RuneTypeConsonant,
	'L': RuneTypeConsonant,
	'M': RuneTypeConsonant,
	'N': RuneTypeConsonant,
	'O': RuneTypeVowel,
	'P': RuneTypeConsonant,
	'Q': RuneTypeConsonant,
	'R': RuneTypeConsonant,
	'S': RuneTypeConsonant,
	'T': RuneTypeConsonant,
	'U': RuneTypeVowel,
	'V': RuneTypeConsonant,
	'W': RuneTypeConsonant,
	'X': RuneTypeConsonant,
	'Y': RuneTypeConsonant,
	'Z': RuneTypeConsonant,

	// Whitespace
	' ':  RuneTypeWhiteSpace,
	'\t': RuneTypeWhiteSpace,
	'\n': RuneTypeWhiteSpace,
	'\r': RuneTypeWhiteSpace,

	// Word Stop
	'”': RuneTypeWordStop,
	'"': RuneTypeWordStop,
	',': RuneTypeWordStop,
	')': RuneTypeWordStop,

	// Sentence Stop
	'.': RuneTypeSentenceStop,
	';': RuneTypeSentenceStop,
	'!': RuneTypeSentenceStop,
	'?': RuneTypeSentenceStop,
}

const (
	RuneTypeSentenceStop RuneType = iota
	RuneTypeWhiteSpace
	RuneTypeWordStop
	RuneTypeVowel
	RuneTypeConsonant
	RuneTypeNumber
	RuneTypeOther
)
