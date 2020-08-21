package algorithms

import (
	"reflect"
	"testing"
)

type Int int

func (number Int) Less(other interface{}) bool {
	return number < other.(Int)
}

func (number Int) Equal(other interface{}) bool {
	return number == other.(Int)
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
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(23073),
					Int(14740),
					Int(22248),
					Int(6874),
					Int(6601),
					Int(1608),
				},
			},
			want: Int(13),
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
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(23073),
					Int(14740),
					Int(22248),
					Int(6874),
					Int(6601),
					Int(1608),
				},
			},
			want: Int(23711),
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

func TestSearchUniversal(t *testing.T) {
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
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(23073),
					Int(14740),
					Int(22248),
					Int(6874),
					Int(6601),
					Int(1608),
				},
				sample: Int(13),
			},
			wantResult: Int(13),
			wantOk:     true,
		},
		{
			name: "not found",
			args: args{
				items: []Equal{
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(23073),
					Int(14740),
					Int(22248),
					Int(6874),
					Int(6601),
					Int(1608),
				},
				sample: Int(131),
			},
			wantResult: nil,
			wantOk:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotOk := SearchUniversal(tt.args.items, tt.args.sample)

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Search() gotResult = %v, wantResult %v", gotResult, tt.wantResult)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Search() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}
