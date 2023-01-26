package grrt

// Copyright 2022 GolangToolKits Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

import (
	"log"
	"net/http"
	"strings"
)

// ReqRoute ReqRoute
type ReqRoute struct {
	//namedRoutes map[string]*Route
	handler      http.Handler
	host         string
	path         string
	matcher      Matcher
	active       bool
	pathVarsUsed bool
	pathVarNames *[]string
	methods      *[]string
}

// New New
func (t *ReqRoute) New() Route {
	var m pathMatcher
	t.matcher = m.New()
	t.pathVarNames = &[]string{}
	t.methods = &[]string{}
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
	if t.chechPath(p) && t.matcher.addPath(p) {
		t.pathVarNames, t.path = t.extractPathAndVarNames(p)
		if len(*t.pathVarNames) > 0 {
			t.pathVarsUsed = true
		}
		t.active = true
	}
	return t
}

// Methods Methods
func (t *ReqRoute) Methods(ms ...string) Route {
	var mts []string
	for _, m := range ms {
		mts = append(mts, strings.ToUpper(m))
	}
	t.methods = &mts
	return t
}

// GetMethods GetMethods
func (t *ReqRoute) GetMethods() *[]string {
	return t.methods
}

// IsMethodAllowed IsMethodAllowed
func (t *ReqRoute) IsMethodAllowed(m string) bool {
	var rtn bool
	if len(*t.methods) == 0 {
		rtn = true
	} else {
		for _, mt := range *t.methods {
			if mt == m {
				rtn = true
				break
			}
		}
	}
	return rtn
}

// IsPathVarsUsed IsPathVarsUsed
func (t *ReqRoute) IsPathVarsUsed() bool {
	return t.pathVarsUsed
}

// GetPathVarsCount GetPathVarsCount
func (t *ReqRoute) GetPathVarsCount() int {
	var rtn int = 0
	if t.pathVarNames != nil {
		rtn = len(*t.pathVarNames)
	}
	return rtn
}

// Host Host --future development----
// func (t *ReqRoute) Host(h string) Route {
// 	return nil
// }

// GetHandler GetHandler
func (t *ReqRoute) GetHandler() http.Handler {
	return t.handler
}

// GetPath GetPath
func (t *ReqRoute) GetPath() string {
	return t.path
}

// GetVarNames GetVarNames
func (t *ReqRoute) GetVarNames() *[]string {
	return t.pathVarNames
}

// IsActive IsActive
func (t *ReqRoute) IsActive() bool {
	return t.active
}

func (t *ReqRoute) chechPath(p string) bool {
	var rtn = false
	if t.chechCurlys(p) && t.chechBackSlash(p) && t.chechCurlyPlacement(p) {
		rtn = true
	} else {
		t.printError(p, "Problem with path")
	}
	return rtn
}

func (t *ReqRoute) chechCurlys(p string) bool {
	var rtn = false
	open := rune('{')
	closed := rune('}')
	var cl int = 0
	for _, c := range p {
		if c == open {
			cl++
		} else if c == closed {
			cl--
		}
	}
	if cl == 0 {
		rtn = true
	} else {
		t.printError(p, "Mismatched curly brackets")
	}
	return rtn
}

func (t *ReqRoute) chechCurlyPlacement(p string) bool {
	//checks to make sure there are no {{ of }}
	var rtn = true
	oc := rune('{')
	cc := rune('}')
	var lastOc int = 0
	var lastCc int = 0
	for i, c := range p {
		if c == oc && i == lastOc+1 {
			t.printError(p, "Route can not have {{")
			rtn = false
			break
		} else if c == oc {
			lastOc = i
		}
		if c == cc && i == lastCc+1 {
			t.printError(p, "Route can not have }}")
			rtn = false
			break
		} else if c == cc {
			lastCc = i
		}
	}
	return rtn
}

func (t *ReqRoute) chechBackSlash(p string) bool {
	var rtn = true
	bs := rune('/')
	var lastBs int = 0
	for i, c := range p {
		if i == 0 && c != bs {
			rtn = false
			t.printError(p, "Route must have leading backslash")
			break
		} else if i != 0 {
			if c == bs && i == lastBs+1 {
				rtn = false
				t.printError(p, "Misplaced backslash")
				break
			} else if c == bs && i != len(p)-1 {
				lastBs = i
			} else if c == bs && i == len(p)-1 {
				rtn = false
				t.printError(p, "Backslash not allowed on end of route")
			}
		}
	}
	return rtn
}

func (t *ReqRoute) extractPathAndVarNames(p string) (*[]string, string) {
	var vars = []string{}
	var pth string
	// rtn:= make([]string,1)
	oc := rune('{')
	cc := rune('}')
	var pend int
	var ind1 int
	var ind2 int
	for i, c := range p {
		if c == oc {
			ind1 = i
			if pend == 0 {
				pend = i
			}
		}
		if c == cc {
			ind2 = i
		}
		if ind1 > 0 && ind1 < ind2 {
			pvar := p[ind1+1 : ind2]
			vars = append(vars, pvar)
			ind1 = 0
			ind2 = 0
		}
	}
	if pend > 0 {
		if pth = p[0 : pend-1]; len(pth) == 0 {
			pth = p[0:pend]
		}
		//pth = p[0 : pend-1]
	} else {
		pth = p
	}
	return &vars, pth
}

func (t *ReqRoute) printError(p string, cause string) {
	log.Println("Path not added to route:")
	log.Println(p)
	log.Println(cause)
}
