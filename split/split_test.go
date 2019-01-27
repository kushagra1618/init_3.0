package split

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	s := Secret{}
	f := "e424cc3ef5c62accff7ebe9c3d797927af59976677501b2c4cd9a2f046218952"
	p := "d41ca9b3ff93b24da439c32ab28c24fd03220fbee13d3c4650f20125172ae72d"
	bf, _ := hex.DecodeString(f)
	bp, _ := hex.DecodeString(p)
	fmt.Println(bf, bp)

	s.SetCid(bf)
	s.SetPassword(bp)
	addr, code, _ := s.CalculateAddrAndCode()
	s.SetPasscode(code)
	s.SetAddrs(addr)
	cid, _ := s.RecalculateCid()
	t.Log(cid)
}
