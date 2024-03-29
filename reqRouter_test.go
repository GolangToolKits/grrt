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

	var nr3 = make(map[string]*[]Route)
	var rt3 = &ReqRoute{
		path:         "/",
		active:       true,
		pathVarsUsed: false,
		pathVarNames: &[]string{},
		methods:      &[]string{},
	}
	nr3["/"] = &[]Route{rt3}

	type fields struct {
		namedRoutes map[string]*[]Route
	}
	type args struct {
		path string
		f    func(http.ResponseWriter, *http.Request)
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Route
		wantPath string
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
			wantPath: "/route/test1",
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
			wantPath: "/route/test1",
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
			wantPath: "/route/test1",
		},
		{
			name: "test 4",
			fields: fields{
				namedRoutes: nr,
			},
			args: args{
				path: "/",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},
				methods:      &[]string{},
			},
			wantPath: "/",
		},
		{
			name: "test 5",
			fields: fields{
				namedRoutes: nr3,
			},
			args: args{
				path: "/{name}",
				f:    nf,
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name"},
				methods:      &[]string{},
			},
			wantPath: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			// if got := tr.HandleFunc(tt.args.path, tt.args.f); !reflect.DeepEqual(got, tt.want) {
			if got := tr.HandleFunc(tt.args.path, tt.args.f); got.GetPath() != tt.wantPath {

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
	var fvars map[string]string
	var nf = func(wnf http.ResponseWriter, rnf *http.Request) {
		fvars = Vars(rnf)
	}
	var hdl = http.HandlerFunc(nf)

	var rts = make(map[string]*[]Route)

	var rt ReqRoute
	rt.active = true
	rt.path = "/test/test1"
	rt.handler = hdl
	rt.methods = &[]string{"POST"}
	// var rts = make(map[string]*[]Route)

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
	//var rts2 = make(map[string]*[]Route)

	//var prt2 []Route
	prt = append(prt, &rt2)
	rts[rt2.path] = &prt

	var rt8 ReqRoute
	rt8.active = true
	rt8.path = "/"
	rt8.handler = hdl
	rt8.methods = &[]string{"GET"}

	//var rts8 = make(map[string]*[]Route)

	var prt8 []Route
	prt8 = append(prt8, &rt8)
	rts[rt8.path] = &prt8

	var rt88 ReqRoute
	rt88.active = true
	rt88.pathVarsUsed = true
	rt88.path = "/"
	rt88.handler = hdl
	rt88.methods = &[]string{"GET"}
	rt88.pathVarNames = &[]string{"param1"}
	prt8 = append(prt8, &rt88)
	rts[rt88.path] = &prt8

	// var prt8 []Route
	// prt8 = append(prt8, &rt8)
	// rts8[rt8.path] = &prt8

	var rt10 ReqRoute
	rt10.active = true
	rt10.pathVarsUsed = true
	rt10.path = "/product"
	rt10.handler = hdl
	rt10.methods = &[]string{"GET"}
	rt10.pathVarNames = &[]string{"id", "sku"}

	var rt10b ReqRoute
	rt10b.active = true
	rt10b.pathVarsUsed = false
	rt10b.path = "/"
	rt10b.handler = hdl
	rt10b.methods = &[]string{"GET"}

	var rts10 = make(map[string]*[]Route)

	var prt10 []Route
	prt10 = append(prt10, &rt10)

	rts10[rt10.path] = &prt10

	var prt10b []Route
	prt10b = append(prt10b, &rt10b)
	rts10[rt10b.path] = &prt10b

	////=--------

	var rt11 ReqRoute
	rt11.active = true
	rt11.pathVarsUsed = true
	rt11.path = "/view"
	rt11.handler = hdl
	rt11.methods = &[]string{"GET"}
	rt11.pathVarNames = &[]string{"name"}

	var rt11b ReqRoute
	rt11b.active = true
	rt11b.pathVarsUsed = false
	rt11b.path = "/"
	rt11b.handler = hdl
	rt11b.methods = &[]string{"GET"}

	var rt11c ReqRoute
	rt11c.active = true
	rt11c.pathVarsUsed = false
	rt11c.path = "/"
	rt11c.handler = hdl
	rt11c.methods = &[]string{"GET"}
	rt11c.pathVarNames = &[]string{"name"}

	var rts11 = make(map[string]*[]Route)

	var prt11 []Route
	prt11 = append(prt11, &rt11)

	rts11[rt11.path] = &prt11

	var prt11b []Route
	prt11b = append(prt11b, &rt11b)
	rts11[rt11b.path] = &prt11b

	var prt11c []Route
	prt11c = append(prt11c, &rt11c)
	rts11[rt11c.path] = &prt11c

	var prrt ReqRoute
	prrt.active = true
	prrt.isPrefix = true
	prrt.prefix = "/testPrefix/"
	prrt.handler = hdl
	prrt.methods = &[]string{}

	var prrt2 ReqRoute
	prrt2.active = true
	prrt2.isPrefix = true
	prrt2.prefix = "/"
	prrt2.handler = hdl
	prrt2.methods = &[]string{}
	var prrts = make(map[string]Route)

	//var pfrt []Route
	//pfrt = append(pfrt, &prrt)
	prrts[prrt.prefix] = &prrt
	prrts[prrt2.prefix] = &prrt2

	tw := httptest.NewRecorder()
	tw2 := httptest.NewRecorder()
	tw3 := httptest.NewRecorder()
	tw4 := httptest.NewRecorder()
	tw5 := httptest.NewRecorder()
	tw6 := httptest.NewRecorder()
	tw6b := httptest.NewRecorder()
	tw7 := httptest.NewRecorder()
	tw8 := httptest.NewRecorder()
	tw9 := httptest.NewRecorder()
	tw10 := httptest.NewRecorder()
	tw11 := httptest.NewRecorder()

	tr, _ := http.NewRequest("POST", "/test/test1", nil)
	tr22, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tr2, _ := http.NewRequest("POST", "/test/te", nil)
	tr3, _ := http.NewRequest("PUT", "/test/test1", nil)
	tr4, _ := http.NewRequest("PUT", "/test/test1/param1/param2", nil)
	tr6, _ := http.NewRequest("POST", "/testPrefix/", nil)
	tr6b, _ := http.NewRequest("GET", "/templates/cleanblog/vendor/fontawesome-free/css/all.min.css", nil)
	tr7, _ := http.NewRequest("OPTIONS", "/testPrefix/", nil)

	tr8, _ := http.NewRequest("GET", "/", nil)
	tr9, _ := http.NewRequest("GET", "/p1", nil)
	tr10, _ := http.NewRequest("GET", "/product/p1/p2", nil)
	tr11, _ := http.NewRequest("GET", "/view/p1", nil)
	//tr12, _ := http.NewRequest("GET", "/templates/cleanblog/vendor/fontawesome-free/css/all.min.css", nil)
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
		name       string
		fields     fields
		args       args
		wantCode   int
		wantW      *httptest.ResponseRecorder
		wantVarLen int
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
			wantW:      tw,
			wantCode:   http.StatusOK,
			wantVarLen: 0,
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
			wantW:      tw2,
			wantCode:   http.StatusNotFound,
			wantVarLen: 0,
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
			wantW:      tw3,
			wantCode:   http.StatusMethodNotAllowed,
			wantVarLen: 0,
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
			wantW:      tw4,
			wantCode:   http.StatusMethodNotAllowed,
			wantVarLen: 0,
		},
		{
			name: "test 5",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw5,
				r: tr22,
			},
			wantW:      tw5,
			wantCode:   http.StatusOK,
			wantVarLen: 2,
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
			wantW:      tw6,
			wantCode:   http.StatusOK,
			wantVarLen: 0,
		},
		{
			name: "test 6 prefix2",
			fields: fields{
				prefixRoutes: prrts,
			},
			args: args{
				w: tw6b,
				r: tr6b,
			},
			wantW:      tw6b,
			wantCode:   http.StatusOK,
			wantVarLen: 0,
		},
		{
			name: "cors test",
			args: args{
				w: tw7,
				r: tr7,
			},
			wantW:      tw7,
			wantCode:   http.StatusOK,
			wantVarLen: 0,
		},
		{
			name: "test 8",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw8,
				r: tr8,
			},
			wantW:      tw8,
			wantCode:   http.StatusOK,
			wantVarLen: 0,
		},
		{
			name: "test 9",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				w: tw9,
				r: tr9,
			},
			wantW:      tw9,
			wantCode:   http.StatusOK,
			wantVarLen: 1,
		},
		{
			name: "test 10",
			fields: fields{
				namedRoutes: rts10,
			},
			args: args{
				w: tw10,
				r: tr10,
			},
			wantW:      tw10,
			wantCode:   http.StatusOK,
			wantVarLen: 2,
		},
		{
			name: "test 11",
			fields: fields{
				namedRoutes: rts11,
			},
			args: args{
				w: tw11,
				r: tr11,
			},
			wantW:      tw11,
			wantCode:   http.StatusOK,
			wantVarLen: 1,
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
			//vs := Vars(tt.args.r)
			if tt.wantW.Code != tt.wantCode || len(fvars) != tt.wantVarLen {
				t.Fail()
			}
			// if tt.name == "test 1" && tw.Code != http.StatusOK {
			// 	t.Fail()
			// }
			// if tt.name == "test 2 404" && tw2.Code != http.StatusNotFound {
			// 	t.Fail()
			// }
			// if tt.name == "test 3 405" && tw3.Code != http.StatusMethodNotAllowed {
			// 	t.Fail()
			// }
			// if tt.name == "test 4 405" && tw3.Code != http.StatusMethodNotAllowed {
			// 	t.Fail()
			// }
			// if tt.name == "test 5" && tw5.Code != http.StatusOK {
			// 	t.Fail()
			// }
			// if tt.name == "test 6 prefix" && tw6.Code != http.StatusOK {
			// 	t.Fail()
			// }
			// if tt.name == "cors test" && tw7.Code != http.StatusOK {
			// 	t.Fail()
			// }
		})
	}
}

func TestReqRouter_findRoute(t *testing.T) {
	//var hdl = http.HandlerFunc(nf)
	var rts = make(map[string]*[]Route)

	//var varNames = &[]string{"param1", "param2"}
	//var methods = &[]string{"GET"}
	var rt ReqRoute
	rt.active = true
	rt.path = "/test/test1"
	rt.pathVarsUsed = true
	rt.pathVarNames = &[]string{"param1", "param2"}
	//rt.handler = hdl
	rt.methods = &[]string{"GET"}
	// var rts = make(map[string]*[]Route)
	var prt []Route
	prt = append(prt, &rt)
	rts[rt.path] = &prt

	//test 2 ----
	//var varNames2 = &[]string{}
	//var methods2 = &[]string{"GET"}
	var rt2 ReqRoute
	rt2.active = true
	rt2.path = "/test/test1/test2/test3"
	rt2.pathVarsUsed = false
	rt2.pathVarNames = &[]string{}
	//rt.handler = hdl
	rt2.methods = &[]string{"GET"}
	//var rts2 = make(map[string]*[]Route)
	var prt2 []Route
	prt2 = append(prt2, &rt2)
	rts[rt2.path] = &prt2

	//test 3 ----
	//var varNames3 = &[]string{}
	//var methods3 = &[]string{"POST"}
	var rt3 ReqRoute
	rt3.active = true
	rt3.path = "/test/test1/test2"
	rt3.pathVarsUsed = false
	rt3.pathVarNames = &[]string{}
	//rt.handler = hdl
	rt3.methods = &[]string{"POST"}

	//var varNames33 = &[]string{"param1"}
	//var methods33 = &[]string{"GET"}
	var rt33 ReqRoute
	rt33.active = true
	rt33.path = "/test/test1/test2"
	rt33.pathVarsUsed = true
	rt33.pathVarNames = &[]string{"param1"}
	//rt.handler = hdl
	rt33.methods = &[]string{"GET"}
	//var rts3 = make(map[string]*[]Route)
	var prt3 []Route
	prt3 = append(prt3, &rt3)
	prt3 = append(prt3, &rt33)
	rts[rt33.path] = &prt3

	var rt4 ReqRoute
	rt4.active = true
	rt4.path = "/"
	rt4.pathVarsUsed = false
	rt4.pathVarNames = &[]string{}
	//rt.handler = hdl
	rt4.methods = &[]string{"GET"}

	//var varNames33 = &[]string{"param1"}
	//var methods33 = &[]string{"GET"}
	var rt41 ReqRoute
	rt41.active = true
	rt41.path = "/"
	rt41.pathVarsUsed = true
	rt41.pathVarNames = &[]string{"name"}
	//rt.handler = hdl
	rt41.methods = &[]string{"GET"}
	//var rts3 = make(map[string]*[]Route)
	var prt4 []Route
	prt4 = append(prt4, &rt4)
	prt4 = append(prt4, &rt41)
	rts[rt4.path] = &prt4

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
				pathVarNames: &[]string{"param1", "param2"},

				methods: &[]string{"GET"},
			},
			want1: &[]string{"p1", "p2"},
		},
		{
			name: "test 2",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				path: "/test/test1/test2/test3",
			},
			want: &ReqRoute{
				path:         "/test/test1/test2/test3",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},
				methods:      &[]string{"GET"},
			},
			want1: &[]string{},
		},
		{
			name: "test 3",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				path: "/test/test1/test2/p1",
			},
			want: &ReqRoute{
				path:         "/test/test1/test2",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"param1"},

				methods: &[]string{"GET"},
			},
			want1: &[]string{"p1"},
		},
		{
			name: "test 4",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				path: "/",
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: false,
				pathVarNames: &[]string{},

				methods: &[]string{"GET"},
			},
			want1: &[]string{},
		},
		{
			name: "test 5",
			fields: fields{
				namedRoutes: rts,
			},
			args: args{
				path: "/info",
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"name"},

				methods: &[]string{"GET"},
			},
			want1: &[]string{"info"},
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
			// if !reflect.DeepEqual(got1, tt.want1) {
			if len(*got1) != len(*tt.want1) {
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
		namedRoutes          map[string]*[]Route
		prefixRoutes         map[string]Route
		corsEnabled          bool
		corsAllowCredentials bool
		allowedHeaders       []string
		allowedOrigins       []string
		allowedMethods       []string
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
				allowedHeaders:       []string{"Content-Type"},
				allowedOrigins:       []string{"test"},
				allowedMethods:       []string{"POST", "GET"},
				corsAllowCredentials: true,
			},
			args: args{
				w: tw,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes:          tt.fields.namedRoutes,
				prefixRoutes:         tt.fields.prefixRoutes,
				corsEnabled:          tt.fields.corsEnabled,
				corsAllowCredentials: tt.fields.corsAllowCredentials,
				allowedHeaders:       tt.fields.allowedHeaders,
				allowedOrigins:       tt.fields.allowedOrigins,
				allowedMethods:       tt.fields.allowedMethods,
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
			if tw.Header().Get("Access-Control-Allow-Credentials") != "true" {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_CORSAllowCredentials(t *testing.T) {
	type fields struct {
		namedRoutes          map[string]*[]Route
		prefixRoutes         map[string]Route
		corsEnabled          bool
		corsAllowCredentials bool
		allowedHeaders       []string
		allowedOrigins       []string
		allowedMethods       []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				corsAllowCredentials: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRouter{
				namedRoutes:          tt.fields.namedRoutes,
				prefixRoutes:         tt.fields.prefixRoutes,
				corsEnabled:          tt.fields.corsEnabled,
				corsAllowCredentials: tt.fields.corsAllowCredentials,
				allowedHeaders:       tt.fields.allowedHeaders,
				allowedOrigins:       tt.fields.allowedOrigins,
				allowedMethods:       tt.fields.allowedMethods,
			}
			tr.CORSAllowCredentials()
			if tr.corsAllowCredentials != true {
				t.Fail()
			}
		})
	}
}

func TestReqRouter_isStaticFile(t *testing.T) {
	type fields struct {
		namedRoutes          map[string]*[]Route
		prefixRoutes         map[string]Route
		corsEnabled          bool
		corsAllowCredentials bool
		allowedHeaders       []string
		allowedOrigins       []string
		allowedMethods       []string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				path: "/templates/cleanblog/vendor/fontawesome-free/css/all.min.css",
			},
			want: true,
		},
		{
			name: "test 1",
			args: args{
				path: "/templates/cleanblog/js/clean-blog.min.js",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRouter{
				namedRoutes:          tt.fields.namedRoutes,
				prefixRoutes:         tt.fields.prefixRoutes,
				corsEnabled:          tt.fields.corsEnabled,
				corsAllowCredentials: tt.fields.corsAllowCredentials,
				allowedHeaders:       tt.fields.allowedHeaders,
				allowedOrigins:       tt.fields.allowedOrigins,
				allowedMethods:       tt.fields.allowedMethods,
			}
			if got := tr.isStaticFile(tt.args.path); got != tt.want {
				t.Errorf("ReqRouter.isStaticFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRouter_findFilePrefix(t *testing.T) {
	var prt = ReqRoute{
		prefix:   "/",
		isPrefix: true,
		active:   true,
	}
	pfx := make(map[string]Route)
	pfx["/"] = &prt
	type fields struct {
		namedRoutes          map[string]*[]Route
		prefixRoutes         map[string]Route
		corsEnabled          bool
		corsAllowCredentials bool
		allowedHeaders       []string
		allowedOrigins       []string
		allowedMethods       []string
	}
	type args struct {
		path string
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
				prefixRoutes: pfx,
			},
			args: args{
				path: "/templates/cleanblog/vendor/fontawesome-free/css/all.min.css",
			},
			want: &ReqRoute{
				prefix:   "/",
				isPrefix: true,
				active:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRouter{
				namedRoutes:          tt.fields.namedRoutes,
				prefixRoutes:         tt.fields.prefixRoutes,
				corsEnabled:          tt.fields.corsEnabled,
				corsAllowCredentials: tt.fields.corsAllowCredentials,
				allowedHeaders:       tt.fields.allowedHeaders,
				allowedOrigins:       tt.fields.allowedOrigins,
				allowedMethods:       tt.fields.allowedMethods,
			}
			if got := tr.findFilePrefix(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRouter.findFilePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
