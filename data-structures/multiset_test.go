package datastructures

import (
	"reflect"
	"testing"
)

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
			if got := NewMultiset(tt.args.items...); !reflect.DeepEqual(got, tt.wantSet) {
				t.Errorf("NewMultiset() = %v, want %v", got, tt.wantSet)
			}
		})
	}
}
