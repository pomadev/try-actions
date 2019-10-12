package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{
			arg1: 1,
			arg2: 2,
			want: 3,
		},
		{
			arg1: 2,
			arg2: -2,
			want: 0,
		},
		{
			arg1: 1,
			arg2: 1,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Add(tt.arg1, tt.arg2)
			if got != tt.want {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
