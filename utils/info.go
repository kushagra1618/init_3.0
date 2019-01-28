//This file contains code for node info and user info
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package utils

import "strings"

type NodeInfo struct {
	Name    string
	Email   string
	Contact string
	Domain  string
	Id      int64
	Pubkey  string
}

//Map for userinfo tree
type UserInfo map[byte]interface{}

//GetUserInfoKey, Returns byte key in for string key of the user info
func GetUserInfoKey(key string) (re byte) {
	key = strings.TrimSpace(key)
	key = strings.ToLower(key)
	switch key {
	case "fname":
		re = 1
	case "mname":
		re = 2
	case "lname":
		re = 3
	case "dob":
		re = 4
	case "gender":
		re = 5
	case "cid":
		re = 6
	case "passcode":
		re = 7
	case "password":
		re = 8
	case "username":
		re = 9
	case "timestamp":
		re = 10
	case "pubaddr":
		re = 11
	default:
		re = 0
	}
	return
}
