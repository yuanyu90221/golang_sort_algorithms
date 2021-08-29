package insertion

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		A *[]int
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		{
			name: "[5,2,4,6,1,3]",
			args: args{
				A: &[]int{5, 2, 4, 6, 1, 3},
			},
			want: &[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
