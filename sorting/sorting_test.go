package sorting

import (
	"reflect"
	"sort"
	"testing"

	algorithms "github.com/irenicaa/go-algorithms"
)

type Int int

func (number Int) Less(other interface{}) bool {
	return number < other.(Int)
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		items []algorithms.Less
	}

	tests := []struct {
		name string
		args args
		want []algorithms.Less
	}{
		{
			name: "empty",
			args: args{items: []algorithms.Less{}},
			want: []algorithms.Less{},
		},
		{
			name: "random",
			args: args{items: []algorithms.Less{
				Int(19732),
				Int(13),
				Int(4197),
				Int(4197),
				Int(23073),
				Int(23711),
				Int(14740),
				Int(14740),
				Int(22248),
				Int(6874),
				Int(1608),
				Int(6601),
			}},
			want: []algorithms.Less{
				Int(19732),
				Int(13),
				Int(4197),
				Int(4197),
				Int(23073),
				Int(23711),
				Int(14740),
				Int(14740),
				Int(22248),
				Int(6874),
				Int(1608),
				Int(6601),
			},
		},
		{
			name: "sorted ascending",
			args: args{items: []algorithms.Less{
				Int(13),
				Int(1608),
				Int(4197),
				Int(4197),
				Int(6601),
				Int(6874),
				Int(14740),
				Int(14740),
				Int(19732),
				Int(22248),
				Int(23073),
				Int(23711),
			}},
			want: []algorithms.Less{
				Int(19732),
				Int(13),
				Int(4197),
				Int(4197),
				Int(23073),
				Int(23711),
				Int(14740),
				Int(14740),
				Int(22248),
				Int(6874),
				Int(1608),
				Int(6601),
			},
		},
		{
			name: "sorted descending",
			args: args{items: []algorithms.Less{
				Int(6874),
				Int(6601),
				Int(4197),
				Int(4197),
				Int(23711),
				Int(23073),
				Int(22248),
				Int(19732),
				Int(1608),
				Int(14740),
				Int(14740),
				Int(13),
			}},
			want: []algorithms.Less{
				Int(19732),
				Int(13),
				Int(4197),
				Int(4197),
				Int(23073),
				Int(23711),
				Int(14740),
				Int(14740),
				Int(22248),
				Int(6874),
				Int(1608),
				Int(6601),
			},
		},
		{
			name: "identical",
			args: args{items: []algorithms.Less{
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
			}},
			want: []algorithms.Less{
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
				Int(6874),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.want, func(i int, j int) bool {
				return tt.want[i].(Int) < tt.want[j].(Int)
			})

			BubbleSort(tt.args.items)
			if !reflect.DeepEqual(tt.args.items, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.args.items, tt.want)
			}
		})
	}
}
