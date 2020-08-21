package algorithms

import (
	"reflect"
	"testing"
)

type IntLess int

func (number IntLess) Less(other interface{}) bool {
	return number < other.(IntLess)
}

func (number IntLess) Equal(other interface{}) bool {
	return number == other.(IntLess)
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

func TestSearch(t *testing.T) {
	type args struct {
		items  []Equal
		sample Equal
	}

	tests := []struct {
		name       string
		args       args
		wantResult Equal
		wantOk     bool
	}{
		{
			name: "found",
			args: args{
				items: []Equal{
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
				sample: IntLess(13),
			},
			wantResult: IntLess(13),
			wantOk:     true,
		},
		{
			name: "not found",
			args: args{
				items: []Equal{
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
				sample: IntLess(131),
			},
			wantResult: nil,
			wantOk:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotOk := Search(tt.args.items, tt.args.sample)

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Search() gotResult = %v, wantResult %v", gotResult, tt.wantResult)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Search() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}
