package flexer

import (
	"iter"
	"monkey/token"
)

const (
	EofByte = 0
)

type Input string

// pure function implementation of Lexer.readChar()
func readChar(input Input) (byte, Input) {
	if len(input) == 0 {
		return EofByte, ""
	}

	return input[0], input[1:]
}

func peakChar(input Input) byte {
	if len(input) == 0 {
		return EofByte
	} else {
		return input[0]
	}
}

func readIdentifier(input Input) (string, Input) {
	return readToken(input, isLetter)
}

func readNumber(input Input) (string, Input) {
	return readToken(input, isDigit)
}

// generic implementation of readNumber or readIdentifier
func readToken(input Input, isType func(byte) bool) (string, Input) {
	readPosition := 0
	for isType(input[readPosition]) {
		readPosition++
	}
	return string(input[:readPosition]), input[readPosition:]
}

func skipWhitespace(input Input) Input {
	if len(input) == 0 {
		return ""
	}

	ch := peakChar(input)

	for isWhitespace(ch) {
		// Consume the peaked character
		_, input = readChar(input)
		ch = peakChar(input)
	}

	return input
}

func nextToken(input Input) (token.Token, Input) {
	input = skipWhitespace(input)

	ch := peakChar(input)

	if ch == EofByte { // EOF
		return newEofToken(), ""
	}

	if ok, input, tok := getDoubleToken(input); ok {
		return tok, input
	}

	if ok, input, tok := getUnitToken(input); ok {
		return tok, input
	}

	if isLetter(ch) {
		literal, input := readIdentifier(input)
		return newLiteralToken(token.LookupIdent(literal), literal), input
	}

	if isDigit(ch) {
		literal, input := readNumber(input)
		return newLiteralToken(token.INT, literal), input
	}

	return newToken(token.ILLEGAL, ch), input

}

func Flex(input Input) iter.Seq2[token.Token, Input] {
	return func(yield func(token.Token, Input) bool) {
		for {
			var tok token.Token
			tok, input = nextToken(input)

			if !yield(tok, input) || tok.Type == token.EOF || tok.Type == token.ILLEGAL {
				break
			}
		}
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newStringToken(tokenType token.TokenType, s string) token.Token {
	return token.Token{Type: tokenType, Literal: s}
}

func newEofToken() token.Token {
	return token.Token{Type: token.EOF, Literal: ""}
}

func getDoubleToken(input Input) (bool, Input, token.Token) {
	var ch1, ch2 byte
	var readInput Input
	ch1, readInput = readChar(input)
	ch2, readInput = readChar(readInput)
	double := string(ch1) + string(ch2)
	found := true
	var tok token.Token

	if double == "==" {
		tok = newStringToken(token.EQ, double)
		return found, readInput, tok
	}

	if double == "!=" {
		tok = newStringToken(token.NOT_EQ, double)
		return found, readInput, tok
	}

	// If no double token is found, return the input as is
	found = false
	return found, input, tok

}

func getUnitToken(input Input) (bool, Input, token.Token) {
	ch, readInput := readChar(input)
	if unitToken, ok := unitTokens[ch]; ok {
		return true, readInput, unitToken
	}

	// If no unit token is found, return the input as is
	return false, input, token.Token{}
}

func newLiteralToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
