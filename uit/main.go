package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Info , Basic info user
type Info struct {
	Name string
	Age  uint8

	Addr string
	List []interface{}
}

func main() {
	var ar []interface{}
	//	ar = append(ar, 1)
	gob.Register(ar)
	var m = make(map[byte]interface{})
	m[0] = []string{"devansh", "kumar", "gupta"}
	m[1] = []byte{5, 7, 99}
	m[2] = 1 //Gender Male=1,Female=2
	m[4] = byte(91)
	m[3] = "devanshguptamrt@gmail.com"
	m[5] = []interface{}{int32(9412366), byte(123)}
	gob.Register(m)
	ar = append(ar, []string{"devansh", "kumar", "gupta"})
	ar = append(ar, []byte{5, 7, 99})
	ar = append(ar, 1)
	ar = append(ar, "devanshguptamrt@gmail.com")
	ar = append(ar, byte(91))
	ar = append(ar, []interface{}{int32(9412366), byte(123)})
	f, _ := os.Create("n")
	e := gob.NewEncoder(f)
	e.Encode(m)
	f.Close()
	b, _ := ioutil.ReadFile("n")
	fmt.Println(len(b))
	h, _ := json.Marshal(m)
	fmt.Println(len(h), string(h))

	//146
	//132
	//134
}
