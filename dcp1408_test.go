package dcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	type args struct {
		lst []int
		i   int
		j   int
	}
	tests := []struct {
		name   string
		args   args
		wanted []int
	}{
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   0,
				j:   4,
			},
			wanted: []int{5, 4, 3, 2, 1},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   0,
				j:   0,
			},
			wanted: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   3,
				j:   0,
			},
			wanted: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   4,
				j:   10,
			},
			wanted: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   2,
				j:   10,
			},
			wanted: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   2,
				j:   3,
			},
			wanted: []int{1, 2, 4, 3, 5},
		},
		{
			name: "reverse",
			args: args{
				lst: []int{1, 2, 3, 4, 5},
				i:   2,
				j:   4,
			},
			wanted: []int{1, 2, 5, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.args.lst, tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanted, result)
		})
	}
}
