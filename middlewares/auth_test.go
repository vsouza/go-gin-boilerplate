package middlewares

import "testing"

func Test_secureCompare(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Parameters a and b are equal",
			args: args{
				a: "abc123",
				b: "abc123",
			},
			want: 1,
		},
		{
			name: "Parameters a and b are not equal",
			args: args{
				a: "123abc",
				b: "abc123",
			},
			want: 0,
		},
		{
			name: "Parameters a and b are almost equal",
			args: args{
				a: "abc123",
				b: "abd123",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secureCompare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("secureCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
