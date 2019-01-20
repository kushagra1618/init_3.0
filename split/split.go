//@author Devansh Gupta
//facebook.com/devansh42
//This is the source code of Secret Split Algorithm being used in our Project

package split

import (
	"errors"
)

const (
	PASSSIZE    = 32
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

//calculateAddrAndCode, Internal method of calculating addr,passcode
func (s Secret) calculateAddrAndCode() (baddr, bcode []byte, err error) {
	if len(s.hashedPassword) == len(s.cid) {
		bpass := (s.hashedPassword)
		bcid := (s.cid)

		var bb []byte

		for i := 0; i < IPFSCIDSIZE; i++ {
			bb = append(bb, bpass[i]^bcid[i])
		}
		for x := 0; x < 4; x++ {
			bcode = append(bcode, xor(bb[0+8*x:8+8*x]))
		}
		for y := 0; y < IPFSCIDSIZE; y++ {
			if y == 0 || y == 8 || y == 16 || y == 24 {

			} else {
				baddr = append(baddr, bb[y])
			}
		}

	} else {
		err = errors.New("Password and Cid should be of same length")
	}
	return
}

//recalculateCid, Internal method to calculate Cid
func (s Secret) recalculateCid() (bcid []byte, err error) {

	if len(s.hashedPassword) == PASSSIZE {
		bpass := (s.hashedPassword)
		bcode := (s.passcode)
		baddr := (s.addr)
		var bb, missing []byte
		for x := 0; x < 4; x++ {
			y := xor(baddr[0+7*x:7+7*x]) ^ bcode[x]
			missing = append(missing, y)

		}

		o, p := 0, 0
		for x := 0; x < IPFSCIDSIZE; x++ {
			if x == 0 || x == 8 || x == 16 || x == 24 {
				bb = append(bb, missing[o])
				o++
			} else {
				bb = append(bb, baddr[p])
				p++
			}
		}
		for x := 0; x < IPFSCIDSIZE; x++ {

			bcid = append(bcid, bpass[x]^bb[x])
		}

	} else {
		return nil, errors.New("Invalid Password")
	}
	return
}
