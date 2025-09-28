package optional

import (
	"testing"
)

func toP[T any](v T) *T {
	return &v
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
