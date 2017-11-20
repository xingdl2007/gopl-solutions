package main

import (
	"testing"
	"sort"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		array  []string
		key    string
		result bool
	}{
		{[]string{"hello", "world"}, "hello", true},
		{[]string{"hello", "world"}, "world", true},
		{[]string{"hello", "world", "again", "what", "up"}, "up", true},
	}

	for _, test := range tests {
		sort.Strings(test.array)
		if binarySearch(test.array, test.key) != test.result {
			t.Errorf("test %q whether contains %q, result: %t, want: %t",
				test.array, test.key, !test.result, test.result)
		}
	}
}
