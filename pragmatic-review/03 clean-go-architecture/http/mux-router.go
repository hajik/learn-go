package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}

func (*muxRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP Server runnig on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
