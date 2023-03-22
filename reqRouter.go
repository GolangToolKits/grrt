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
	varsKey                    contextKey = iota
	corsAllowOriginHeader      string     = "Access-Control-Allow-Origin"
	corsAllowHeadersHeader     string     = "Access-Control-Allow-Headers"
	corsAllowMethodsHeader     string     = "Access-Control-Allow-Methods"
	corsAllowCredentialsHeader string     = "Access-Control-Allow-Credentials"
)

// ReqRouter RequestRouter
type ReqRouter struct {
	namedRoutes          map[string]*[]Route
	prefixRoutes         map[string]Route
	corsEnabled          bool
	corsAllowCredentials bool
	allowedHeaders       []string
	allowedOrigins       []string
	allowedMethods       []string
}

// ServeHTTP ServeHTTP dispatches the handler registered in the matched route.
func (t *ReqRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// When there is a match, the route variables can be retrieved calling
	// mux.Vars(request).

	if t.corsEnabled && r.Method == http.MethodOptions {
		t.handleCors(w)
	} else {
		path := r.URL.Path
		var rt Route
		sfile := t.isStaticFile(path)
		if !sfile {
			frt, fvars := t.findRouteAndVars(path)
			rt = frt
			if rt != nil && len(*fvars) > 0 {
				r = t.requestWithVars(r, rt.GetVarNames(), fvars)
				// rt = frt
			}
		}
		if rt == nil && !sfile {
			rt = t.findPrefix(path)
		} else if rt == nil && sfile {
			rt = t.findFilePrefix(path)
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
}

// NewRoute NewRoute
func (t *ReqRouter) NewRoute() Route {
	var rt ReqRoute
	rrt := rt.New()
	return rrt
}

// Handle Handle
func (t *ReqRouter) Handle(path string, handler http.Handler) Route {
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
func (t *ReqRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Route {
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
func (t *ReqRouter) PathPrefix(px string) Route {
	rt := t.NewRoute().PathPrefix(px)
	fts := t.prefixRoutes[rt.GetPrefix()]
	if fts == nil {
		t.prefixRoutes[px] = rt
	}
	return rt
}

// EnableCORS EnableCORS
func (t *ReqRouter) EnableCORS() {
	t.corsEnabled = true
}

// CORSAllowCredentials CORSAllowCredentials
func (t *ReqRouter) CORSAllowCredentials() {
	t.corsAllowCredentials = true
}

// SetCorsAllowedHeaders SetAllowedHeaders
func (t *ReqRouter) SetCorsAllowedHeaders(hdr string) {
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
func (t *ReqRouter) SetCorsAllowedOrigins(org string) {
	org = strings.ReplaceAll(org, " ", "")
	var origins = strings.Split(org, ",")
	t.allowedOrigins = origins
}

// SetCorsAllowedMethods AllowedMethods
func (t *ReqRouter) SetCorsAllowedMethods(mths string) {
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

func (t *ReqRouter) handleCors(w http.ResponseWriter) {
	w.Header().Set(corsAllowOriginHeader, strings.Join(t.allowedOrigins, ", "))
	w.Header().Set(corsAllowHeadersHeader, strings.Join(t.allowedHeaders, ", "))
	w.Header().Set(corsAllowMethodsHeader, strings.Join(t.allowedMethods, ", "))
	if t.corsAllowCredentials {
		w.Header().Set(corsAllowCredentialsHeader, "true")
	}
	w.WriteHeader(http.StatusOK)
}

func (t *ReqRouter) findPrefix(px string) Route {
	var rtn Route
	rtn = t.prefixRoutes[px]
	return rtn
}

func (t *ReqRouter) findFilePrefix(path string) Route {
	var rtn Route
	sp := strings.Split(path, "/")
	if len(sp) > 1 {
		rtn = t.prefixRoutes[sp[1]]
		if rtn == nil {
			rtn = t.prefixRoutes["/"]
		}
	}
	return rtn
}

func (t *ReqRouter) findRouteAndVars(path string) (Route, *[]string) {
	var rnt Route
	sp := strings.Split(path, "/")
	plen := len(sp)
	var vars []string
	var searchPath = path
	var found = false
	for range sp {
		//fmt.Println(i)
		if found {
			break
		}
		rts := t.namedRoutes[searchPath]
		if rts != nil {
			for _, rt := range *rts {
				if rt.IsPathVarsUsed() {
					vars = sp[plen:]
					if len(vars) == rt.GetPathVarsCount() {
						found = true
						rnt = rt
						break
					}
				} else if rt.GetPath() == path {
					found = true
					rnt = rt
					break
				}
			}
		} else {
			plen--
			si := strings.LastIndex(searchPath, "/")
			if si == 0 {
				si++
			}
			searchPath = searchPath[:si]
		}
	}
	return rnt, &vars
}

func (t *ReqRouter) requestWithVars(r *http.Request, pVarNames, pvars *[]string) *http.Request {
	var vars = make(map[string]string)
	if len(*pVarNames) == len(*pvars) {
		for i, n := range *pVarNames {
			vars[n] = (*pvars)[i]
		}
	}
	ctx := context.WithValue(r.Context(), varsKey, vars)
	return r.WithContext(ctx)
}

func (t *ReqRouter) isStaticFile(path string) bool {
	var rtn bool
	var ind = strings.LastIndex(path, ".")
	var ind2 = strings.LastIndex(path, "@")
	if ind > 0 && ind2 < 0 {
		disp := len(path) - (ind + 1)
		if disp < 4 {
			rtn = true
		}
	}
	return rtn
}
