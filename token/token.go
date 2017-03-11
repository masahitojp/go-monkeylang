package token

import "fmt"

type TokenType string

// TODO(): Store the filename and line number on the token?
type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) Useful() string {
	return fmt.Sprintf("token.Token -> %s.%s", t.Type, t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// Binary Comparision
	EQUALS     = "=="
	NOT_EQUALS = "!"
)

var (
	keywords = map[string]TokenType{ // TODO(): Change variable name
		"fn":     FUNCTION,
		"let":    LET,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}
)

func LookupIdent(ident string) TokenType {
	if token_type, ok := keywords[ident]; ok {
		return token_type
	}

	return IDENT
}
