//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
//This file contains code for splitting up Cid

package split

import "errors"

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
