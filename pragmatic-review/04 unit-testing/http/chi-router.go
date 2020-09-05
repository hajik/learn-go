package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type chaiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chaiRouter{}
}

func (*chaiRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(url, f)
}

func (*chaiRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(url, f)
}

func (*chaiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP Server runnig on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
