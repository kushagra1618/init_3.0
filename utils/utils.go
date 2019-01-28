//This file contains some helper function to simplify some frequent tasks
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42

package utils

//CondOp, This implements  C like Conditional Operator (bool_expresion)?left:right
func CondOp(c bool, left, right interface{}) interface{} {
	if c {
		return left
	}
	return right
}
