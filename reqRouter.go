package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import (
	"context"
	"log"
	"net/http"
)

// ReqRouter RequestRouter
type ReqRouter struct {
	namedRoutes map[string]Route
}

func (t ReqRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.getReqVars(r)
	path := r.URL.Path
	rt := t.namedRoutes[path]
	if rt == nil || !rt.IsActive() {
		w.WriteHeader(http.StatusNotFound)
	} else if !rt.IsMethodAllowed(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		hd := rt.GetHandler()
		hd.ServeHTTP(w, r)
	}
}

// NewRoute NewRoute
func (t ReqRouter) NewRoute() Route {
	var rt ReqRoute
	rrt := rt.New()
	return rrt
}

// HandleFunc HandleFunc
func (t ReqRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route {
	rt := t.NewRoute().Path(path).HandlerFunc(f)
	t.namedRoutes[rt.GetPath()] = rt
	return rt
}

func (t ReqRouter) getReqVars(r *http.Request) {
	var vars map[string]string
	ctx := context.WithValue(r.Context(), 0, vars)
	log.Println("ctx: ", ctx)
}
