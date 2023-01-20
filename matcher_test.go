package grrt

import (
	//"reflect"
	"testing"
)

func Test_pathMatcher_addPath(t *testing.T) {
	type fields struct {
		paths map[string]bool
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
			name: "addPathAlreadyExist",
			fields: fields{
				paths: map[string]bool{
					"/test/test1": true,
				},
			},
			args: args{
				p: "/test/test1",
			},
			want: false,
		},
		{
			name: "addPathOk",
			fields: fields{
				paths: map[string]bool{},
			},
			args: args{
				p: "/test/test1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &pathMatcher{
				paths: tt.fields.paths,
			}
			if got := m.addPath(tt.args.p); got != tt.want {
				t.Errorf("pathMatcher.addPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathMatcher_New(t *testing.T) {
	// type fields struct {
	// 	paths map[string]bool
	// }
	tests := []struct {
		name string
		//fields fields
		//want   Matcher
		want string
	}{
		// TODO: Add test cases.
		{
			name: "new",
			want: "needs to be new object",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// m := &pathMatcher{
			// 	//paths: tt.fields.paths,
			// }
			var m pathMatcher

			if got := m.New(); got == nil {
				t.Errorf("pathMatcher.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
