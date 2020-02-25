package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)


type middlewareHandler struct {
	router *httprouter.Router
	limiter *connLimiter
}

func NewMiddlewareHandler(router *httprouter.Router, cc int) http.Handler {
	mh := middlewareHandler{}
	mh.router = router
	mh.limiter = NewConnLimiter(cc)
	return mh
}

func (mh middlewareHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if !mh.limiter.GetConn() {
		sendErrorResponse(writer, http.StatusTooManyRequests, "Too Many Requests!")
		return
	}
	mh.router.ServeHTTP(writer, req)
	defer mh.limiter.ReleaseConn()
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	return router
}


func main() {
	r := RegisterHandler()
	mh := NewMiddlewareHandler(r, 2)
	log.Printf("Listening at Address: 127.0.0.1:9000")
	http.ListenAndServe("127.0.0.1:9000", mh)
}