package grrt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReqRouter_NewRoute(t *testing.T) {
	var m pathMatcher
	matcher := m.New()
	var vs = &[]string{}
	var mt = &[]string{}
	type fields struct {
		namedRoutes map[string]Route
	}
	tests := []struct {
		name   string
		fields fields
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			// fields: fields{

			// },
			want: &ReqRoute{
				matcher:      matcher,
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
	var nr = make(map[string]Route)
	var nf func(http.ResponseWriter, *http.Request)
	var m pathMatcher
	matcher := m.New()
	type fields struct {
		namedRoutes map[string]Route
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
				matcher:      matcher,
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
	var nr = make(map[string]Route)
	var nf func(http.ResponseWriter, *http.Request)
	var m pathMatcher
	matcher := m.New()
	type fields struct {
		namedRoutes map[string]Route
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
				matcher:      matcher,
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
	var rts = make(map[string]Route)

	tw := httptest.NewRecorder()
	tw2 := httptest.NewRecorder()
	tw3 := httptest.NewRecorder()

	tr, _ := http.NewRequest("POST", "/test/test1", nil)
	tr2, _ := http.NewRequest("POST", "/test/te", nil)
	tr3, _ := http.NewRequest("PUT", "/test/test1", nil)
	tr4, _ := http.NewRequest("POST", "/test/test1/param1/param2", nil)
	rts[rt.path] = &rt

	type fields struct {
		namedRoutes map[string]Route
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
				w: tw3,
				r: tr4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := ReqRouter{
				namedRoutes: tt.fields.namedRoutes,
			}
			tr.ServeHTTP(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 2 404" && tw2.Code != http.StatusNotFound {
				t.Fail()
			}
			if tt.name == "test 2 405" && tw3.Code != http.StatusMethodNotAllowed {
				t.Fail()
			}
		})
	}
}
