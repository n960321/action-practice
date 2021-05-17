package main

import "testing"

func Test_foo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first test",
			args: args{
				a: 3,
				b: 7,
			},
			want: 10,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
