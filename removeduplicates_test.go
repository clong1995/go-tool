package tool

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type args[T comparable] struct {
		slice []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "TestRemoveDuplicates",
			args: args[int]{
				slice: []int{1, 2, 3, 3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.args.slice)
			t.Logf("RemoveDuplicates() = %v", got)
		})
	}
}
