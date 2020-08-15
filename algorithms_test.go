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
		{
			name: "success",
			args: args{
				items: []Less{
					IntLess(19732),
					IntLess(13),
					IntLess(4197),
					IntLess(23711),
					IntLess(23073),
					IntLess(14740),
					IntLess(22248),
					IntLess(6874),
					IntLess(6601),
					IntLess(1608),
				},
			},
			want: IntLess(13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minimum(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	type args struct {
		items []Less
	}

	tests := []struct {
		name string
		args args
		want Less
	}{
		{
			name: "success",
			args: args{
				items: []Less{
					IntLess(19732),
					IntLess(13),
					IntLess(4197),
					IntLess(23711),
					IntLess(23073),
					IntLess(14740),
					IntLess(22248),
					IntLess(6874),
					IntLess(6601),
					IntLess(1608),
				},
			},
			want: IntLess(23711),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maximum(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
