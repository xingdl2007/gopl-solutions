// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package split_test

import (
	"testing"
	"strings"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s, sep string
		want   int
	}{
		{"a:b:c", ":", 3},
		{"a b c", ":", 1},
		{"^a^b c", "^", 3},
	}
	for _, test := range tests {
		if got := len(strings.Split(test.s, test.sep)); got != test.want {
			t.Errorf("Split(%q,%q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
