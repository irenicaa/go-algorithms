package sorting

import (
	"reflect"
	"testing"

	algorithms "github.com/irenicaa/go-algorithms"
)

type Int int

func (number Int) Less(other interface{}) bool {
	return number < other.(Int)
}

func TestSorting(t *testing.T) {
	type args struct {
		items []algorithms.Less
	}

	functions := []struct {
		name    string
		handler func(items []algorithms.Less)
	}{
		{
			name:    "BubleSort",
			handler: BubbleSort,
		},
		{
			name:    "CocktailSort",
			handler: CocktailSort,
		},
		{
			name:    "CombSort",
			handler: CombSort,
		},
		{
			name:    "InsertionSort",
			handler: InsertionSort,
		},
		{
			name:    "ShellSort",
			handler: ShellSort,
		},
		{
			name:    "SelectionSort",
			handler: SelectionSort,
		},
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
			args: args{
				items: []algorithms.Less{
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
			want: []algorithms.Less{
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
			},
		},
		{
			name: "sorted ascending",
			args: args{
				items: []algorithms.Less{
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
				},
			},
			want: []algorithms.Less{
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
			},
		},
		{
			name: "sorted descending",
			args: args{
				items: []algorithms.Less{
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
				},
			},
			want: []algorithms.Less{
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
			},
		},
		{
			name: "identical",
			args: args{
				items: []algorithms.Less{
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

	for _, function := range functions {
		for _, tt := range tests {
			t.Run(function.name+"/"+tt.name, func(t *testing.T) {
				items := []algorithms.Less{}
				for _, item := range tt.args.items {
					items = append(items, item)
				}

				function.handler(items)

				if !reflect.DeepEqual(items, tt.want) {
					t.Errorf("%s = %v, want %v", function.name, items, tt.want)
				}
			})
		}
	}
}
