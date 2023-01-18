package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

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
