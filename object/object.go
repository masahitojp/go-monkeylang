package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER      = "INTEGER"
	BOOLEAN      = "BOOLEAN"
	RETURN_VALUE = "RETURN_VALUE"
	ERROR        = "ERROR"
	NULL         = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
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
// TODO(): There must be a better way for storing these, specially as they will be falling
// in and out of scope all the time
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
