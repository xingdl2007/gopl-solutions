package main

import (
	"fmt"
	"time"
)

// recursive version
func permute(nums []int) [][]int {
	var ret [][]int
	size := len(nums)
	if size == 1 {
		ret = append(ret, nums)
	} else {
		p := permute(nums[:size-1])
		for i := 0; i < len(p); i++ {
			for j := 0; j < size; j++ {
				// must be new each iteration
				var tmp []int
				tmp = append(tmp, p[i][:j]...)
				tmp = append(tmp, nums[size-1])
				tmp = append(tmp, p[i][j:]...)

				ret = append(ret, tmp)
			}
		}
	}
	return ret
}

func judge(nums []int, n int) bool {
	var sum int

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	// row
	for i := 1; i < n; i++ {
		var tmp int
		for j := 0; j < n; j++ {
			tmp += nums[i*n+j]
		}
		if tmp != sum {
			return false
		}
	}

	// column
	for i := 0; i < n; i++ {
		var tmp int
		for j := 0; j < n; j++ {
			tmp += nums[i+j*n]
		}
		if tmp != sum {
			return false
		}
	}

	// diagonal
	var tmp int
	for i := 0; i < n; i++ {
		tmp += nums[i+i*n]
	}
	if tmp != sum {
		return false
	}

	tmp = 0
	for i := 0; i < n; i++ {
		tmp += nums[n-1-i+i*n]
	}
	if tmp != sum {
		return false
	}

	return true
}

func magicSquare(n int) (count int) {
	defer trace("magicSquare")()
	var arr = make([]int, n*n)
	for i := 0; i < n*n; i++ {
		arr[i] = i + 1
	}
	perms := permute(arr)

	for _, p := range perms {
		if judge(p, n) {
			fmt.Println(p)
			count++
		}
	}
	return
}

func trace(f string) func() {
	start := time.Now()
	fmt.Printf("Entering %s\n", f)
	return func() {
		fmt.Printf("Exiting %s, elapsed %fs\n", f, time.Since(start).Seconds())
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

	fmt.Println(judge([]int{2, 7, 6, 9, 5, 1, 4, 3, 8}, 3))
	fmt.Println(judge([]int{2, 7, 6, 9, 1, 5, 4, 3, 8}, 3))

	fmt.Println(magicSquare(3))
}
