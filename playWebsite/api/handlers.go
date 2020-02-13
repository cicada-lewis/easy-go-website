package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("Create User Handler"))
	io.WriteString(writer, "Create User Handler")
	return
}

func LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	io.WriteString(writer, fmt.Sprintf("user: %s has login.", params.ByName("user_name")))
}
