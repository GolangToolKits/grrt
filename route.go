package grrt

import "net/http"

// Route Route
type Route interface {
	Handler(handler http.Handler) Route
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route
	Path(p string) Route
	Host(h string) Route
	
}
