package main

import "fmt"

// versatile function
func LomutoPartition(a []int, l, r int) int {
	s := l
	for i := l + 1; i <= r; i++ {
		if a[i] < a[l] {
			s += 1
			a[s], a[i] = a[i], a[s]
		}
	}
	a[s], a[l] = a[l], a[s]
	return s
}

// recursive solution
func quickSelect(a []int, l, r, k int) int {
	s := LomutoPartition(a, l, r)
	if s == k-1 {
		return a[s]
	} else if s > k-1 {
		return quickSelect(a, l, s-1, k)
	} else {
		return quickSelect(a, s+1, r, k)
	}
}

// iterative solution
func quickSelect2(a []int, l, r, k int) int {
	var s int
	for s = LomutoPartition(a, l, r); s != k-1; {
		if s > k-1 {
			s = LomutoPartition(a, l, s-1)
		} else {
			s = LomutoPartition(a, s+1, r)
		}
	}
	return a[s]
}

func QuickSelect(a []int, k int) int {
	if k < 1 || k > len(a) {
		return 0
	}
	return quickSelect2(a, 0, len(a)-1, k)
}

func main() {
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 1))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 2))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 3))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 4))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 5))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 6))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 7))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 8))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 9))

	// boundary check
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 0))
	fmt.Println(QuickSelect([]int{4, 1, 10, 8, 7, 12, 9, 2, 15}, 10))
}
