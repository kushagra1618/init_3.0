//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
//This file contains code for forming Cid with passcode and user addr

package split

import "errors"

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
