package grrt

import (
	"net/http"
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
