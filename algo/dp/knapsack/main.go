package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// very famous problem
func knapsack(goods [][]int, w int) int {
	var value = make([][]int, len(goods)+1)
	for i := 0; i < len(value); i++ {
		value[i] = make([]int, w+1)
	}
	for i := 1; i <= len(goods); i++ {
		for j := 1; j <= w; j++ {
			if goods[i-1][0] > j {
				value[i][j] = value[i-1][j]
			} else {
				value[i][j] = max(value[i-1][j], goods[i-1][1]+value[i-1][j-goods[i-1][0]])
			}
		}
	}
	// last one
	return value[len(value)-1][w]
}

// with memory function, compute as necessary
func knapsack2(goods [][]int, w int) int {
	var value = make([][]int, len(goods)+1)
	for i := 0; i < len(value); i++ {
		value[i] = make([]int, w+1)
	}
	for i := 1; i <= len(goods); i++ {
		for j := 1; j <= w; j++ {
			value[i][j] = -1
		}
	}
	return fill(goods, value, len(value)-1, w)
}

// only will fill item which is needed
func fill(goods, value [][]int, i, j int) int {
	if value[i][j] < 0 {
		if goods[i-1][0] > j {
			value[i][j] = fill(goods, value, i-1, j)
		} else {
			value[i][j] = max(fill(goods, value, i-1, j), goods[i-1][1]+fill(goods, value, i-1, j-goods[i-1][0]))
		}
	}
	return value[i][j]
}

func main() {
	goods := [][]int{{2, 12}, {1, 10}, {3, 20}, {2, 15}}
	w := 5
	fmt.Println(knapsack(goods, w))
	fmt.Println(knapsack2(goods, w))
}
