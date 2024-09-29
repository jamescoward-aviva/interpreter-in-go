package flexer

import (
	"iter"
	"monkey/token"
)

const (
	EofByte = 0
)

// pure function implementation of Lexer.readChar()
func readChar(position int, input string) (string, byte) {
	if position >= len(input) {
		return "", EofByte
	} else {
		char := input[position]
		position++
		return input[position:], char
	}
}

func peakChar(postion int, input string) byte {
	if postion >= len(input) {
		return EofByte
	} else {
		return input[postion]
	}
}

func readIdentifier(position int, input string) (string, string) {
	return readToken(position, input, isLetter)
}

func readNumber(position int, input string) (string, string) {
	return readToken(position, input, isDigit)
}

// generic implementation of readNumber or readIdentifier
func readToken(position int, input string, isType func(byte) bool) (string, string) {
	readPosition := position
	for isType(input[readPosition]) {
		readPosition++
	}
	return input[readPosition:], input[:readPosition]
}

func skipWhitespace(position int, input string) string {
	if position >= len(input) {
		return ""
	}

	ch := peakChar(position, input)
	for isWhitespace(ch) {
		position++
		ch = peakChar(position, input)
	}
	return input[position:]
}

func nextToken(input string) (string, token.Token) {
	position := 0
	input = skipWhitespace(position, input)

	char := peakChar(position, input)

	if char == EofByte { // EOF
		return "", newEofToken()
	}

	if ok, input, tok := getDoubleToken(char, position, input); ok {
		return input, tok
	}

	if ok, input, tok := getUnitToken(input); ok {
		return input, tok
	}

	if isLetter(char) {
		input, literal := readIdentifier(position, input)
		return input, newLiteralToken(token.LookupIdent(literal), literal)
	}

	if isDigit(char) {
		input, literal := readNumber(position, input)
		return input, newLiteralToken(token.INT, literal)
	}

	return input, newToken(token.ILLEGAL, char)

}

func Flex(input string) iter.Seq2[string, token.Token] {
	return func(yield func(string, token.Token) bool) {
		for {
			var tok token.Token
			input, tok = nextToken(input)

			if !yield(input, tok) || tok.Type == token.EOF || tok.Type == token.ILLEGAL {
				break
			}
		}
	}
}

func readCharAndGetToken(position int, input string, tt token.TokenType) (string, token.Token) {
	input, char := readChar(position, input)
	return input, newToken(tt, char)
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

func getDoubleToken(ch byte, position int, input string) (bool, string, token.Token) {
	readPosition := position + 1
	double := string(ch) + string(peakChar(readPosition, input))
	found := true
	var tok token.Token

	if double == "==" {
		input, _ = readChar(readPosition, input)
		tok = newStringToken(token.EQ, double)
		return found, input, tok
	} else if double == "!=" {
		input, _ = readChar(readPosition, input)
		tok = newStringToken(token.NOT_EQ, double)
		return found, input, tok
	} else {
		found = false
		return found, input, tok
	}

}

func getUnitToken(input string) (bool, string, token.Token) {
	position := 0
	newInput, char := readChar(position, input)

	if unitToken, ok := unitTokens[char]; ok {
		return true, newInput, unitToken
	} else {
		return false, input, token.Token{}
	}
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
