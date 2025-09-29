package optional

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOption_UnmarshalJSON(t *testing.T) {
	type R struct {
		Val1 Option[int]     `json:"val1"`
		Val2 Option[string]  `json:"val2"`
		Val3 Option[float64] `json:"val3"`
		Val4 Option[int]     `json:"val4"`
	}
	type args struct {
		data []byte
	}
	type testCase struct {
		name    string
		args    args
		want    R
		wantErr bool
	}
	tests := []testCase{
		{
			name: "unmarshal json",
			args: args{
				data: []byte(`
{
	"val1": 1,
	"val2": "hello",
	"val3": 3.14,
	"val4": null
}
`),
			},
			want: R{
				Val1: Some(1),
				Val2: Some("hello"),
				Val3: Some(3.14),
				Val4: None[int](),
			},
			wantErr: false,
		},
		{
			name: "unmarshal omitted json",
			args: args{
				data: []byte(`
{
	"val1": 1,
	"val2": "hello",
	"val3": 3.14
}
`),
			},
			want: R{
				Val1: Some(1),
				Val2: Some("hello"),
				Val3: Some(3.14),
				Val4: None[int](),
			},
			wantErr: false,
		},
		{
			name: "unmarshal invalid type",
			args: args{
				data: []byte(`
{
	"val1": "1"
}`),
			},
			want:    R{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got R
			err := json.Unmarshal(tt.args.data, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmp.AllowUnexported(Option[int]{}, Option[string]{}, Option[float64]{}),
			}
			if diff := cmp.Diff(tt.want, got, opts); diff != "" {
				t.Errorf("UnmarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOption_MarshalJSON(t *testing.T) {
	type R struct {
		Val1 Option[int]     `json:"val1"`
		Val2 Option[string]  `json:"val2"`
		Val3 Option[float64] `json:"val3"`
		Val4 Option[int]     `json:"val4"`
	}
	type testCase struct {
		name    string
		src     R
		want    string
		wantErr bool
	}
	tests := []testCase{
		{
			name: "marshal json",
			src: R{
				Val1: Some(1),
				Val2: Some("hello"),
				Val3: Some(3.14),
				Val4: None[int](),
			},
			want: `
{
	"val1": 1,
	"val2": "hello",
	"val3": 3.14,
	"val4": null
}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Remove all whitespace for comparison
			want := strings.ReplaceAll(tt.want, "\n", "")
			want = strings.ReplaceAll(want, "\t", "")
			want = strings.ReplaceAll(want, " ", "")
			if diff := cmp.Diff(want, string(got)); diff != "" {
				t.Errorf("MarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestOption_MarshalJSON_OmitZero(t *testing.T) {
	type R struct {
		Val1 Option[int] `json:"val1,omitzero"`
	}
	type testCase struct {
		name    string
		src     R
		want    string
		wantErr bool
	}
	tests := []testCase{
		{
			name: "marshal json",
			src: R{
				Val1: Some(1),
			},
			want:    `{"val1":1}`,
			wantErr: false,
		},
		{
			name: "marshal json with omitzero",
			src: R{
				Val1: None[int](),
			},
			want:    `{}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Errorf("MarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
