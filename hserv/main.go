//This file contains code for http server of our website
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package main

import (
	"encoding/json"
	"fmt"
	"init/hserv/pages"
	"init/utils"
	"log"
	"net/http"
	"strings"

	"github.com/mr-tron/base58"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create/app", createapphandler)
	http.HandleFunc("/create/user", createuserhandler)
	log.Fatal(http.ListenAndServe("localhost:5050", nil))
	fmt.Println("Listening at 5050")
}

//Handler for New App creation
func createapphandler(resp http.ResponseWriter, req *http.Request) {
	var res = make(map[string]interface{})
	if strings.ToLower(req.Method) == "post" {
		req.ParseForm() //Parsing Form
		var form = req.Form
		if form.Get("mode") == "create" {
			info := utils.NodeInfo{}
			info.Name = form.Get("name")
			info.Email = form.Get("email")
			info.Domain = form.Get("domain")
			info.Contact = form.Get("contact")
			info.Pubkey = form.Get("pubkey")
			var iscreated = pages.CreateAppinDb(&info)
			res["id"] = info.Id
			res["name"] = info.Name
			res["code"] = utils.CondOp(iscreated, http.StatusCreated, http.StatusFound)

		}

	}

	resp.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(res)
	resp.Write(b)

}

//This function handles user creatation
func createuserhandler(resp http.ResponseWriter, req *http.Request) {
	if strings.ToLower(req.Method) == "post" {
		req.ParseForm()
		var fr = req.Form
		var user = make(utils.UserInfo)
		user[utils.GetUserInfoKey("fname")] = fr.Get("fname")
		user[utils.GetUserInfoKey("mname")] = fr.Get("mname")
		user[utils.GetUserInfoKey("lname")] = fr.Get("lname")
		user[utils.GetUserInfoKey("dob")] = fr.Get("dob")
		user[utils.GetUserInfoKey("gender")] = fr.Get("gender")
		user[utils.GetUserInfoKey("username")] = fr.Get("username")
		user[utils.GetUserInfoKey("password")] = fr.Get("password")
		user, iscreated := pages.CreateUser(user) //Creation of user on ipfs network
		resp.WriteHeader(http.StatusOK)

		if iscreated {

			passcode := base58.Encode(user[utils.GetUserInfoKey("passcode")].([]byte))
			pubaddr := base58.Encode(user[utils.GetUserInfoKey("pubaddr")].([]byte))
			var v = make(map[string]interface{})
			v["code"] = passcode
			v["addr"] = pubaddr
			b, _ := json.Marshal(v)
			resp.Header().Set("content-type", "application/json")
			resp.Write(b)
		} else {
			resp.Write([]byte("Couldn't make it"))
		}
	}
}
