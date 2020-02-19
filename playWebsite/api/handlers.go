package main

import (
	"encoding/json"
	"fmt"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/dbops"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/defs"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/session"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)


func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	 res, _ := ioutil.ReadAll(request.Body)
	 ubody := &defs.UserCredential{}

	 if err := json.Unmarshal(res, ubody); err != nil {
	 	sendErrorResponse(writer, defs.ErrorResponseBodyParseFailed)
	 	return
	 }
	 if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
	 	sendErrorResponse(writer, defs.ErrorDBError)
	 }
	 sid := session.GenerateNewSessionId(ubody.Username)
	 su := &defs.SignedUp{Success: true, SessionId: sid}
	 if resp, err := json.Marshal(su); err != nil {
		return
	} else {
		sendNormalResponse(writer, string(resp), 200)
	 }
}

func LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	io.WriteString(writer, fmt.Sprintf("user: %s has login.", params.ByName("user_name")))
}
