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
	varsKey                contextKey = iota
	corsAllowOriginHeader  string     = "Access-Control-Allow-Origin"
	corsAllowHeadersHeader string     = "Access-Control-Allow-Headers"
	corsAllowMethodsHeader string     = "Access-Control-Allow-Methods"
)

// ReqRouter RequestRouter
type ReqRouter struct {
	namedRoutes    map[string]*[]Route
	prefixRoutes   map[string]Route
	corsEnabled    bool
	allowedHeaders []string
	allowedOrigins []string
	allowedMethods []string
}

// ServeHTTP ServeHTTP dispatches the handler registered in the matched route.
func (t ReqRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// When there is a match, the route variables can be retrieved calling
	// mux.Vars(request).

	if t.corsEnabled && r.Method == http.MethodOptions {
		t.handleCors(w)
	} else {
		path := r.URL.Path
		var rt = t.findPrefix(path)
		if rt == nil {
			frt, fvars := t.findRouteAndVars(path)
			rt = frt
			if len(*fvars) > 0 {
				r = t.requestWithVars(r, rt.GetVarNames(), fvars)
				// rt = frt
			}
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

	// path := r.URL.Path
	// var rt = t.findPrefix(path)
	// if rt == nil {
	// 	frt, fvars := t.findRouteAndVars(path)
	// 	rt = frt
	// 	if len(*fvars) > 0 {
	// 		r = t.requestWithVars(r, rt.GetVarNames(), fvars)
	// 		// rt = frt
	// 	}
	// }
	// if rt == nil || !rt.IsActive() {
	// 	w.WriteHeader(http.StatusNotFound)
	// } else if !rt.IsMethodAllowed(r.Method) {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// } else {
	// 	hd := rt.GetHandler()
	// 	hd.ServeHTTP(w, r)
	// }
}

// NewRoute NewRoute
func (t ReqRouter) NewRoute() Route {
	var rt ReqRoute
	rrt := rt.New()
	return rrt
}

// Handle Handle
func (t ReqRouter) Handle(path string, handler http.Handler) Route {
	rt := t.NewRoute().Path(path).Handler(handler)
	fts := t.namedRoutes[rt.GetPath()]
	if fts == nil {
		t.namedRoutes[rt.GetPath()] = &[]Route{rt}
	} else {
		var addRt = true
		for _, rtf := range *fts {
			if rt.GetPathVarsCount() == rtf.GetPathVarsCount() {
				addRt = false
				log.Println("Path not added to route, it already exists:", path)
			}
		}
		if addRt {
			*fts = append(*fts, rt)
		}
	}
	return rt
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
				log.Println("Path not added to route, it already exists:", path)
			}
		}
		if addRt {
			*fts = append(*fts, rt)
		}
	}
	return rt
}

// PathPrefix PathPrefix
func (t ReqRouter) PathPrefix(px string) Route {
	rt := t.NewRoute().PathPrefix(px)
	fts := t.prefixRoutes[rt.GetPrefix()]
	if fts == nil {
		t.prefixRoutes[px] = rt
	}
	return rt
}

// EnableCORS EnableCORS
func (t ReqRouter) EnableCORS() {
	t.corsEnabled = true
}

// SetCorsAllowedHeaders SetAllowedHeaders
func (t ReqRouter) SetCorsAllowedHeaders(hdr string) {
	hdr = strings.ReplaceAll(hdr, " ", "")
	headers := strings.Split(hdr, ",")
	for _, v := range headers {
		nHeader := http.CanonicalHeaderKey(strings.TrimSpace(v))
		if nHeader == "" {
			continue
		}
		t.allowedHeaders = append(t.allowedHeaders, nHeader)
	}
}

// SetCorsAllowedOrigins AllowedOrigins
func (t ReqRouter) SetCorsAllowedOrigins(org string) {
	org = strings.ReplaceAll(org, " ", "")
	var origins = strings.Split(org, ",")
	t.allowedOrigins = origins
}

// SetCorsAllowedMethods AllowedMethods
func (t ReqRouter) SetCorsAllowedMethods(mths string) {
	mths = strings.ReplaceAll(mths, " ", "")
	var methods = strings.Split(mths, ",")
	for _, v := range methods {
		nMethod := strings.ToUpper(strings.TrimSpace(v))
		if nMethod == "" {
			continue
		}
		t.allowedMethods = append(t.allowedMethods, nMethod)
	}
}

func (t ReqRouter) handleCors(w http.ResponseWriter) {
	w.Header().Set(corsAllowOriginHeader, strings.Join(t.allowedOrigins, ","))
	w.Header().Set(corsAllowHeadersHeader, strings.Join(t.allowedHeaders, ","))
	w.Header().Set(corsAllowMethodsHeader, strings.Join(t.allowedMethods, ","))
	w.WriteHeader(http.StatusOK)
}

func (t ReqRouter) findPrefix(px string) Route {
	var rtn Route
	rtn = t.prefixRoutes[px]
	return rtn
}

func (t ReqRouter) findRouteAndVars(path string) (Route, *[]string) {
	var rnt Route
	sp := strings.Split(path, "/")
	var vars []string
	var vcnt = len(sp) - 2
	//log.Println("sp:", sp)
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
