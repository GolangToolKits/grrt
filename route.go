package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import "net/http"

// Route Route
type Route interface {
	Handler(handler http.Handler) Route
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) Route
	Path(p string) Route
	PathPrefix(p string) Route
	GetHandler() http.Handler
	GetPath() string
	GetPrefix() string
	GetVarNames() *[]string
	IsActive() bool
	Methods(ms ...string) Route
	GetMethods() *[]string
	IsMethodAllowed(m string) bool
	IsPathVarsUsed() bool
	GetPathVarsCount() int
	//-------for future development---------------------------
	//Host(h string) Route
}
