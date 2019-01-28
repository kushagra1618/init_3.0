//This file contains code for MasterNode
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42

package main

import (
	"bytes"
	"flag"
	"fmt"
	"init/utils"
	"log"
	"net"
	"os"
)

func main() {
	do := flag.String("do", "", "Action to take against current node")
	port := flag.Int("port", utils.MAINNODEPORT, "Port to start server on")

	flag.Parse()
	switch *do {
	case "start":
		startNode(*port)
	default:
		fmt.Println("Better luck next time")
		fmt.Println("Here is the options ")
		flag.PrintDefaults()
	}
}

func startNode(port int) {

	list, err := net.Listen(utils.TRANSPORTPROTOCOL, fmt.Sprintf("%s:%d", utils.MAINNODE, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, er := list.Accept()
		if er != nil {
			fmt.Println(er)
		}
		//	defer conn.Close()
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var b [128]byte

	buf := bytes.NewBuffer(b[:]) //Reading in Buffer

	for {

		_, err := buf.ReadFrom(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		buf.WriteTo(os.Stdout)
	}

}
