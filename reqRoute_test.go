package grrt

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestReqRoute_chechPath(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "url match success",
			args: args{
				p: "/test/route/{parm1}/{parm2}",
			},
			want: true,
		},
		{
			name: "url match fail",
			args: args{
				p: "/test/route//{parm1}/{parm2}",
			},
			want: false,
		},
		{
			name: "url match success",
			args: args{
				p: "/",
			},
			want: true,
		},
		{
			name: "url match success",
			args: args{
				p: "/{parm1}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
			}
			if got := tr.chechPath(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_chechCurlys(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "url match success",
			args: args{
				p: "/test/route/{parm1}/{parm2}",
			},
			want: true,
		},
		{
			name: "url match fail",
			args: args{
				p: "/test/route/{{parm1}/{parm2}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
			}
			if got := tr.chechCurlys(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechCurlys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_chechBackSlash(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "start / missing test 1",
			args: args{
				p: "route/test/{parm1}",
			},
			want: false,
		},
		{
			name: "double // test 2",
			args: args{
				p: "//route/test/{parm1}",
			},
			want: false,
		},
		{
			name: "no double // test 3",
			args: args{
				p: "/route/test/{parm1}",
			},
			want: true,
		},
		{
			name: "no double // test 4",
			args: args{
				p: "/route//test/{parm1}",
			},
			want: false,
		},
		{
			name: "no double // test 5",
			args: args{
				p: "/route/test//{parm1}",
			},
			want: false,
		},
		{
			name: "no double // test 6",
			args: args{
				p: "/route/test/{parm1}/",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
			}
			if got := tr.chechBackSlash(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechBackSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_chechCurlyPlacement(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.

		{
			name: "test 1 ok",
			args: args{
				p: "/route/test/{parm1}",
			},
			want: true,
		},
		{
			name: "test 2 fail",
			args: args{
				p: "/route/test/{{parm1}",
			},
			want: false,
		},
		{
			name: "test 3 fail",
			args: args{
				p: "/route/test/{parm1}}",
			},
			want: false,
		},
		{
			name: "test 3 fail",
			args: args{
				p: "/{parm1}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
			}
			if got := tr.chechCurlyPlacement(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechCurlyPlacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_extractPathAndVarNames(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]string
		want1  string
	}{
		// TODO: Add test cases.
		{
			name: "test 1 one var success",
			args: args{
				p: "/route/test/{param1}",
			},
			want:  &[]string{"param1"},
			want1: "/route/test",
		},
		{
			name: "test 2 two var success",
			args: args{
				p: "/route/test/{param1}/{param2}",
			},
			want:  &[]string{"param1", "param2"},
			want1: "/route/test",
		},
		{
			name: "test 3 four var success",
			args: args{
				p: "/route/test/{param1}/{param2}/{param3}/{param4}",
			},
			want:  &[]string{"param1", "param2", "param3", "param4"},
			want1: "/route/test",
		},
		{
			name: "test 4 four var success",
			args: args{
				p: "/{ss}",
			},
			want:  &[]string{"ss"},
			want1: "/",
		},
		{
			name: "test 5 no var success",
			args: args{
				p: "/route/test",
			},
			want:  &[]string{},
			want1: "/route/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			got, got1 := tr.extractPathAndVarNames(tt.args.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.extractPathAndVarNames() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReqRoute.extractPathAndVarNames() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReqRoute_Path(t *testing.T) {

	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
	}
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1 ",
			args: args{
				p: "/route/test/{param1}",
			},
			fields: fields{
				pathVarNames: &[]string{},
			},
			want: &ReqRoute{
				path:         "/route/test",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"param1"},
			},
		},
		{
			name: "test 2 ",
			args: args{
				p: "/route/test",
			},
			fields: fields{
				pathVarNames: &[]string{},
			},
			want: &ReqRoute{
				path:         "/route/test",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{},
			},
		},
		{
			name: "test 3 ",
			args: args{
				p: "/route/test/{{param1}",
			},
			fields: fields{
				pathVarNames: &[]string{},
			},
			want: &ReqRoute{
				path:         "",
				active:       false,
				pathVarsUsed: false,
				pathVarNames: &[]string{},
			},
		},
		{
			name: "test 4 ",
			args: args{
				p: "/",
			},
			fields: fields{
				pathVarNames: &[]string{},
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{},
			},
		},
		{
			name: "test 5 ",
			args: args{
				p: "/{param1}",
			},
			fields: fields{
				pathVarNames: &[]string{},
			},
			want: &ReqRoute{
				path:         "/",
				active:       true,
				pathVarsUsed: true,
				pathVarNames: &[]string{"param1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			// if got := tr.Path(tt.args.p); !reflect.DeepEqual(got, tt.want) {
			if got := tr.Path(tt.args.p); got.GetPath() != tt.want.GetPath() ||
				len(*got.GetVarNames()) != len(*tt.want.GetVarNames()) ||
				got.IsActive() != tt.want.IsActive() {
				fmt.Println("got: ", got)
				t.Errorf("ReqRoute.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_Handler(t *testing.T) {
	var hd http.Handler
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
	}
	type args struct {
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
			args: args{
				handler: hd,
			},
			fields: fields{
				handler: hd,
				active:  true,
			},
			want: &ReqRoute{
				path:    "",
				active:  true,
				handler: hd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			if got := tr.Handler(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_HandlerFunc(t *testing.T) {
	var nf func(http.ResponseWriter, *http.Request)
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
	}
	type args struct {
		f func(http.ResponseWriter, *http.Request)
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
			args: args{
				f: nf,
			},
			fields: fields{
				handler: http.HandlerFunc(nf),
				active:  true,
			},
			want: &ReqRoute{
				path:    "",
				active:  true,
				handler: http.HandlerFunc(nf),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			if got := tr.HandlerFunc(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.HandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_GetHandler(t *testing.T) {
	var nf func(http.ResponseWriter, *http.Request)
	var hd = http.HandlerFunc(nf)
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Handler
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				handler: hd,
				active:  true,
			},
			want: hd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			if got := tr.GetHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.GetHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_New(t *testing.T) {
	var vs = &[]string{}
	var mt = &[]string{}
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   Route
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				pathVarNames: vs,
				methods:      mt,
			},
			want: &ReqRoute{
				pathVarNames: vs,
				methods:      mt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			if got := tr.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_Methods(t *testing.T) {
	var mt = &[]string{"POST", "PUT"}
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	type args struct {
		ms []string
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
				methods: &[]string{},
			},
			args: args{
				ms: []string{"post", "put"},
			},
			want: &ReqRoute{
				methods: mt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.Methods(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.Methods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_GetMethods(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				methods: &[]string{"POST", "PUT"},
			},
			want: &[]string{"POST", "PUT"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.GetMethods(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.GetMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_IsMethodAllowed(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	type args struct {
		m string
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
			fields: fields{
				methods: &[]string{"POST", "GET"},
			},
			args: args{
				m: "POST",
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				methods: &[]string{},
			},
			args: args{
				m: "POST",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.IsMethodAllowed(tt.args.m); got != tt.want {
				t.Errorf("ReqRoute.IsMethodAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_IsPathVarsUsed(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				pathVarsUsed: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.IsPathVarsUsed(); got != tt.want {
				t.Errorf("ReqRoute.IsPathVarsUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_chechPrefix(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		prefix       string
		active       bool
		pathVarsUsed bool
		pathVarNames *[]string
		methods      *[]string
	}
	type args struct {
		p string
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
				p: "/testprefix",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				p: "testprefix",
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				p: "/testprefix/",
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				p: "//testprefix/",
			},
			want: false,
		},
		{
			name: "test 5",
			args: args{
				p: "/testprefix/pre",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				prefix:       tt.fields.prefix,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.chechPrefix(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_PathPrefix(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		prefix       string
		active       bool
		pathVarsUsed bool
		isPrefix     bool
		pathVarNames *[]string
		methods      *[]string
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
				active:   true,
				isPrefix: true,
				prefix:   "/testPrefix",
			},
			args: args{
				px: "/testPrefix",
			},
			want: &ReqRoute{
				active:   true,
				isPrefix: true,
				prefix:   "/testPrefix",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				prefix:       tt.fields.prefix,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				isPrefix:     tt.fields.isPrefix,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.PathPrefix(tt.args.px); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.PathPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_GetPrefix(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		prefix       string
		active       bool
		pathVarsUsed bool
		isPrefix     bool
		pathVarNames *[]string
		methods      *[]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				prefix: "/testPrefix/",
			},
			want: "/testPrefix/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				prefix:       tt.fields.prefix,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				isPrefix:     tt.fields.isPrefix,
				pathVarNames: tt.fields.pathVarNames,
				methods:      tt.fields.methods,
			}
			if got := tr.GetPrefix(); got != tt.want {
				t.Errorf("ReqRoute.GetPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
