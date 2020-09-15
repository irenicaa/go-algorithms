package datastructures

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	type args struct {
		items []interface{}
	}

	tests := []struct {
		name string
		args args
		want Set
	}{
		{
			name: "empty",
			args: args{items: []interface{}{}},
			want: Set{},
		},
		{
			name: "unique items",
			args: args{items: []interface{}{"one", "two"}},
			want: Set{"one": {}, "two": {}},
		},
		{
			name: "duplicate items",
			args: args{items: []interface{}{"one", "one"}},
			want: Set{"one": {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name string
		set  Set
		args args
		want bool
	}{
		{
			name: "nonexistent",
			set:  Set{"one": {}, "two": {}},
			args: args{item: "three"},
			want: false,
		},
		{
			name: "existent",
			set:  Set{"one": {}, "two": {}},
			args: args{item: "two"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Contains(tt.args.item)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name    string
		set     Set
		args    args
		wantSet Set
	}{
		{
			name:    "nonexistent",
			set:     Set{"one": {}, "two": {}},
			args:    args{item: "three"},
			wantSet: Set{"one": {}, "two": {}, "three": {}},
		},
		{
			name:    "existent",
			set:     Set{"one": {}, "two": {}},
			args:    args{item: "two"},
			wantSet: Set{"one": {}, "two": {}},
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

func TestSet_Remove(t *testing.T) {
	type args struct {
		item interface{}
	}

	tests := []struct {
		name    string
		set     Set
		args    args
		wantSet Set
	}{
		{
			name:    "nonexistent",
			set:     Set{"one": {}, "two": {}},
			args:    args{item: "three"},
			wantSet: Set{"one": {}, "two": {}},
		},
		{
			name:    "existent",
			set:     Set{"one": {}, "two": {}},
			args:    args{item: "two"},
			wantSet: Set{"one": {}},
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

func TestSet_Union(t *testing.T) {
	type args struct {
		other Set
	}
	tests := []struct {
		name string
		set  Set
		args args
		want Set
	}{
		{
			name: "both are empty",
			set:  Set{},
			args: args{other: Set{}},
			want: Set{},
		},
		{
			name: "first is empty",
			set:  Set{},
			args: args{other: Set{"one": {}, "two": {}}},
			want: Set{"one": {}, "two": {}},
		},
		{
			name: "second is empty",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{}},
			want: Set{"one": {}, "two": {}},
		},
		{
			name: "nobody is empty",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"two": {}, "three": {}}},
			want: Set{"one": {}, "two": {}, "three": {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		other Set
	}
	tests := []struct {
		name string
		set  Set
		args args
		want Set
	}{
		{
			name: "both are empty",
			set:  Set{},
			args: args{other: Set{}},
			want: Set{},
		},
		{
			name: "nobody is empty (with intersection)",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"two": {}, "three": {}}},
			want: Set{"two": {}},
		},
		{
			name: "nobody is empty (without intersection)",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"three": {}, "four": {}}},
			want: Set{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Intersection(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type args struct {
		other Set
	}
	tests := []struct {
		name string
		set  Set
		args args
		want Set
	}{
		{
			name: "both are empty",
			set:  Set{},
			args: args{other: Set{}},
			want: Set{},
		},
		{
			name: "nobody is empty (with intersection)",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"two": {}, "three": {}}},
			want: Set{"one": {}},
		},
		{
			name: "nobody is empty (without intersection)",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"three": {}, "four": {}}},
			want: Set{"one": {}, "two": {}},
		},
		{
			name: "nobody is empty (same)",
			set:  Set{"one": {}, "two": {}},
			args: args{other: Set{"one": {}, "two": {}}},
			want: Set{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Difference(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
