//This file contains code for New app registration
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package pages

import (
	"database/sql"
	"encoding/hex"
	"init/utils"

	"github.com/mr-tron/base58"

	_ "github.com/go-sql-driver/mysql" //Mysql Driver for golang
)

//Creates App in the Database,
//Returns true if new app was created
func CreateAppinDb(info *utils.NodeInfo) bool {

	db, err := sql.Open(utils.DBDRIVER, utils.DBDATASOURCE) //Sql connection on localhost username=root and password=root
	checkerr(err)
	defer db.Close()
	stmt, err := db.Prepare("SELECT ID,NAME FROM apps where EMAIL=?") //Query to check If App already exists
	checkerr(err)
	defer stmt.Close()
	row, err := stmt.Query(info.Email)
	checkerr(err)
	defer row.Close()
	if row.Next() { //Checking for results
		row.Scan(&info.Id, &info.Name)
		return false
	} else { //Creating new App id
		bpubkey, _ := hex.DecodeString(info.Pubkey)
		stmt, err := db.Prepare("INSERT INTO apps(NAME,EMAIL,DOMAIN,CONTACT,RSAPUBKEY)VALUES(?,?,?,?,?)")
		checkerr(err)
		rs, err := stmt.Exec(info.Name, info.Email, info.Domain, info.Contact, base58.Encode(bpubkey))
		checkerr(err)
		id, err := rs.LastInsertId()
		info.Id = id

	}
	return true

}

//For Error Check
func checkerr(err error) {
	if err != nil {
		panic(err)
	}

}
