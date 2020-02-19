package main

import (
	"encoding/json"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/defs"
	"io"
	"net/http"
)



func sendErrorResponse(writer http.ResponseWriter, errResp defs.ErrResponse) {
	writer.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(writer, string(resStr))
}

func sendNormalResponse(writer http.ResponseWriter, resp string, sc int){
	writer.WriteHeader(sc)
	io.WriteString(writer, resp)
}