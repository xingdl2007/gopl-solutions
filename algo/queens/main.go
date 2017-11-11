package main

import "fmt"

// invalid: false
// O(n): i=> row, res[i]=> column
func isValid(res []int, k int) bool {
	for i := 0; i < k; i++ {
		if res[i] == res[k] ||
			i+res[i] == k+res[k] ||
			i-res[i] == k-res[k] {
			return false
		}
	}
	return true
}

// if only restrict row/column, then will have N! solution
func isValid2(res []int, k int) bool {
	for i := 0; i < k; i++ {
		if res[i] == res[k] {
			return false
		}
	}
	return true
}

// 8-queens problems
// first try to find one solution, then continue
func eightQueens() (count int) {
	const N = 8
	res := [...]int{-1, -1, -1, -1, -1, -1, -1, -1}

	for k := 0; k >= 0; {
		// find one solution
		for k != N && k >= 0 {
			res[k]++
			// subtle: backtrace
			if res[k] == N {
				res[k] = -1
				k--
				continue
			}
			if isValid(res[:], k) {
				k++
			}
		}
		// counter
		if k == N {
			count++
			fmt.Println(res[:])

			// minor optimization
			res[k-1] = -1
			// restart start from last but one
			k -= 2
		}
	}
	return
}

// n-queens problems
func NQueens(N int) (count int) {
	res := make([]int, N, N)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for k := 0; k >= 0; {
		// find one solution
		for k != N && k >= 0 {
			res[k]++
			// subtle: backtrace
			if res[k] == N {
				res[k] = -1
				k--
				continue
			}
			if isValid(res[:], k) {
				k++
			}
		}
		// counter
		if k == N {
			count++

			// minor optimization
			res[k-1] = -1
			// restart start from last but one
			k -= 2
		}
	}
	return
}

func main() {
	t := []int{1, 3, 0, 2}
	fmt.Println(isValid(t, 0))
	fmt.Println(isValid(t, 1))
	fmt.Println(isValid(t, 2))
	fmt.Println(isValid(t, 3))

	fmt.Println(eightQueens())
	fmt.Println(NQueens(12))
}
