package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Matcher Matcher
type Matcher interface {
	addPath(p string) bool
}

type pathMatcher struct {
	paths map[string]bool
}

func (m *pathMatcher) New() Matcher {
	m.paths = make(map[string]bool)
	return m
}

func matchHost(h string) bool {
	return false
}

func matchPath(p string) bool {
	return false
}

func (m *pathMatcher) addPath(p string) bool {
	var rtn = false
	if !m.paths[p] {
		m.paths[p] = true
		rtn = true
	}
	return rtn
}
