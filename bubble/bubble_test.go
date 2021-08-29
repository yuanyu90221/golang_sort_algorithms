package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		A *[]int
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
			},
			want: &[]int{1, 2, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
