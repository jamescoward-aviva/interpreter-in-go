package flexer

import "monkey/token"

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

func NextToken(input string) (string, token.Token) {
	var tok token.Token
	position := 0
	input = skipWhitespace(position, input)

	char := peakChar(position, input)
	switch char {
	case '=':
		readPosition := position + 1
		if peakChar(readPosition, input) == '=' {
			input, _ = readChar(readPosition, input)
			tok = newStringToken(token.EQ, "==")
		} else {
			input, tok = readCharAndGetToken(position, input, token.ASSIGN)
		}
	case '+':
		input, tok = readCharAndGetToken(position, input, token.PLUS)
	case '-':
		input, tok = readCharAndGetToken(position, input, token.MINUS)
	case '!':
		readPosition := position + 1
		if peakChar(readPosition, input) == '=' {
			input, _ = readChar(readPosition, input)
			tok = newStringToken(token.NOT_EQ, "!=")
		} else {
			input, tok = readCharAndGetToken(position, input, token.BANG)
		}
	case '/':
		input, tok = readCharAndGetToken(position, input, token.SLASH)
	case '*':
		input, tok = readCharAndGetToken(position, input, token.ASTERISK)
	case '<':
		input, tok = readCharAndGetToken(position, input, token.LT)
	case '>':
		input, tok = readCharAndGetToken(position, input, token.GT)
	case ';':
		input, tok = readCharAndGetToken(position, input, token.SEMICOLON)
	case '(':
		input, tok = readCharAndGetToken(position, input, token.LPAREN)
	case ')':
		input, tok = readCharAndGetToken(position, input, token.RPAREN)
	case '{':
		input, tok = readCharAndGetToken(position, input, token.LBRACE)
	case '}':
		input, tok = readCharAndGetToken(position, input, token.RBRACE)
	case ',':
		input, tok = readCharAndGetToken(position, input, token.COMMA)
	case EofByte:
		tok = newStringToken(token.EOF, "")
	default:
		var literal string
		if isLetter(char) {
			input, literal = readIdentifier(position, input)
			tok = newLiteralToken(token.LookupIdent(literal), literal)
		} else if isDigit(char) {
			input, literal = readNumber(position, input)
			tok = newLiteralToken(token.INT, literal)
		} else {
			input, tok = readCharAndGetToken(position, input, token.ILLEGAL)
		}
	}

	return input, tok

}

func Flex(input string) []token.Token {
	var tokens []token.Token
	for input, tok := NextToken(input); tok.Type != token.EOF; input, tok = NextToken(input) {
		tokens = append(tokens, tok)
	}
	return tokens
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
