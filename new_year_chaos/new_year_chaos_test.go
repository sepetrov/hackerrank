package new_year_chaos_test

import (
	"errors"
	"testing"

	"github.com/sepetrov/hackerrank/new_year_chaos"
)

func TestRun(t *testing.T) {
	errTooChaotic := errors.New("Too chaotic")
	tests := []struct {
		arg  []int32
		want int32
		err  error
	}{
		{
			[]int32{2, 1, 5, 3, 4},
			3,
			nil,
		},
		{
			[]int32{2, 5, 1, 3, 4},
			0,
			errTooChaotic,
		},
		{
			[]int32{5, 1, 2, 3, 7, 8, 6, 4},
			0,
			errTooChaotic,
		},
		{
			[]int32{1, 2, 5, 3, 7, 8, 6, 4},
			7,
			nil,
		},
		{
			[]int32{1, 2, 5, 3, 4, 7, 8, 6},
			4,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			arg := make([]int32, len(tt.arg))
			copy(arg, tt.arg)
			got, err := new_year_chaos.Run(arg)
			if got != tt.want || (err == nil) != (tt.err == nil) {
				t.Fatalf("Run(%v) = %v, %v; want %v, error %v", tt.arg, got, err, tt.want, tt.err)
			}
		})
	}
}
