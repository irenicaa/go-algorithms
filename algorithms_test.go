package algorithms

import (
	"reflect"
	"sort"
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

func TestSearchSorted(t *testing.T) {
	type args struct {
		items  []Less
		sample Less
	}

	tests := []struct {
		name       string
		args       args
		wantResult Less
		wantOk     bool
	}{
		{
			name: "empty",
			args: args{
				items:  []Less{},
				sample: Int(6874),
			},
			wantResult: nil,
			wantOk:     false,
		},
		{
			name: "found smallest",
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
				sample: Int(13),
			},
			wantResult: Int(13),
			wantOk:     true,
		},
		{
			name: "found in the middle",
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
				sample: Int(6874),
			},
			wantResult: Int(6874),
			wantOk:     true,
		},
		{
			name: "found largest",
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
				sample: Int(23711),
			},
			wantResult: Int(23711),
			wantOk:     true,
		},
		{
			name: "not found too small",
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
				sample: Int(12),
			},
			wantResult: nil,
			wantOk:     false,
		},
		{
			name: "not found in the middle",
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
				sample: Int(10000),
			},
			wantResult: nil,
			wantOk:     false,
		},
		{
			name: "not found too large",
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
				sample: Int(23712),
			},
			wantResult: nil,
			wantOk:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.args.items, func(i int, j int) bool {
				return tt.args.items[i].(Int) < tt.args.items[j].(Int)
			})

			gotResult, gotOk := SearchSorted(tt.args.items, tt.args.sample)

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Search() gotResult = %v, wantResult %v", gotResult, tt.wantResult)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Search() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestUniqueUniversal(t *testing.T) {
	type args struct {
		items []Equal
	}

	tests := []struct {
		name string
		args args
		want []Equal
	}{
		{
			name: "success",
			args: args{
				items: []Equal{
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(13),
					Int(14740),
					Int(22248),
					Int(13),
					Int(6601),
					Int(1608),
					Int(1608),
				},
			},
			want: []Equal{
				Int(19732),
				Int(13),
				Int(4197),
				Int(23711),
				Int(14740),
				Int(22248),
				Int(6601),
				Int(1608),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUniversal(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUniversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueSorted(t *testing.T) {
	type args struct {
		items []Equal
	}

	tests := []struct {
		name string
		args args
		want []Equal
	}{
		{
			name: "success",
			args: args{
				items: []Equal{
					Int(19732),
					Int(13),
					Int(4197),
					Int(23711),
					Int(13),
					Int(14740),
					Int(22248),
					Int(13),
					Int(6601),
					Int(1608),
					Int(1608),
				},
			},
			want: []Equal{
				Int(19732),
				Int(13),
				Int(4197),
				Int(23711),
				Int(14740),
				Int(22248),
				Int(6601),
				Int(1608),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.args.items, func(i int, j int) bool {
				return tt.args.items[i].(Int) < tt.args.items[j].(Int)
			})
			sort.Slice(tt.want, func(i int, j int) bool {
				return tt.want[i].(Int) < tt.want[j].(Int)
			})

			if got := UniqueSorted(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFast(t *testing.T) {
	type args struct {
		items []interface{}
	}

	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "success",
			args: args{
				items: []interface{}{
					19732,
					13,
					4197,
					23711,
					13,
					14740,
					22248,
					13,
					6601,
					1608,
					1608,
				},
			},
			want: []interface{}{
				19732,
				13,
				4197,
				23711,
				14740,
				22248,
				6601,
				1608,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.want, func(i int, j int) bool {
				return tt.want[i].(int) < tt.want[j].(int)
			})

			got := UniqueFast(tt.args.items)
			sort.Slice(got, func(i int, j int) bool {
				return got[i].(int) < got[j].(int)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSorted(t *testing.T) {
	type args struct {
		left  []Less
		right []Less
	}

	tests := []struct {
		name string
		args args
		want []Less
	}{
		{
			name: "empty",
			args: args{left: []Less{}, right: []Less{}},
			want: []Less{},
		},
		{
			name: "left is longer",
			args: args{
				left:  []Less{Int(1), Int(3), Int(5), Int(10), Int(11)},
				right: []Less{Int(2), Int(7), Int(8), Int(20)},
			},
			want: []Less{
				Int(1),
				Int(2),
				Int(3),
				Int(5),
				Int(7),
				Int(8),
				Int(10),
				Int(11),
				Int(20),
			},
		},
		{
			name: "right is longer",
			args: args{
				left:  []Less{Int(2), Int(7), Int(8), Int(20)},
				right: []Less{Int(1), Int(3), Int(5), Int(10), Int(11)},
			},
			want: []Less{
				Int(1),
				Int(2),
				Int(3),
				Int(5),
				Int(7),
				Int(8),
				Int(10),
				Int(11),
				Int(20),
			},
		},
		{
			name: "equal in length (without duplicates)",
			args: args{
				left:  []Less{Int(2), Int(7), Int(8), Int(12), Int(20)},
				right: []Less{Int(1), Int(3), Int(5), Int(10), Int(11)},
			},
			want: []Less{
				Int(1),
				Int(2),
				Int(3),
				Int(5),
				Int(7),
				Int(8),
				Int(10),
				Int(11),
				Int(12),
				Int(20),
			},
		},
		{
			name: "equal in length (with duplicates)",
			args: args{
				left:  []Less{Int(2), Int(5), Int(7), Int(8), Int(12), Int(20)},
				right: []Less{Int(1), Int(3), Int(5), Int(7), Int(10), Int(11)},
			},
			want: []Less{
				Int(1),
				Int(2),
				Int(3),
				Int(5),
				Int(5),
				Int(7),
				Int(7),
				Int(8),
				Int(10),
				Int(11),
				Int(12),
				Int(20),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSorted(tt.args.left, tt.args.right)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
