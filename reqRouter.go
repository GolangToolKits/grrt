package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import (
	"context"
	"log"
	"net/http"
	"strings"
)

type contextKey int

const (
	varsKey contextKey = iota
)

// ReqRouter RequestRouter
type ReqRouter struct {
	namedRoutes map[string]*[]Route
}

// ServeHTTP ServeHTTP dispatches the handler registered in the matched route.
func (t ReqRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// When there is a match, the route variables can be retrieved calling
	// mux.Vars(request).
	path := r.URL.Path
	rt, fvars := t.findRouteAndVars(path)
	if len(*fvars) > 0 {
		r = t.requestWithVars(r, rt.GetVarNames(), fvars)
	}
	if rt == nil || !rt.IsActive() {
		w.WriteHeader(http.StatusNotFound)
	} else if !rt.IsMethodAllowed(r.Method) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		hd := rt.GetHandler()
		hd.ServeHTTP(w, r)
	}
}

func (t ReqRouter) requestWithVars(r *http.Request, pVarNames, pvars *[]string) *http.Request {
	var vars = make(map[string]string)
	if len(*pVarNames) == len(*pvars) {
		for i, n := range *pVarNames {
			vars[n] = (*pvars)[i]
		}
	}
	ctx := context.WithValue(r.Context(), varsKey, vars)
	return r.WithContext(ctx)
}

func (t ReqRouter) findRouteAndVars(path string) (Route, *[]string) {
	var rnt Route
	sp := strings.Split(path, "/")
	var vars []string
	var vcnt = len(sp) - 2
	log.Println("sp:", sp)
	var found = false
	var searchPath = ""
	for i, p := range sp {
		if i == 0 {
			continue
		} else if found {
			break
		}
		searchPath += "/" + p
		rts := t.namedRoutes[searchPath]
		if rts != nil {
			for _, rt := range *rts {
				if rt.GetPathVarsCount() == vcnt {
					rnt = rt
					found = true
					vars = sp[i+1:]
					break
				}
			}
		} else {
			vcnt--
		}
	}
	return rnt, &vars
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
	fts := t.namedRoutes[rt.GetPath()]
	if fts == nil {
		t.namedRoutes[rt.GetPath()] = &[]Route{rt}
	} else {
		var addRt = true
		for _, rtf := range *fts {
			if rt.GetPathVarsCount() == rtf.GetPathVarsCount() {
				addRt = false
			}
		}
		if addRt {
			*fts = append(*fts, rt)
		}

	}
	return rt
}
