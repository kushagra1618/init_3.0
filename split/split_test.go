package split

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	s := Secret{}
	f := "e424cc3ef5c62accff7ebe9c3d797927"
	p := "d41ca9b3ff93b24da439c32ab28c24fd"
	s.SetPassword(p)
	s.SetCid(f)
	addr, code, err := s.CalculateAddrAndCode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(addr, code)
}
