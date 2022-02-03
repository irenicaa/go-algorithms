package datastructures

import (
	"reflect"
	"testing"
)

func TestNewMultiset(t *testing.T) {
	type args struct {
		items []interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantSet Multiset
	}{
		{
			name:    "empty",
			args:    args{items: []interface{}{}},
			wantSet: Multiset{},
		},
		{
			name:    "unique items",
			args:    args{items: []interface{}{"one", "two"}},
			wantSet: Multiset{"one": 1, "two": 1},
		},
		{
			name:    "duplicate items",
			args:    args{items: []interface{}{"one", "one"}},
			wantSet: Multiset{"one": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMultiset(tt.args.items...)

			if !reflect.DeepEqual(got, tt.wantSet) {
				t.Errorf("NewMultiset() = %v, want %v", got, tt.wantSet)
			}
		})
	}
}

func TestMultiset_Contains(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name string
		set  Multiset
		args args
		want bool
	}{
		{
			name: "nonexistent",
			set:  Multiset{"one": 2, "two": 3},
			args: args{item: "three"},
			want: false,
		},
		{
			name: "existent (with a nonzero value)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{item: "two"},
			want: true,
		},
		{
			name: "existent (with a zero value)",
			set:  Multiset{"one": 2, "two": 0},
			args: args{item: "two"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Contains(tt.args.item); got != tt.want {
				t.Errorf("Multiset.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Add(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name    string
		set     Multiset
		args    args
		wantSet Multiset
	}{
		{
			name:    "nonexistent",
			set:     Multiset{"one": 2, "two": 3},
			args:    args{item: "three"},
			wantSet: Multiset{"one": 2, "two": 3, "three": 1},
		},
		{
			name:    "existent",
			set:     Multiset{"one": 2, "two": 3},
			args:    args{item: "two"},
			wantSet: Multiset{"one": 2, "two": 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Add(tt.args.item)

			if !reflect.DeepEqual(tt.set, tt.wantSet) {
				t.Errorf("got %v, want %v", tt.set, tt.wantSet)
			}
		})
	}
}

func TestMultiset_Remove(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name    string
		set     Multiset
		args    args
		wantSet Multiset
	}{
		{
			name:    "nonexistent",
			set:     Multiset{"one": 2, "two": 3},
			args:    args{item: "three"},
			wantSet: Multiset{"one": 2, "two": 3},
		},
		{
			name:    "existent (more than one)",
			set:     Multiset{"one": 2, "two": 3},
			args:    args{item: "two"},
			wantSet: Multiset{"one": 2, "two": 2},
		},
		{
			name:    "existent (one)",
			set:     Multiset{"one": 2, "two": 1},
			args:    args{item: "two"},
			wantSet: Multiset{"one": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Remove(tt.args.item)

			if !reflect.DeepEqual(tt.set, tt.wantSet) {
				t.Errorf("got %v, want %v", tt.set, tt.wantSet)
			}
		})
	}
}

func TestMultiset_Sum(t *testing.T) {
	type args struct {
		other Multiset
	}

	tests := []struct {
		name string
		set  Multiset
		args args
		want Multiset
	}{
		{
			name: "both are empty",
			set:  Multiset{},
			args: args{other: Multiset{}},
			want: Multiset{},
		},
		{
			name: "first is empty",
			set:  Multiset{},
			args: args{other: Multiset{"one": 2, "two": 3}},
			want: Multiset{"one": 2, "two": 3},
		},
		{
			name: "second is empty",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{}},
			want: Multiset{"one": 2, "two": 3},
		},
		{
			name: "nobody is empty",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 4, "three": 2}},
			want: Multiset{"one": 2, "two": 7, "three": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Sum(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Union(t *testing.T) {
	type args struct {
		other Multiset
	}

	tests := []struct {
		name string
		set  Multiset
		args args
		want Multiset
	}{
		{
			name: "both are empty",
			set:  Multiset{},
			args: args{other: Multiset{}},
			want: Multiset{},
		},
		{
			name: "first is empty",
			set:  Multiset{},
			args: args{other: Multiset{"one": 2, "two": 3}},
			want: Multiset{"one": 2, "two": 3},
		},
		{
			name: "second is empty",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{}},
			want: Multiset{"one": 2, "two": 3},
		},
		{
			name: "nobody is empty",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 4, "three": 2}},
			want: Multiset{"one": 2, "two": 4, "three": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Intersection(t *testing.T) {
	type args struct {
		other Multiset
	}

	tests := []struct {
		name string
		set  Multiset
		args args
		want Multiset
	}{
		{
			name: "both are empty",
			set:  Multiset{},
			args: args{other: Multiset{}},
			want: Multiset{},
		},
		{
			name: "nobody is empty (with intersection, the left item is bigger)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 1, "three": 2}},
			want: Multiset{"two": 1},
		},
		{
			name: "nobody is empty (with intersection, the left item is smaller)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 4, "three": 2}},
			want: Multiset{"two": 3},
		},
		{
			name: "nobody is empty (without intersection)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"three": 4, "four": 2}},
			want: Multiset{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Intersection(tt.args.other)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiset_Difference(t *testing.T) {
	type args struct {
		other Multiset
	}

	tests := []struct {
		name string
		set  Multiset
		args args
		want Multiset
	}{
		{
			name: "both are empty",
			set:  Multiset{},
			args: args{other: Multiset{}},
			want: Multiset{},
		},
		{
			name: "nobody is empty (with intersection, without removing)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 1, "three": 2}},
			want: Multiset{"one": 2, "two": 2},
		},
		{
			name: "nobody is empty (with intersection, with removing)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"two": 4, "three": 2}},
			want: Multiset{"one": 2},
		},
		{
			name: "nobody is empty (without intersection)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"three": 4, "four": 2}},
			want: Multiset{"one": 2, "two": 3},
		},
		{
			name: "nobody is empty (same)",
			set:  Multiset{"one": 2, "two": 3},
			args: args{other: Multiset{"one": 2, "two": 3}},
			want: Multiset{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Difference(tt.args.other)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiset.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
