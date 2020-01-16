package main

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

const ()
