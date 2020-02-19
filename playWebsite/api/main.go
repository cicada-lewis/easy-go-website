package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

type middlewareHandler struct {

	router *httprouter.Router

}

func NewMiddlewareHandler(r *httprouter.Router) http.Handler {

	//mw := middlewareHandler{
	//	router: r,
	//}
	mw := middlewareHandler{}
	mw.router =r
	return mw

}


func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)

	m.router.ServeHTTP(w, r)
}


func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	// methods

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", LoginUser)

	return router
}

func main() {
	r := RegisterHandlers()
	mwh := NewMiddlewareHandler(r)
	err := http.ListenAndServe("127.0.0.1:9527", mwh)
	if err != nil {
		os.Exit(-1)
	}
}
