package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listening for Connections ..........")
	for {
		conn, er := l.Accept()
		if er != nil {
			fmt.Println(er)
			return
		}
		defer conn.Close()
		go handelIO(conn)
	}
}

func handelIO(conn net.Conn) {
	var buf [128]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n, "Bytes Recieved from ", conn.LocalAddr().Network(), "\t", string(buf[:]))

	}
}

//1101491668
//11
