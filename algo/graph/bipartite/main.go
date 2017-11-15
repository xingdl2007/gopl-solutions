// bipartite (2-colorable) graph detection

package main

import "fmt"

func bipartiteDFS(graph [][]int) bool {
	var m = make(map[int]bool)
	var color = make([]bool, len(graph))
	isBipartite := true

	// DFS
	var dfs func(n int)
	dfs = func(n int) {
		m[n] = true
		for _, a := range graph[n] {
			if !m[a] {
				color[a] = !color[n]
				dfs(a)
			} else {
				if color[a] == color[n] {
					isBipartite = false
				}
			}
		}
	}

	// scan all vertex
	for i := 0; i < len(graph); i++ {
		if !m[i] {
			dfs(i)
		}
	}

	return isBipartite
}

// BFS
func bipartiteBFS(graph [][]int) bool {
	var m = make(map[int]bool)
	var color = make([]bool, len(graph))
	isBipartite := true

	var scanBuf []int

	// BFS
	var bfs func(n int)
	bfs = func(n int) {
		scanBuf = append(scanBuf, n)
		for len(scanBuf) > 0 {
			n := scanBuf[0]
			if !m[n] {
				m[n] = true
				for _, v := range graph[n] {
					if !m[v] {
						scanBuf = append(scanBuf, v)
						color[v] = !color[n]
					} else {
						if color[v] == color[n] {
							isBipartite = false
						}
					}
				}
			}
			scanBuf = scanBuf[1:]
		}
	}

	// scan all vertex
	for i := 0; i < len(graph); i++ {
		if !m[i] {
			bfs(i)
		}
	}

	return isBipartite
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

	// DFS
	fmt.Println(bipartiteDFS(graph))
	fmt.Println(bipartiteDFS(bi))
	fmt.Println(bipartiteDFS(nbi))

	// BFS
	fmt.Println()
	fmt.Println(bipartiteBFS(graph))
	fmt.Println(bipartiteBFS(bi))
	fmt.Println(bipartiteBFS(nbi))
}
