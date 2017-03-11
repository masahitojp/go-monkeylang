package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/vishen/go-monkeylang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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
func (ls LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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
func (i Identifier) String() string { return i.Value }

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
func (rs ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
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
func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
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
func (il IntegerLiteral) String() string { return il.Token.Literal }

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
func (pe PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Infix Expression
type InfixExpression struct {
	Token    token.Token // Operator token; +, /
	Operator string
	Left     Expression
	Right    Expression
}

func (ie InfixExpression) expressionNode()      {}
func (ie InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie InfixExpression) Useful() string {
	return fmt.Sprintf("ast.InfixExpression -> Token=%s, Left=%s, Operator=%s, Right=%s",
		ie.Token.Useful(), "NI", ie.Operator, "NI")
}
func (ie InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean
type Boolean struct {
	Token token.Token
	Value bool
}

func (b Boolean) expressionNode()      {}
func (b Boolean) TokenLiteral() string { return b.Token.Literal }
func (b Boolean) String() string       { return b.Token.Literal }

// If expression
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie IfExpression) expressionNode()      {}
func (ie IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs BlockStatement) statementNode()       {}
func (bs BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl FunctionLiteral) expressionNode()      {}
func (fl FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce CallExpression) expressionNode()      {}
func (ce CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
