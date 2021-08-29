package merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		A *[]int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		{
			name: "[2,4,5,7,1,2,3,6]",
			args: args{
				A: &[]int{2, 4, 5, 7, 1, 2, 3, 6},
				p: 0,
				r: 7,
			},
			want: &[]int{1, 2, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.A, tt.args.p, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
