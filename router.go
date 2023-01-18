package grrt

import "net/http"

// Router Router
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route
	NewRoute() Route
}

// NewRouter NewRouter
func NewRouter() Router {
	var rtn = ReqRouter{namedRoutes: make(map[string]Route)}
	return &rtn
}

// go mod init github.com/Ulbora/grr
