package flesch

func Tokenize(b byte) Token {
	var token Token
	r := rune(b)
	token.Type = RuneToTokenType(r)
	token.Value = r

	return token
}
