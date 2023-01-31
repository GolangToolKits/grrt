package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import (
	"context"
	"net/http"
)

// Router Router
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Handle(path string, handler http.Handler) Route
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route
	PathPrefix(path string) Route
	NewRoute() Route

	//CORS methods
	EnableCORS()
	CORSAllowCredentials()
	SetCorsAllowedHeaders(headers string)
	SetCorsAllowedOrigins(origins string)
	SetCorsAllowedMethods(methods string)
}

// NewRouter NewRouter creates new Router
func NewRouter() Router {
	var rtn = ReqRouter{namedRoutes: make(map[string]*[]Route)}
	return &rtn
}

// Vars Vars returns the path variables for the current request
func Vars(r *http.Request) map[string]string {
	var rtn map[string]string
	if rv := r.Context().Value(varsKey); rv != nil {
		rtn = rv.(map[string]string)
	}
	return rtn
}

// SetURLVars SetURLVars
func SetURLVars(r *http.Request, vars map[string]string) *http.Request {
	ctxi := context.WithValue(r.Context(), varsKey, vars)
	return r.WithContext(ctxi)
}

// go mod init github.com/GolangToolKits/grrt
