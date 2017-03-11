package ast

import (
	"fmt"

	"github.com/vishen/go-monkeylang/token"
)

type Node interface {
	TokenLiteral() string
	Useful() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Root node of the tree
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Let statement
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls LetStatement) statementNode()       {}
func (ls LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls LetStatement) Useful() string {
	return fmt.Sprintf("ast.LetStatement -> Token=%s, Name=%s", ls.Token.Useful(), ls.Name.Useful())
}

// Identifier statement
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i Identifier) expressionNode()      {}
func (i Identifier) TokenLiteral() string { return i.Token.Literal }
func (i Identifier) Useful() string {
	return fmt.Sprintf("ast.Identifier -> Token=%s, Value=%s", i.Token.Useful(), i.Value)
}

// Return statement
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs ReturnStatement) statementNode()       {}
func (rs ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs ReturnStatement) Useful() string {
	return fmt.Sprintf("ast.ReturnStatement -> Token=%s, ReturnValue=%s", rs.Token.Useful(), "NI")
}

// Expression statement
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es ExpressionStatement) statementNode()       {}
func (es ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es ExpressionStatement) Useful() string {
	return fmt.Sprintf("ast.ExpressionStatement -> Token=%s, Expression=%s", es.Token.Useful(), "NI")
}

// Integer Literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il IntegerLiteral) expressionNode()      {}
func (il IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il IntegerLiteral) Useful() string {
	return fmt.Sprintf("ast.IntegerLiteral -> Token=%s Value=%d", il.Token.Useful(), il.Value)
}

// Prefix Expression
type PrefixExpression struct {
	Token    token.Token // Prefix token; !, -
	Operator string
	Right    Expression
}

func (pe PrefixExpression) expressionNode()      {}
func (pe PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe PrefixExpression) Useful() string {
	return fmt.Sprintf("ast.PrefixExpression -> Token=%s, Operator=%s, Right=%s",
		pe.Token.Useful(), pe.Operator, "NI")
}
