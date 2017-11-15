// bipartite (2-colorable) graph detection

package main

import "fmt"

// connected component
func connectedComponent(graph [][]int) (ret [][]int) {
	m := make(map[int]bool)

	var dfs func(n int) []int
	dfs = func(n int) (comp []int) {
		m[n] = true
		comp = append(comp, n)

		for _, v := range graph[n] {
			if !m[v] {
				comp = append(comp, dfs(v)...)
			}
		}
		return comp
	}

	for i := 0; i < len(graph); i++ {
		if !m[i] {
			ret = append(ret, dfs(i))
		}
	}
	return ret
}

// ref: Algorithms 4ed (Robert Sedgewick, p.350)
func CC(graph [][]int) (ret [][]int) {
	m := make(map[int]bool)
	id := make([]int, len(graph))
	var counter int

	var dfs func(n int)
	dfs = func(n int) {
		m[n] = true
		id[n] = counter
		for _, v := range graph[n] {
			if !m[v] {
				dfs(v)
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		if !m[i] {
			dfs(i)
			counter++
		}
	}

	ret = make([][]int, counter)
	for n, v := range id {
		ret[v] = append(ret[v], n)
	}

	return ret
}

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

	bi := [][]int{
		{1, 5},
		{0, 2, 4},
		{1, 3},
		{2, 4},
		{1, 3, 5},
		{0, 4},
	}

	nbi := [][]int{
		{1, 2},
		{0, 2, 3},
		{0, 1, 3},
		{1, 2},
	}

	// Connected Component
	fmt.Println(connectedComponent(graph))

	// better solution
	fmt.Println(CC(graph))
}
