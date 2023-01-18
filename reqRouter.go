package grrt

import "net/http"

// ReqRouter RequestRouter
type ReqRouter struct {
	namedRoutes map[string]Route
}

func (t ReqRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// NewRoute NewRoute
func (t ReqRouter) NewRoute() Route {
	return nil
}

// HandleFunc HandleFunc
func (t ReqRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route {
	return nil
}
