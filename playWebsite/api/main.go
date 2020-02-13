package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	// methods

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", LoginUser)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
