package flesch

// Sentence is encountered whenever you find a word that
// ends in a specific punctuation symbol: a period, question
// mark, or exclamation point.
type Sentence struct {
}

// Word is contiguous sequence of alphabetic characters.
// Whitespace defines word boundaries.
type Word struct {
}

// Syllable is considered to have been encountered whenever
// you detect a vowel at the start of a word or a vowel
// following a consonant in a word. A lone ‘e’ at the end
// of a word does not count as a syllable.
type Syllable struct {
}

type Character struct {
}

type TokenType int

func RuneToTokenType(r rune) TokenType {
	tokenType, ok := tokenTypeLookup[r]
	if !ok {
		return TokenTypeOther
	}
	return tokenType
}

var tokenTypeLookup = map[rune]TokenType{
	// Numbers
	'0': TokenTypeNumber,
	'1': TokenTypeNumber,
	'2': TokenTypeNumber,
	'3': TokenTypeNumber,
	'4': TokenTypeNumber,
	'5': TokenTypeNumber,
	'6': TokenTypeNumber,
	'7': TokenTypeNumber,
	'8': TokenTypeNumber,
	'9': TokenTypeNumber,

	// Lowercase
	'a': TokenTypeVowel,
	'b': TokenTypeConsonant,
	'c': TokenTypeConsonant,
	'd': TokenTypeConsonant,
	'e': TokenTypeVowel,
	'f': TokenTypeConsonant,
	'g': TokenTypeConsonant,
	'h': TokenTypeConsonant,
	'i': TokenTypeVowel,
	'j': TokenTypeConsonant,
	'k': TokenTypeConsonant,
	'l': TokenTypeConsonant,
	'm': TokenTypeConsonant,
	'n': TokenTypeConsonant,
	'o': TokenTypeVowel,
	'p': TokenTypeConsonant,
	'q': TokenTypeConsonant,
	'r': TokenTypeConsonant,
	's': TokenTypeConsonant,
	't': TokenTypeConsonant,
	'u': TokenTypeVowel,
	'v': TokenTypeConsonant,
	'w': TokenTypeConsonant,
	'x': TokenTypeConsonant,
	'y': TokenTypeConsonant,
	'z': TokenTypeConsonant,

	// Uppercase
	'A': TokenTypeVowel,
	'B': TokenTypeConsonant,
	'C': TokenTypeConsonant,
	'D': TokenTypeConsonant,
	'E': TokenTypeVowel,
	'F': TokenTypeConsonant,
	'G': TokenTypeConsonant,
	'H': TokenTypeConsonant,
	'I': TokenTypeVowel,
	'J': TokenTypeConsonant,
	'K': TokenTypeConsonant,
	'L': TokenTypeConsonant,
	'M': TokenTypeConsonant,
	'N': TokenTypeConsonant,
	'O': TokenTypeVowel,
	'P': TokenTypeConsonant,
	'Q': TokenTypeConsonant,
	'R': TokenTypeConsonant,
	'S': TokenTypeConsonant,
	'T': TokenTypeConsonant,
	'U': TokenTypeVowel,
	'V': TokenTypeConsonant,
	'W': TokenTypeConsonant,
	'X': TokenTypeConsonant,
	'Y': TokenTypeConsonant,
	'Z': TokenTypeConsonant,

	// Whitespace
	' ':  TokenTypeWhiteSpace,
	'\t': TokenTypeWhiteSpace,
	'\n': TokenTypeWhiteSpace,
	'\r': TokenTypeWhiteSpace,

	// Sentence Stop
	'.': TokenTypeSentenceStop,
	':': TokenTypeSentenceStop,
	'!': TokenTypeSentenceStop,
	'?': TokenTypeSentenceStop,
}

const (
	TokenTypeSentenceStop TokenType = iota
	TokenTypeWhiteSpace
	TokenTypeVowel
	TokenTypeConsonant
	TokenTypeNumber
	TokenTypeOther
)

type Token struct {
	Type  TokenType
	Value rune
}
