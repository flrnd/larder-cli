package main

import "testing"

func TestJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"join(hello, world)", args{[]string{"hello", "world"}}, "helloworld"},
		{"join(こんにちは, 世界)", args{[]string{"こんにちは", "世界"}}, "こんにちは世界"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := join(tt.args.strs...); got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}
