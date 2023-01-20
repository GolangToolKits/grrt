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
	GetHandler()http.Handler
	Host(h string) Route
}
