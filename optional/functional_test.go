package optional

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T any, V any] struct {
		o Option[T]
		f func(T) V
	}
	type testCase[T any, V any] struct {
		name string
		args args[T, V]
		want Option[V]
	}
	tests := []testCase[int, string]{
		{
			name: "int to string",
			args: args[int, string]{
				o: Some(1),
				f: strconv.Itoa,
			},
			want: Some("1"),
		},
		{
			name: "none to none",
			args: args[int, string]{
				o: None[int](),
				f: strconv.Itoa,
			},
			want: None[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
