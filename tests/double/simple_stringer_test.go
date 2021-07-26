package double

import "testing"

func TestExample_StringVal(t *testing.T) {
	tests := []struct {
		i    Example
		want string
	}{
		{
			i:    Example_exampleBar0,
			want: "exampleBar0",
		},
		{
			i:    Example_exampleBar2,
			want: "exampleBar2",
		},
		{
			i:    Example_exampleBarEmpty,
			want: "Example_exampleBarEmpty",
		},
		{
			i:    Example_exampleBar4,
			want: "exampleBar4",
		},
		{
			i:    Example_exampleBar5,
			want: "exampleBar5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.i.StringVal(); got != tt.want {
				t.Errorf("StringVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample3_StringVal(t *testing.T) {
	tests := []struct {
		i    Example3
		want string
	}{
		{
			i:    Example3_example3Bar0,
			want: "exampleBar0",
		},
		{
			i:    Example3_example3Bar2,
			want: "Example3_example3Bar2",
		},
		{
			i:    Example3_example3Bar8,
			want: "Example3_example3Bar8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.i.StringVal(); got != tt.want {
				t.Errorf("StringVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
