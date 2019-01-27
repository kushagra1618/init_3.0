package utils

//CondOp, This implements  C like Conditional Operator (bool_expresion)?left:right
func CondOp(c bool, left, right interface{}) interface{} {
	if c {
		return left
	}
	return right
}
