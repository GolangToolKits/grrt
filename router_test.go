package grrt

import (
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	var rn = make(map[string]Route)
	tests := []struct {
		name string
		want Router
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: &ReqRouter{
				namedRoutes: rn,
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
