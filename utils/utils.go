package utils

//CondOp it implements the C like Conditional Operator
func CondOp(c bool, left, right interface{}) interface{} {
	if c {
		return left
	}
	return right
}
