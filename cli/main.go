//This file contains code of friend node executable file
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package main

import (
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"init/utils"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	startLogging()
	defer logfile.Close()

	parseArgs()

}

var Log *log.Logger
var logfile *os.File

func parseArgs() {
	do := flag.String("do", "", "Process command against current node")
	port := flag.Int("port", utils.FRIENDNODEPORT, fmt.Sprint("Start node at given port default is ", utils.FRIENDNODEPORT))
	flag.Parse()
	switch *do {
	case "start":
		startNode(*port)
	case "reg":
		registerNode()
	default:
		fmt.Println("Okay, See u soon")
	}
}

//This function handles Initialization of node
func startNode(port int) {
	listener, err := net.Listen(utils.TRANSPORTPROTOCOL, fmt.Sprintf("localhost:%d", port))

	if err != nil {
		Log.Println("There is a Problem,", err)
		return
	}
	defer listener.Close()

	for { //Listen's Requests forever
		conn, er := listener.Accept()
		if er != nil {
			Log.Fatal(er)

		}
		defer conn.Close()
		go io.Copy(os.Stdout, conn)
	}

}

//This function handles logging for Node
func startLogging() {
	logfile, err := os.Create("logs.log")
	if err != nil {
		fmt.Println("Couldn't open Log File")
		return
	}

	Log = log.New(logfile, "Info: ", 1)
	io.Copy(os.Stdout, logfile)

}

//Function to handle friend node registration
func registerNode() {
	info := utils.NodeInfo{}
	fmt.Print("\nEnter Application name\t")
	fmt.Scanf("%s", &info.Name)
	fmt.Print("\nEnter Email\t")
	fmt.Scanf("%s", &info.Email)
	fmt.Print("\nEnter Domain\t")
	fmt.Scanf("%s", &info.Domain)
	fmt.Print("\nEnter Contact\t")
	fmt.Scanf("%s", &info.Contact)
	pr, pu := utils.GenerateRsaKeyPair() //Generating new PKI pair using rsa
	pub, _ := pem.Decode(pu)
	pubks := hex.EncodeToString(pub.Bytes)
	info.Pubkey = pubks

	var m = make(url.Values)
	m.Set("pubkey", info.Pubkey)
	m.Set("name", info.Name)
	m.Set("email", info.Email)
	m.Set("domain", info.Domain)
	m.Set("contact", info.Contact)
	m.Set("mode", "create")
	fmt.Println("Sending request to Web Http Server..")
	resp, err := http.PostForm(fmt.Sprintf("%s/create/app", utils.MAINNODEWEBADDR), m)
	if err != nil {
		fmt.Println(err)
	}
	rbody := resp.Body
	defer rbody.Close()
	by, _ := ioutil.ReadAll(rbody)
	var v = make(map[string]interface{})
	json.Unmarshal(by, &v)
	switch int(v["code"].(float64)) {
	case http.StatusCreated:
		utils.WritePrivateKeyTOPemFile(pr, pu, "keys") //Writing new Private key to local file system
		fmt.Println("Congratulations! You have registered your node ")
		fmt.Println("Your App ", v["name"], "'s App Id is ", v["id"])
		fmt.Println("All info is avilable in app.json")
		fmt.Println("Your RSA Public key  is \n")
		pem.Encode(os.Stdout, pub)
		info.Name = string(v["name"].(string))
		info.Id = int64(v["id"].(float64))
		infobb, _ := json.Marshal(info)
		appf, _ := os.Create("app.json")
		defer appf.Close()
		appf.Write(infobb)

	case http.StatusFound:
		fmt.Println("Your have already registered your node ")
		fmt.Println("Your App ", v["name"], "'s App Id is ", v["id"])

		fmt.Println("All information should be in app.json")

	}

}
