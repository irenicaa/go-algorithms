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
