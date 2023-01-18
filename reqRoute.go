package grrt

import "net/http"

// ReqRoute ReqRoute
type ReqRoute struct {
	namedRoutes map[string]*Route
	handler     http.Handler
	host        string
	path        string
}

// Handler Handler
func (t *ReqRoute) Handler(handler http.Handler) Route {
	t.handler = handler
	return t
}

// HandlerFunc HandlerFunc
func (t *ReqRoute) HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route {
	t.handler = http.HandlerFunc(f)
	return t
}

// Path Path
func (t *ReqRoute) Path(p string) Route {
	var rtn Route
	if matchPath(p) {
		rtn = t
	}
	return rtn
}

// Host Host
func (t *ReqRoute) Host(h string) Route {
	return nil
}
