package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/vishen/go-monkeylang/ast"
)

type ObjectType string

const (
	INTEGER      = "INTEGER"
	BOOLEAN      = "BOOLEAN"
	RETURN_VALUE = "RETURN_VALUE"
	FUNCTION     = "FUNCTION"
	ERROR        = "ERROR"
	NULL         = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f Function) Type() ObjectType { return FUNCTION }
func (f Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type Integer struct {
	Value int64
}

func (i Integer) Type() ObjectType { return INTEGER }
func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b Boolean) Type() ObjectType { return BOOLEAN }
func (b Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n Null) Type() ObjectType { return NULL }
func (n Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv ReturnValue) Type() ObjectType { return RETURN_VALUE }
func (rv ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Environment for storing variables...
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
