package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import "net/http"

// ReqRoute ReqRoute
type ReqRoute struct {
	//namedRoutes map[string]*Route
	handler http.Handler
	host    string
	path    string
	matcher Matcher
	active  bool
}

// New New
func (t *ReqRoute) New() Route {
	var m pathMatcher
	t.matcher = m.New()
	return t
}

// Handler Handler
func (t *ReqRoute) Handler(handler http.Handler) Route {
	if t.active {
		t.handler = handler
	}
	return t
}

// HandlerFunc HandlerFunc
func (t *ReqRoute) HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route {
	return t.Handler(http.HandlerFunc(f))
}

// Path Path
func (t *ReqRoute) Path(p string) Route {
	if t.matcher.addPath(p) {
		t.path = p
		t.active = true
	}
	return t
}

// Host Host
func (t *ReqRoute) Host(h string) Route {
	return nil
}

// GetHandler GetHandler
func (t *ReqRoute) GetHandler() http.Handler {
	return t.handler
}
