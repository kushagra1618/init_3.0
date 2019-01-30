//This file contains some helper function to simplify some frequent tasks
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42

package utils

import "crypto/sha256"

//This function calculates sha256 hash of given string
func Sha256Hash(msg []byte) []byte {
	s := sha256.New()
	s.Write(msg)
	return s.Sum(nil)
}

//CondOp, This implements  C like Conditional Operator (bool_expresion)?left:right
func CondOp(c bool, left, right interface{}) interface{} {
	if c {
		return left
	}
	return right
}
