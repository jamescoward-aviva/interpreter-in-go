package flexer

import "monkey/token"

// unitTokens are tokens that are exactly one character long
// and whose lexical representation is the same as the string on the token type
var unitTokens = map[byte]token.Token{
	'=': {
		Type:    token.ASSIGN,
		Literal: "=",
	},
	'+': {
		Type:    token.PLUS,
		Literal: "+",
	},
	'-': {
		Type:    token.MINUS,
		Literal: "-",
	},
	'!': {
		Type:    token.BANG,
		Literal: "!",
	},
	'*': {
		Type:    token.ASTERISK,
		Literal: "*",
	},
	'/': {
		Type:    token.SLASH,
		Literal: "/",
	},
	',': {
		Type:    token.COMMA,
		Literal: ",",
	},
	';': {
		Type:    token.SEMICOLON,
		Literal: ";",
	},
	'(': {
		Type:    token.LPAREN,
		Literal: "(",
	},
	')': {
		Type:    token.RPAREN,
		Literal: ")",
	},
	'{': {
		Type:    token.LBRACE,
		Literal: "{",
	},
	'}': {
		Type:    token.RBRACE,
		Literal: "}",
	},
	'<': {
		Type:    token.LT,
		Literal: "<",
	},
	'>': {
		Type:    token.GT,
		Literal: ">",
	},
}
