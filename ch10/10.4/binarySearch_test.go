package main

import (
	"testing"
	"sort"
)

func TestSearchStrings(t *testing.T) {
	var tests = []struct {
		array  []string
		key    string
		result int
	}{
		{[]string{"hello", "world"}, "hello", 0},
		{[]string{"hello", "world"}, "world", 1},
		{[]string{"hello", "world", "again", "what", "up"}, "up", 2},
	}

	for _, test := range tests {
		sort.Strings(test.array)
        if res := sort.SearchStrings(test.array, test.key); res != test.result {
			t.Errorf("test %q whether contains %q, result: %v, want: %v",
				test.array, test.key, res, test.result)
		}
	}
}
