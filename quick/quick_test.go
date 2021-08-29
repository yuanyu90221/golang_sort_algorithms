package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
			name: "[2,8,7,1,3,5,6,4]",
			args: args{
				A: &[]int{2, 8, 7, 1, 3, 5, 6, 4},
				p: 0,
				r: 7,
			},
			want: &[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.A, tt.args.p, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
