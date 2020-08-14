package algorithms

import (
	"reflect"
	"testing"
)

type IntLess int

func (number IntLess) Less(other interface{}) bool {
	return number < other.(IntLess)
}

func TestMinimum(t *testing.T) {
	type args struct {
		items []Less
	}

	tests := []struct {
		name string
		args args
		want Less
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minimum(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
