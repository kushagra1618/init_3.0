//@author Devansh Gupta
//facebook.com/devansh42
//github.com/devansh42
//This file contains code to generate fbs of uit  tree
package pages

import (
	"init/utils"

	"init/uit/fbs/fb"
	_ "init/uit/fbs/fb"

	gfb "github.com/google/flatbuffers/go"
)

//createUserFB, This function creates Flatbuffer of user info tree
//returns byte slice containing data written to fb
func createUserFB(user utils.UserInfo) []byte {
	buf := gfb.NewBuilder(1024)  //Initial Buffer size, Resizeable
	var fields = []gfb.UOffsetT{ //Fields for fbs
		buf.CreateString(user[utils.GetUserInfoKey("fname")].(string)),
		buf.CreateString(user[utils.GetUserInfoKey("lname")].(string)),
		buf.CreateString(user[utils.GetUserInfoKey("mname")].(string)),
		buf.CreateString(user[utils.GetUserInfoKey("username")].(string))}

	fb.UserInfoStart(buf)
	fb.UserInfoAddFname(buf, fields[0])
	fb.UserInfoAddLname(buf, fields[1])
	fb.UserInfoAddMname(buf, fields[2])
	fb.UserInfoAddGender(buf, 1) //Male
	fb.UserInfoAddEmail(buf, fields[3])
	fb.UserInfoAddCountry(buf, 1) //Bharat
	t := fb.UserInfoEnd(buf)
	buf.Finish(t) //Finished

	return buf.FinishedBytes()

}
