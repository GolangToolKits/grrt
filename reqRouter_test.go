package grrt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReqRouter_NewRoute(t *testing.T) {

	var vs = &[]string{}
	var mt = &[]string{}
	type fields struct {
		namedRoutes map[string]*[]Route
	}
	tests := []struct {
		name   string
		fields fields
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",

			want: &ReqRoute{

				pathVarNames: vs,
				methods:      mt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			if got := tr.NewRoute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRouter.NewRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_HandleFunc(t *testing.T) {

	var nr = make(map[string]*[]Route)
	var nf func(http.ResponseWriter, *http.Request)
	var nr2 = make(map[string]*[]Route)
	var rt2 = &ReqRoute{
		path:         "/route/test1",
		active:       true,
		pathVarsUsed: true,
		pathVarNames: &[]string{"id"},
		methods:      &[]string{},
	}
	nr2["/route/test1"] = &[]Route{rt2}

	type fields struct {
		namedRoutes map[string]*[]Route
	}
	type args struct {
		path string
		f    func(http.ResponseWriter, *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				namedRoutes: nr,
			},
			args: args{
				path: "/route/test1",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},
				methods:      &[]string{},
			},
		},
		{
			name: "test 2",
			fields: fields{
				namedRoutes: nr2,
			},
			args: args{
				path: "/route/test1/{name}/{cat}",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name", "cat"},
				methods:      &[]string{},
			},
		},
		{
			name: "test 3",
			fields: fields{
				namedRoutes: nr2,
			},
			args: args{
				path: "/route/test1/{name}",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name"},
				methods:      &[]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			// if got := tr.HandleFunc(tt.args.path, tt.args.f); !reflect.DeepEqual(got, tt.want) {
			if got := tr.HandleFunc(tt.args.path, tt.args.f); got.GetPath() != "/route/test1" {

				t.Errorf("ReqRouter.HandleFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_HandleFuncFullCall(t *testing.T) {
	var nr = make(map[string]*[]Route)
	var nf func(http.ResponseWriter, *http.Request)

	type fields struct {
		namedRoutes map[string]*[]Route
	}
	type args struct {
		path string
		f    func(http.ResponseWriter, *http.Request)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				namedRoutes: nr,
			},
			args: args{
				path: "/route/test1",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},

				methods: &[]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			// if got := tr.HandleFunc(tt.args.path, tt.args.f); !reflect.DeepEqual(got, tt.want) {
			if got := tr.HandleFunc(tt.args.path, tt.args.f).Methods("post", "put"); got.GetPath() != "/route/test1" ||
				got.GetHandler() == nil || len(*got.GetMethods()) != 2 ||
				(*got.GetMethods())[0] != "POST" ||
				len(tr.namedRoutes) != 1 {
				fmt.Println("got", got)
				var ms = got.GetMethods()
				fmt.Println("got method 1", (*ms)[0])
				fmt.Println("routes", tr.namedRoutes)
				t.Errorf("ReqRouter.HandleFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_ServeHTTP(t *testing.T) {
	var nf = func(http.ResponseWriter, *http.Request) {

	}
	var hdl = http.HandlerFunc(nf)

	var rt ReqRoute
	rt.active = true
	rt.path = "/test/test1"
	rt.handler = hdl
	rt.methods = &[]string{"POST"}
	var rts = make(map[string]*[]Route)

	var prt []Route
	prt = append(prt, &rt)
	rts[rt.path] = &prt

	var rt2 ReqRoute
	rt2.active = true
	rt2.pathVarsUsed = true
	rt2.path = "/test/test1"
	rt2.handler = hdl
	rt2.methods = &[]string{"GET"}
	rt2.pathVarNames = &[]string{"param1", "param2"}
	var rts2 = make(map[string]*[]Route)

	var prt2 []Route
	prt2 = append(prt2, &rt2)
	rts2[rt2.path] = &prt2

	var prrt ReqRoute
	prrt.active = true
	prrt.isPrefix = true
	prrt.prefix = "/testPrefix/"
	prrt.handler = hdl
	prrt.methods = &[]string{}
	var prrts = make(map[string]Route)

	//var prt []Route
	//prt = append(prt, &rt)
	prrts[prrt.prefix] = &prrt

	tw := httptest.NewRecorder()
	tw2 := httptest.NewRecorder()
	tw3 := httptest.NewRecorder()
	tw4 := httptest.NewRecorder()
	tw5 := httptest.NewRecorder()
	tw6 := httptest.NewRecorder()
	tw7 := httptest.NewRecorder()

	tr, _ := http.NewRequest("POST", "/test/test1", nil)
	tr22, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tr2, _ := http.NewRequest("POST", "/test/te", nil)
	tr3, _ := http.NewRequest("PUT", "/test/test1", nil)
	tr4, _ := http.NewRequest("POST", "/test/test1/param1/param2", nil)
	tr6, _ := http.NewRequest("POST", "/testPrefix/", nil)
	tr7, _ := http.NewRequest("OPTIONS", "/testPrefix/", nil)
	// var prt []Route
	// prt = append(prt, &rt)
	// rts[rt.path] = &prt

	type fields struct {
		namedRoutes  map[string]*[]Route
		prefixRoutes map[string]Route
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw,
				r: tr,
			},
		},
		{
			name: "test 2 404",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw2,
				r: tr2,
			},
		},
		{
			name: "test 3 405",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw3,
				r: tr3,
			},
		},
		{
			name: "test 4 405",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw4,
				r: tr4,
			},
		},
		{
			name: "test 5",
			fields: fields{
				namedRoutes: rts2,
			},
			args: args{
				w: tw5,
				r: tr22,
			},
		},
		{
			name: "test 6 prefix",
			fields: fields{
				prefixRoutes: prrts,
			},
			args: args{
				w: tw6,
				r: tr6,
			},
		},
		{
			name: "cors test",
			args: args{
				w: tw7,
				r: tr7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:  tt.fields.namedRoutes,
				prefixRoutes: tt.fields.prefixRoutes,
				corsEnabled:  true,
			}
			tr.ServeHTTP(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 2 404" && tw2.Code != http.StatusNotFound {
				t.Fail()
			}
			if tt.name == "test 3 405" && tw3.Code != http.StatusMethodNotAllowed {
				t.Fail()
			}
			if tt.name == "test 4 405" && tw3.Code != http.StatusMethodNotAllowed {
				t.Fail()
			}
			if tt.name == "test 5" && tw5.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 6 prefix" && tw6.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "cors test" && tw7.Code != http.StatusOK {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_findRoute(t *testing.T) {
	//var hdl = http.HandlerFunc(nf)
	var varNames = &[]string{"param1", "param2"}
	var methods = &[]string{"GET"}
	var rt ReqRoute
	rt.active = true
	rt.path = "/test/test1"
	rt.pathVarsUsed = true
	rt.pathVarNames = varNames
	//rt.handler = hdl
	rt.methods = methods
	var rts = make(map[string]*[]Route)
	var prt []Route
	prt = append(prt, &rt)
	rts[rt.path] = &prt

	//test 2 ----
	var varNames2 = &[]string{}
	var methods2 = &[]string{"GET"}
	var rt2 ReqRoute
	rt2.active = true
	rt2.path = "/test/test1/test2/test3"
	rt2.pathVarsUsed = false
	rt2.pathVarNames = varNames2
	//rt.handler = hdl
	rt2.methods = methods2
	var rts2 = make(map[string]*[]Route)
	var prt2 []Route
	prt2 = append(prt2, &rt2)
	rts2[rt2.path] = &prt2

	//test 3 ----
	var varNames3 = &[]string{}
	var methods3 = &[]string{"POST"}
	var rt3 ReqRoute
	rt3.active = true
	rt3.path = "/test/test1/test2"
	rt3.pathVarsUsed = false
	rt3.pathVarNames = varNames3
	//rt.handler = hdl
	rt3.methods = methods3

	var varNames33 = &[]string{"param1"}
	var methods33 = &[]string{"GET"}
	var rt33 ReqRoute
	rt33.active = true
	rt33.path = "/test/test1/test2"
	rt33.pathVarsUsed = true
	rt33.pathVarNames = varNames33
	//rt.handler = hdl
	rt33.methods = methods33
	var rts3 = make(map[string]*[]Route)
	var prt3 []Route
	prt3 = append(prt3, &rt3)
	prt3 = append(prt3, &rt33)
	rts3[rt33.path] = &prt3

	type fields struct {
		namedRoutes map[string]*[]Route
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
		want1  *[]string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				path: "/test/test1/p1/p2",
			},
			want: &ReqRoute{
				path:         "/test/test1",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: varNames,

				methods: methods,
			},
			want1: &[]string{"p1", "p2"},
		},
		{
			name: "test 2",
			fields: fields{
				namedRoutes: rts2,
			},
			args: args{
				path: "/test/test1/test2/test3",
			},
			want: &ReqRoute{
				path:         "/test/test1/test2/test3",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: varNames2,
				methods:      methods2,
			},
			want1: &[]string{},
		},
		{
			name: "test 3",
			fields: fields{
				namedRoutes: rts3,
			},
			args: args{
				path: "/test/test1/test2/p1",
			},
			want: &ReqRoute{
				path:         "/test/test1/test2",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: varNames33,

				methods: methods33,
			},
			want1: &[]string{"p1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			got, got1 := tr.findRouteAndVars(tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRouter.findRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReqRouter.findRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReqRouter_requestWithVars(t *testing.T) {
	// var rts = make(map[string]*[]Route)
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	type fields struct {
		namedRoutes map[string]*[]Route
	}
	type args struct {
		r         *http.Request
		pVarNames *[]string
		pvars     *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *http.Request
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				r:         tr,
				pVarNames: &[]string{"param1", "param2"},
				pvars:     &[]string{"p1", "p2"},
			},
			want: tr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			if got := tr.requestWithVars(tt.args.r, tt.args.pVarNames, tt.args.pvars); reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRouter.requestWithVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_PathPrefix(t *testing.T) {
	// r1 := ReqRoute{
	// 	prefix:   "/testPrefix/",
	// 	active:   true,
	// 	isPrefix: true,
	// }
	//nr  := make(map[string]*[]Route)
	tm1 := make(map[string]Route)
	//tm1["/testPrefix/"] = &r1

	type fields struct {
		namedRoutes  map[string]*[]Route
		prefixRoutes map[string]Route
		corsEnabled  bool
	}
	type args struct {
		px string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				prefixRoutes: tm1,
			},
			args: args{
				px: "/testPrefix/",
			},
			want: &ReqRoute{
				prefix:   "/testPrefix/",
				active:   true,
				isPrefix: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:  tt.fields.namedRoutes,
				prefixRoutes: tt.fields.prefixRoutes,
				corsEnabled:  tt.fields.corsEnabled,
			}
			if got := tr.PathPrefix(tt.args.px); got.GetPrefix() != tt.want.GetPrefix() {
				t.Errorf("ReqRouter.PathPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_Handle(t *testing.T) {
	var nr = make(map[string]*[]Route)
	var nf func(http.ResponseWriter, *http.Request)
	var hdl = http.HandlerFunc(nf)

	var nr2 = make(map[string]*[]Route)
	nrt2 := &ReqRoute{
		path:         "/route/test1",
		active:       true,
		pathVarsUsed: true,
		pathVarNames: &[]string{"name"},
		methods:      &[]string{},
	}
	nr2["/route/test1"] = &[]Route{nrt2}

	type fields struct {
		namedRoutes  map[string]*[]Route
		prefixRoutes map[string]Route
		corsEnabled  bool
	}
	type args struct {
		path    string
		handler http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				namedRoutes: nr,
			},
			args: args{
				path:    "/route/test1",
				handler: hdl,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},
				methods:      &[]string{},
			},
		},
		{
			name: "test 2",
			fields: fields{
				namedRoutes: nr,
			},
			args: args{
				path:    "/route/test1/{name}/{cat}",
				handler: hdl,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name", "cat"},
				methods:      &[]string{},
			},
		},
		{
			name: "test 3",
			fields: fields{
				namedRoutes: nr2,
			},
			args: args{
				path:    "/route/test1/{name}",
				handler: hdl,
			},
			want: &ReqRoute{
				path:         "/route/test1",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name"},
				methods:      &[]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:  tt.fields.namedRoutes,
				prefixRoutes: tt.fields.prefixRoutes,
				corsEnabled:  tt.fields.corsEnabled,
			}
			if got := tr.Handle(tt.args.path, tt.args.handler); got.GetPath() != tt.want.GetPath() ||
				got.IsActive() != tt.want.IsActive() || got.IsPathVarsUsed() != tt.want.IsPathVarsUsed() {
				t.Errorf("ReqRouter.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_findPrefix(t *testing.T) {
	px1 := &ReqRoute{
		prefix:   "/testPrefix/",
		isPrefix: true,
	}
	pxm := make(map[string]Route)
	pxm["/testPrefix/"] = px1

	type fields struct {
		namedRoutes  map[string]*[]Route
		prefixRoutes map[string]Route
		corsEnabled  bool
	}
	type args struct {
		px string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				prefixRoutes: pxm,
			},
			args: args{
				px: "/testPrefix/",
			},
			want: &ReqRoute{
				prefix:   "/testPrefix/",
				isPrefix: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:  tt.fields.namedRoutes,
				prefixRoutes: tt.fields.prefixRoutes,
				corsEnabled:  tt.fields.corsEnabled,
			}
			if got := tr.findPrefix(tt.args.px); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRouter.findPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_SetAllowedHeaders(t *testing.T) {
	type fields struct {
		namedRoutes    map[string]*[]Route
		prefixRoutes   map[string]Route
		corsEnabled    bool
		allowedHeaders []string
		allowedOrigins []string
		allowedMethods []string
	}
	type args struct {
		headers string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				allowedHeaders: []string{"Content-Type", "api-key"},
			},
			args: args{
				headers: "Content-Type, api-key, ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:    tt.fields.namedRoutes,
				prefixRoutes:   tt.fields.prefixRoutes,
				corsEnabled:    tt.fields.corsEnabled,
				allowedHeaders: tt.fields.allowedHeaders,
				allowedOrigins: tt.fields.allowedOrigins,
				allowedMethods: tt.fields.allowedMethods,
			}
			tr.SetCorsAllowedHeaders(tt.args.headers)
			if tr.allowedHeaders[0] != "Content-Type" {
				t.Fail()
			}
			if tr.allowedHeaders[1] != "api-key" {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_EnableCORS(t *testing.T) {
	type fields struct {
		namedRoutes    map[string]*[]Route
		prefixRoutes   map[string]Route
		corsEnabled    bool
		allowedHeaders []string
		allowedOrigins []string
		allowedMethods []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				corsEnabled: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:    tt.fields.namedRoutes,
				prefixRoutes:   tt.fields.prefixRoutes,
				corsEnabled:    tt.fields.corsEnabled,
				allowedHeaders: tt.fields.allowedHeaders,
				allowedOrigins: tt.fields.allowedOrigins,
				allowedMethods: tt.fields.allowedMethods,
			}
			tr.EnableCORS()
		})
	}
}

func TestReqRouter_AllowedOrigins(t *testing.T) {
	type fields struct {
		namedRoutes    map[string]*[]Route
		prefixRoutes   map[string]Route
		corsEnabled    bool
		allowedHeaders []string
		allowedOrigins []string
		allowedMethods []string
	}
	type args struct {
		origins string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				allowedOrigins: []string{"*", "test"},
			},
			args: args{
				origins: "*,   test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:    tt.fields.namedRoutes,
				prefixRoutes:   tt.fields.prefixRoutes,
				corsEnabled:    tt.fields.corsEnabled,
				allowedHeaders: tt.fields.allowedHeaders,
				allowedOrigins: tt.fields.allowedOrigins,
				allowedMethods: tt.fields.allowedMethods,
			}
			tr.SetCorsAllowedOrigins(tt.args.origins)
			if tr.allowedOrigins[0] != "*" {
				t.Fail()
			}
			if tr.allowedOrigins[1] != "test" {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_AllowedMethods(t *testing.T) {
	type fields struct {
		namedRoutes    map[string]*[]Route
		prefixRoutes   map[string]Route
		corsEnabled    bool
		allowedHeaders []string
		allowedOrigins []string
		allowedMethods []string
	}
	type args struct {
		methods string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				allowedMethods: []string{"POST"},
			},
			args: args{
				methods: "POST,   ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:    tt.fields.namedRoutes,
				prefixRoutes:   tt.fields.prefixRoutes,
				corsEnabled:    tt.fields.corsEnabled,
				allowedHeaders: tt.fields.allowedHeaders,
				allowedOrigins: tt.fields.allowedOrigins,
				allowedMethods: tt.fields.allowedMethods,
			}
			tr.SetCorsAllowedMethods(tt.args.methods)
			if tr.allowedMethods[0] != "POST" {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_handleCors(t *testing.T) {
	tw := httptest.NewRecorder()
	type fields struct {
		namedRoutes    map[string]*[]Route
		prefixRoutes   map[string]Route
		corsEnabled    bool
		allowedHeaders []string
		allowedOrigins []string
		allowedMethods []string
	}
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				allowedHeaders: []string{"Content-Type"},
				allowedOrigins: []string{"test"},
				allowedMethods: []string{"POST", "GET"},
			},
			args: args{
				w: tw,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:    tt.fields.namedRoutes,
				prefixRoutes:   tt.fields.prefixRoutes,
				corsEnabled:    tt.fields.corsEnabled,
				allowedHeaders: tt.fields.allowedHeaders,
				allowedOrigins: tt.fields.allowedOrigins,
				allowedMethods: tt.fields.allowedMethods,
			}

			tr.handleCors(tt.args.w)
			fmt.Println("headers: ", tw.Header().Get("Access-Control-Allow-Headers"))
			if tw.Result().StatusCode != http.StatusOK {
				t.Fail()
			}
			if tw.Header().Get("Access-Control-Allow-Headers") != "Content-Type" {
				t.Fail()
			}
			if tw.Header().Get("Access-Control-Allow-Origin") != "test" {
				t.Fail()
			}
			if tw.Header().Get("Access-Control-Allow-Methods") != "POST, GET" {
				t.Fail()
			}
		})
	}
}
