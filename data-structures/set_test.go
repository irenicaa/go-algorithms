package datastructures

import (
	"reflect"
	"testing"
)

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
