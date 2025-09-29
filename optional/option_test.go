package optional

import (
	"reflect"
	"testing"
)

func toP[T any](v T) *T {
	return &v
}

func TestFromPtr(t *testing.T) {
	type args[T any] struct {
		v *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "pointer value to some",
			args: args[int]{
				v: toP(1),
			},
			want: Some(1),
		},
		{
			name: "nil to none",
			args: args[int]{
				v: nil,
			},
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromPtr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_IsSone(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Some is true",
			o:    Some(1),
			want: true,
		},
		{
			name: "None is false",
			o:    None[int](),
			want: false,
		},
		{
			name: "Some initialized by FromPtr is true",
			o:    FromPtr(toP(1)),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsSone(); got != tt.want {
				t.Errorf("IsSone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_IsNone(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Some is false",
			o:    Some(1),
			want: false,
		},
		{
			name: "None is true",
			o:    None[int](),
			want: true,
		},
		{
			name: "Some initialized by FromPtr is false",
			o:    FromPtr(toP(1)),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsNone(); got != tt.want {
				t.Errorf("IsNone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Get(t *testing.T) {
	type testCase[T any] struct {
		name    string
		o       Option[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "get value from some",
			o:       Some(1),
			want:    1,
			wantErr: false,
		},
		{
			name:    "get error from none",
			o:       None[int](),
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetOr(t *testing.T) {
	type args[T any] struct {
		fallback T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "get value from some",
			o:    Some(1),
			args: args[int]{
				fallback: 9,
			},
			want: 1,
		},
		{
			name: "get fallback value from none",
			o:    None[int](),
			args: args[int]{
				fallback: 9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOr(tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_ToPtr(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want *T
	}
	tests := []testCase[int]{
		{
			name: "some to pointer value",
			o:    Some(1),
			want: toP(1),
		},
		{
			name: "none to nil",
			o:    None[int](),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.ToPtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
