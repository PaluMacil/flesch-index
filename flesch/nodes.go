package flesch

import "strings"

type Document struct {
	Sentences []Sentence
	name      string
}

func (d Document) Name() string {
	if d.name == "" {
		return "(no name)"
	}

	return d.name
}

func (d Document) WordCount() int {
	var count int
	for _, sentence := range d.Sentences {
		count += len(sentence.Words)
	}

	return count
}

func (d Document) Words() []Word {
	var words []Word
	for _, sentence := range d.Sentences {
		words = append(words, sentence.Words...)
	}

	return words
}

func (d Document) UniqueWords() []Word {
	keys := make(map[string]bool)
	var list []Word
	for _, word := range d.Words() {
		// case invariant
		wordString := strings.ToUpper(word.String())
		if _, exists := keys[wordString]; !exists {
			keys[wordString] = true
			list = append(list, word)
		}
	}
	return list
}

func (d Document) Syllables() int {
	var count int
	for _, s := range d.Sentences {
		count += s.Syllables()
	}

	return count
}

func (d Document) Score() float32 {
	syllables := float32(d.Syllables())
	words := float32(d.WordCount())
	sentences := float32(len(d.Sentences))

	avgWordPerSen := words / sentences
	// average syllables per word
	avgSylPerWord := syllables / words

	score := 206.835 - (84.6 * avgSylPerWord) - (1.015 * avgWordPerSen)

	return score
}

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

func (s Sentence) Syllables() int {
	var count int
	for _, w := range s.Words {
		count += w.Syllables()
	}

	return count
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
// of a word does not count as a syllable. Three letter words
// or less are always one syllable.
func (w Word) Syllables() int {
	word := w.Runes()

	return syllablesFromRunes(word)
}

func syllablesFromRunes(runes []rune) int {
	var syllables int
	if len(runes) <= 3 {
		return 1
	}

	// vowel at start of runes
	if TypeOfRune(runes[0]) == RuneTypeVowel {
		syllables++
	}

	// ...or a vowel following a consonant in a runes
	for i, r := range runes {
		// Not relevant for first rune
		if i == 0 {
			continue
		}
		if TypeOfRune(r) == RuneTypeVowel && TypeOfRune(runes[i-1]) == RuneTypeConsonant {
			// do not count if last rune and a LONE 'e'
			if i == /* last: */ len(runes)-1 && /* is 'e': */ r == 'e' && /* lone: */ runes[i-1] != 'e' {
				continue
			}
			// otherwise, count it
			syllables++
		}
	}

	// all words are at least one syllable
	if syllables == 0 {
		syllables = 1
	}

	return syllables
}

func SyllablesFromString(word string) int {
	return syllablesFromRunes([]rune(word))
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
	':': RuneTypeWordStop,

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
