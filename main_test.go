package main

import (
	"reflect"
	"testing"
)

func TestBoundary(t *testing.T) {
	root := tn(1,
		tn(2, tn(4, nil, nil), tn(5, tn(8, nil, nil), tn(9, nil, nil))),
		tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	want := []int{1, 2, 4, 8, 9, 6, 7, 3}
	if got := Boundary(root); !reflect.DeepEqual(got, want) {
		t.Errorf("Boundary = %v, want %v", got, want)
	}

	if got := Boundary(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Boundary(nil) = %v", got)
	}
	// Single node counted once.
	if got := Boundary(tn(7, nil, nil)); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("single = %v", got)
	}
	// Left-skewed chain: all nodes appear once.
	chain := tn(1, tn(2, tn(3, nil, nil), nil), nil)
	if got := Boundary(chain); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("chain = %v", got)
	}
	// Right-skewed chain.
	rchain := tn(1, nil, tn(2, nil, tn(3, nil, nil)))
	if got := Boundary(rchain); !reflect.DeepEqual(got, []int{1, 3, 2}) {
		t.Errorf("rchain = %v", got)
	}
}
