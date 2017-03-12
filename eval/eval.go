package eval

import (
	//	"fmt"

	"github.com/vishen/go-monkeylang/ast"
	"github.com/vishen/go-monkeylang/object"
)

// Commonly used objects
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	//	fmt.Printf("Node=%#v\n", node)
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		if node.Value {
			return TRUE
		} else {
			return FALSE
		}
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range statements {
		result = Eval(stmt)
		//		fmt.Printf("i=%d stmt=%#v result=%#v", i, stmt, result)
	}

	return result
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {

	if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.Integer).Value

		switch operator {
		// Return Integers
		case "+":
			return &object.Integer{Value: leftVal + rightVal}
		case "-":
			return &object.Integer{Value: leftVal - rightVal}
		case "*":
			return &object.Integer{Value: leftVal * rightVal}
		case "/":
			return &object.Integer{Value: leftVal / rightVal}
			// Return Boolean
		case "<":
			return nativeBoolToBooleanObject(leftVal < rightVal)
		case ">":
			return nativeBoolToBooleanObject(leftVal > rightVal)
		case "==":
			return nativeBoolToBooleanObject(leftVal == rightVal)
		case "!=":
			return nativeBoolToBooleanObject(leftVal != rightVal)
		default:
			return NULL
		}
	} else if operator == "==" {
		return nativeBoolToBooleanObject(left == right)
	} else if operator == "!=" {
		return nativeBoolToBooleanObject(left != right)
	} else {
		return NULL
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		// TODO(): Should this return an error?
		return NULL
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	// TODO(): Should this return an error?
	if right.Type() != object.INTEGER {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func nativeBoolToBooleanObject(truthy bool) object.Object {
	if truthy {
		return TRUE
	} else {
		return FALSE
	}
}
