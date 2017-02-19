package lexer

import (
	"github.com/vishen/go-monkeylang/token"
)

type Lexer struct {
	input    string
	pos      int
	read_pos int
	ch       byte // TODO(): Needs to be a rune to be able to handle UTF-8
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peek() == '=' {
			l.advance()
			t.Type = token.EQUALS
			t.Literal = l.input[l.pos-1 : l.pos+1]
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peek() == '=' {
			l.advance()
			t.Type = token.NOT_EQUALS
			t.Literal = l.input[l.pos-1 : l.pos+1]
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.advance()
	return t
}

func (l *Lexer) advance() {
	if l.read_pos >= len(l.input) {
		l.ch = 0 // Ascii code for NUL
	} else {
		l.ch = l.input[l.read_pos]
	}
	l.pos = l.read_pos
	l.read_pos += 1
}

func (l *Lexer) peek() byte {
	if l.read_pos >= len(l.input) {
		return 0 // Ascii code for NUL
	} else {
		return l.input[l.read_pos]
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos

	for isLetter(l.ch) {
		l.advance()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos

	for isDigit(l.ch) {
		l.advance()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.advance()
	}
}

// Utils
func newToken(token_type token.TokenType, ch byte) token.Token {
	return token.Token{Type: token_type, Literal: string(ch)}
}

// TODO(): Change function name to something more meaningful
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
