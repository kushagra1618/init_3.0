//This file contains code for http server of our website
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"init/utils"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create/app", create_app_handler)

	log.Fatal(http.ListenAndServe("localhost:5050", nil))
	fmt.Println("Listening at 5050")
}

//Handler for New App creation
func create_app_handler(resp http.ResponseWriter, req *http.Request) {
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
			var iscreated = CreateAppinDb(&info)
			res["id"] = info.Id
			res["name"] = info.Name
			res["code"] = utils.CondOp(iscreated, http.StatusCreated, http.StatusFound)

		}

	}

	resp.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(res)
	resp.Write(b)

}

//Creates App in the Database
func CreateAppinDb(info *utils.NodeInfo) bool {

	db, err := sql.Open("mysql", "root:root@/init")
	checkerr(err)
	defer db.Close()
	stmt, err := db.Prepare("SELECT ID,NAME FROM apps where EMAIL=?")
	checkerr(err)
	defer stmt.Close()
	row, err := stmt.Query(info.Email)
	checkerr(err)
	defer row.Close()
	if row.Next() {
		row.Scan(&info.Id, &info.Name)
		return false
	} else { //Creating new App id
		//pr, pu := utils.GenerateRsaKeyPair()
		stmt, err := db.Prepare("INSERT INTO apps(NAME,EMAIL,DOMAIN,CONTACT,RSAPUBKEY)VALUES(?,?,?,?,?)")
		checkerr(err)
		rs, err := stmt.Exec(info.Name, info.Email, info.Domain, info.Contact, info.Pubkey)
		checkerr(err)
		id, err := rs.LastInsertId()
		info.Id = id

	}
	return true

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}

}
