package lexer

import "github.com/Buitragox/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0 // NUL
	} else {
		lex.char = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.char {
	case '=':
		tok = newToken(token.ASSIGN, lex.char)
	case ';':
		tok = newToken(token.SEMICOLON, lex.char)
	case '(':
		tok = newToken(token.LPAREN, lex.char)
	case ')':
		tok = newToken(token.RPAREN, lex.char)
	case '{':
		tok = newToken(token.LBRACE, lex.char)
	case '}':
		tok = newToken(token.RBRACE, lex.char)
	case ',':
		tok = newToken(token.COMMA, lex.char)
	case '+':
		tok = newToken(token.PLUS, lex.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lex.readChar()
	return tok
}
