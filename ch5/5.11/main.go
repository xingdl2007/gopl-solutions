// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
	"log"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	// a loop
	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	order, acyclic := topoSort(prereqs)
	if !acyclic {
		log.Fatalln("Cyclic prerequisite")
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// false means found cyclic prerequisite
func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)

	// record all uncompleted search
	stack := make(map[string]bool)
	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// start search from `item`
				stack[item] = true
				if !visitAll(m[item]) {
					return false
				}
				// `item` is done
				stack[item] = false
				order = append(order, item)
			} else if stack[item] {
				// detect a loop
				return false
			}
		}
		return true
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// have valid toposort
	if found := visitAll(keys); found {
		return order, true
	}
	return nil, false
}

//!-main
