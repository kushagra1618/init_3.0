//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
//This is the source code of Secret Split Algorithm being used in our Project


package split

import (
	"errors"
)

const (
	//PASSSIZE User password size 256 bits
	PASSSIZE = 32
	//IPFSCIDSIZE CID for the UIT Tree
	IPFSCIDSIZE = 32
)

//Secret, struct for Secret message spliting
type Secret struct {
	//Contains unexported members
	hashedPassword []byte //Hashed Password of User
	cid            []byte //Cid of UIT
	passcode       []byte //Passcode of the user
	addr           []byte //Address of the user
}

//SetPassword, sets Password of the User, If Password is not hashed yet, It will return an error
func (s *Secret) SetPassword(password []byte) error {
	if len(password) != PASSSIZE {
		return errors.New("Please provide a hashed Password")
	}
	s.hashedPassword = password
	return nil
}

//SetCid, sets Cid of the UIT, if incorrect CID is given it will return an error
func (s *Secret) SetCid(cid []byte) error {
	if len(cid) != IPFSCIDSIZE {
		return errors.New("Please provide proper Cid of UIT Tree")
	}
	s.cid = cid
	return nil
}

//GetPassword Returns password of the user
func (s Secret) GetPassword() []byte {
	return s.hashedPassword
}

//GetPasscode Returns passcode of the user after secret split algorithm
func (s Secret) GetPasscode() []byte {
	return s.passcode
}

//Getcid Returns cid of the content
func (s Secret) GetCid() []byte {
	return s.cid
}

//GetAddrs Returns Addrs of the user calculated
func (s Secret) GetAddrs() []byte {
	return s.addr
}

//xor, Internal method for xoring
func xor(b []byte) (c byte) {
	c = b[0]
	for x := 1; x < len(b); x++ {
		c = c ^ b[x]
	}
	return
}

//CalculateAddrAndCode, calculates User Address, Passcode
//returns an error if anythings goes wrong
//(Addr,Passcode,error)
func (s Secret) CalculateAddrAndCode() (addr, code []byte, err error) {
	return s.calculateAddrAndCode()
}

//SetPasscode, sets passcode for the user
//Returns error if occured
func (s *Secret) SetPasscode(p []byte) error {
	if len(p) < 4 {
		return errors.New("Passcode is to short")
	}
	s.passcode = p
	return nil
}

//SetAddrs, sets addrs for the user
//Returns error if occured
func (s *Secret) SetAddrs(a []byte) error {
	if len(a) < 28 {
		return errors.New("Addrs is to short")
	}
	s.addr = a
	return nil
}

//CalculateCid, this method calculates cid for user
//Requires user passcode and Password and addrs
//Returns error if anything misbehaves
//(cid,error)
func (s Secret) RecalculateCid() (cid []byte, err error) {
	return s.recalculateCid()

}

//This function xor to arrays of same length
func XorArrays(ar1, ar2 []byte) (ar []byte, err error) {
	if len(ar1) == len(ar2) {
		for i := 0; i < len(ar1); i++ {
			ar = append(ar, ar1[i]^ar2[i])
		}
	} else {
		err = errors.New("Arrays should of same length")
	}
	return
}
