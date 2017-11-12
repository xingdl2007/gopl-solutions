// bipartite (2-colorable) graph detection

package main

import "fmt"

func loopDetect(graph [][]int) (has bool) {
	m := make(map[int]bool)

	var dfs func(n, f int)
	dfs = func(n, f int) {
		m[n] = true
		for _, v := range graph[n] {
			if !m[v] {
				dfs(v, n)
			} else {
				if v != f {
					has = true
				}
			}
		}
	}
	for i := 0; i < len(graph); i++ {
		if !m[i] {
			dfs(i, i)
		}
	}
	return
}

// TODO: path, give some path

func main() {
	// tinyG, adjacent list
	graph := [][]int{
		{1, 2, 5, 6},
		{0},
		{0},
		{4, 5},
		{3, 5, 6},
		{0, 3, 4},
		{0, 4},
		{8},
		{7},
		{10, 11, 12},
		{9},
		{9, 12},
		{9, 11},
	}

	// Connected Component
	fmt.Println(loopDetect(graph))
}
