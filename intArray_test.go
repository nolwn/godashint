package godashint

import (
	"testing"
)

// TestFindIndex does what it sounds like. Fucking Go...
func TestFindIndex(t *testing.T) {
	arr := []int{3, 5, 1, 12, 6, 34, 1}
	find05 := func(value int) bool { return value == 5 }
	find01 := func(value int) bool { return value == 1 }
	find03 := func(value int) bool { return value == 3 }
	find40 := func(value int) bool { return value == 40 }
	findEv := func(value int) bool { return value%2 == 0 }
	findBg := func(value int) bool { return value > 30 }
	var tests = []struct {
		predicate    func(int) bool
		from, expect int
	}{
		{find05, 0, 1},
		{find01, 0, 2},
		{find03, 0, 0},
		{find40, 0, -1},
		{findEv, 0, 3},
		{findBg, 0, 5},
		{find05, 2, -1},
		{find05, 1, 1},
		{find01, 4, 6},
		{findEv, 4, 4},
	}

	for _, vars := range tests {
		var found int = findIndex(arr, vars.predicate, vars.from)
		if found != vars.expect {
			t.Errorf("Expected: %d\tFound: %d", vars.expect, found)
		}
	}
}
