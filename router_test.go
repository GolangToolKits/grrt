package grrt

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	var rn = make(map[string]*[]Route)
	var px = make(map[string]Route)
	tests := []struct {
		name string
		want Router
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: &ReqRouter{
				namedRoutes: rn,
				prefixRoutes: px,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVars(t *testing.T) {
	tr1, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)

	m1 := make(map[string]string)

	ctx := context.WithValue(tr1.Context(), varsKey, m1)
	tr1 = tr1.WithContext(ctx)
	m1["param1"] = "p1"
	m1["param2"] = "p2"

	tr2, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	m2 := make(map[string]string)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				r: tr1,
			},
			want: m1,
		},
		{
			name: "test 2",
			args: args{
				r: tr2,
			},
			want: m2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vars(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetURLVars(t *testing.T) {
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	vars := make(map[string]string)
	vars["param1"] = "test1"
	vars["param2"] = "test2"
	type args struct {
		r    *http.Request
		vars map[string]string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				r:    tr,
				vars: vars,
			},
			want: tr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := SetURLVars(tt.args.r, tt.args.vars)
			fvars := Vars(got)
			if fvars["param1"] != "test1" || fvars["param2"] != "test2" {
				t.Fail()
			}
		})
	}
}
