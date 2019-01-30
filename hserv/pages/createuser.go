//This file contains code for User creation on ipfs platform
//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
package pages

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"init/split"
	"init/utils"

	"github.com/btcsuite/btcd/btcec"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mr-tron/base58"

	shell "github.com/ipfs/go-ipfs-api"
)

//CreateUser,Creates users identity on ipfs network
//Username uniqueness check is supressed
func CreateUser(user utils.UserInfo) (utils.UserInfo, bool) {

	password := user[utils.GetUserInfoKey("password")].(string)
	username := user[utils.GetUserInfoKey("username")].(string)
	delete(user, utils.GetUserInfoKey("password")) //This deletes password from user info map
	delete(user, utils.GetUserInfoKey("username")) //This deletes username from user info map
	_, pu := getPKIPair(username, password)

	bb, _ := json.Marshal(user)
	bb, _ = encryptUIT(bb, pu) //Encrypts UIT of the user
	buf := bytes.NewBuffer(bb)
	s := shell.NewShell(utils.LOCALIPFSADDR) //Calling IPFS Shell

	ucid, err := s.Add(buf) //Added Userfile into ipfs
	checkerr(err)
	bytecid, _ := base58.Decode(ucid) //Decode Cid from base58 to byte array
	//Taking 32 bytes from last of the decoded cid
	addr, code, err := getPassCodeAndAddr(bytecid[2:], getHashedPassword(password))
	checkerr(err)
	user[utils.GetUserInfoKey("passcode")] = code
	user[utils.GetUserInfoKey("pubaddr")] = addr
	createPairingOfUsernameAndAddr(username, addr)
	fmt.Println(base58.Encode(code))
	fmt.Println(ucid)
	return user, true
}

//This function returns sha256 hashed password for user
func getHashedPassword(pass string) []byte {
	s := sha256.New()
	s.Write([]byte(pass))
	return s.Sum(nil)
}

//This function returns Passcode and Public Addr for User
func getPassCodeAndAddr(cid, password []byte) ([]byte, []byte, error) {
	s := split.Secret{}
	s.SetCid(cid)
	s.SetPassword(password)

	return s.CalculateAddrAndCode()

}

//This function created pairing of username and public address of the user
//It updates existing pairing if one exists
func createPairingOfUsernameAndAddr(usernamestr string, pubaddr []byte) error {

	s := sha256.New() //Calculation of hash of username
	s.Write([]byte(usernamestr))
	username := s.Sum(nil)
	susername := base58.Encode(username)
	spubaddr := base58.Encode(pubaddr)
	db, err := sql.Open(utils.DBDRIVER, utils.DBDATASOURCE)
	checkerr(err)
	defer db.Close()
	stmt, err := db.Prepare("SELECT PUBADDR FROM users where USERNAME=? limit 1")
	checkerr(err)
	defer stmt.Close()
	row, err := stmt.Query(susername)
	checkerr(err)
	defer row.Close()
	if row.Next() {
		stmt, err = db.Prepare("UPDATE users set PUBADDR=? WHERE USERNAME=? limit 1")
		checkerr(err)

		_, err = stmt.Exec((spubaddr), (susername))

	} else {
		stmt, err = db.Prepare("INSERT INTO users(USERNAME,PUBADDR)VALUES(?,?)")
		checkerr(err)
		_, err = stmt.Exec((susername), (spubaddr))
	}
	return err
}

//Password is double hashed because it will prevent
//xor operation to nullify everything as a^a=0
//username and password is same then seed will be 0000000000000000000000000...
func getPKIPair(username, password string) (*btcec.PrivateKey, *btcec.PublicKey) {
	var busername, bpassword = []byte(username), []byte(password)
	busername = utils.Sha256Hash(busername)
	bpassword = utils.Sha256Hash(utils.Sha256Hash(bpassword))
	seed, _ := split.XorArrays(busername, bpassword)
	master, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	var x uint32 = 0
	for !master.IsPrivate() {
		master, _ := master.Child(hdkeychain.HardenedKeyStart + x)
		master.ECPrivKey() //Just to supperess an error
		x = x + 1
	}
	prk, _ := master.ECPrivKey()

	return prk, prk.PubKey()
}

//This function encrypts uit tree with user's public key
func encryptUIT(bb []byte, pu *btcec.PublicKey) ([]byte, error) {
	return btcec.Encrypt(pu, bb)
}

//This function decrypts uit tree with user's private key
func decryptUIT(bb []byte, pr *btcec.PrivateKey) ([]byte, error) {
	return btcec.Decrypt(pr, bb)
}
