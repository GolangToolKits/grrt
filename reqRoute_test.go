package grrt

import (
	"net/http"
	"reflect"
	"testing"
)

func TestReqRoute_chechPath(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		matcher      Matcher
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				matcher:      tt.fields.matcher,
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
		matcher      Matcher
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
				matcher:      tt.fields.matcher,
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
		matcher      Matcher
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
				matcher:      tt.fields.matcher,
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
		matcher      Matcher
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				matcher:      tt.fields.matcher,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
			}
			if got := tr.chechCurlyPlacement(tt.args.p); got != tt.want {
				t.Errorf("ReqRoute.chechCurlyPlacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqRoute_extractPathVarNames(t *testing.T) {
	type fields struct {
		handler      http.Handler
		host         string
		path         string
		matcher      Matcher
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
	}{
		// TODO: Add test cases.
		{
			name: "test 1 one var success",
			args: args{
				p: "/route/test/{param1}",
			},
			want: &[]string{"param1"},
		},
		{
			name: "test 2 two var success",
			args: args{
				p: "/route/test/{param1}/{param2}",
			},
			want: &[]string{"param1", "param2"},
		},
		{
			name: "test 3 four var success",
			args: args{
				p: "/route/test/{param1}/{param2}/{param3}/{param4}",
			},
			want: &[]string{"param1", "param2", "param3", "param4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &ReqRoute{
				handler:      tt.fields.handler,
				host:         tt.fields.host,
				path:         tt.fields.path,
				matcher:      tt.fields.matcher,
				active:       tt.fields.active,
				pathVarsUsed: tt.fields.pathVarsUsed,
				pathVarNames: tt.fields.pathVarNames,
			}
			if got := tr.extractPathVarNames(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReqRoute.extractPathVarNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
