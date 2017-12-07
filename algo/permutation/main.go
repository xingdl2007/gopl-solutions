package main

import (
	"fmt"
	"sort"
)

// this article introduces common algorithms for permutation generation
// https://blog.cykerway.com/posts/2016/12/25/permutation-generation-algorithms.html

// recursive version, straight-forward
// Can be modified to satisfy minimal-change requirement
func Recursive(n []int) [][]int {
	size := len(n)
	if size <= 1 {
		return [][]int{n}
	}
	var res [][]int
	for _, p := range Recursive(n[:size-1]) {
		// even permutation, descending order
		for i := 0; i <= len(p); i++ {
			var tmp = make([]int, 0, len(n))
			tmp = append(tmp, p[:i]...)
			tmp = append(tmp, n[size-1])
			tmp = append(tmp, p[i:]...)
			res = append(res, tmp)
		}
	}
	return res
}

// easy to understand, and faster than Recursive
func Backtracking(l []int) [][]int {
	var perm func([]int, int)
	var res [][]int
	perm = func(l []int, n int) {
		if n == 1 {
			var tmp = make([]int, len(l))
			copy(tmp, l)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			l[n-1], l[i] = l[i], l[n-1]
			perm(l, n-1)
			l[n-1], l[i] = l[i], l[n-1]
		}
	}
	perm(l, len(l))
	return res
}

// STJ(plain changes): minimal-change requirement
// https://en.wikipedia.org/wiki/Steinhaus%E2%80%93Johnson%E2%80%93Trotter_algorithm
//
// The Steinhaus–Johnson–Trotter algorithm or Johnson–Trotter algorithm, also called plain changes,
// is an algorithm named after Hugo Steinhaus, Selmer M. Johnson and Hale F. Trotter that generates
// all of the permutations of n elements. Each permutation in the sequence that it generates differs
// from the previous permutation by swapping two adjacent elements of the sequence.
// Equivalently, this algorithm finds a Hamiltonian path in the permutohedron.
//
// Sedgewick (1977) calls it "perhaps the most prominent permutation enumeration algorithm".
// As well as being simple and computationally efficient, it has the advantage that subsequent
// computations on the permutations that it generates may be sped up because these permutations
// are so similar to each other.
// ref: SJT.c
// Actually very interesting algorithm, see Introduction to Design and Analysis of Algorithms
// P.112 for another explanation.
// Impressive, fastest!
func JohnsonTrotter(num int) [][]int {
	var res [][]int
	p := make([]int, num+1)   // permutation, one extra space
	pi := make([]int, num+1)  // permutation reverse, coordinate
	dir := make([]int, num+1) // dir: -1 or +1
	for i := 0; i <= num; i++ {
		p[i] = i
		pi[i] = i
		dir[i] = -1
	}
	// n: means numbers, 1,2,3,4...n
	var perm func(n int, p, pi, dir []int)
	perm = func(n int, p, pi, dir []int) {
		if n > num {
			var tmp = make([]int, num)
			copy(tmp, p[1:])
			res = append(res, tmp)
			return
		}
		// bigger number have high priority
		perm(n+1, p, pi, dir)
		for i := 0; i < n-1; i++ {
			// swap n with its neighbor
			z := p[pi[n]+dir[n]]
			p[pi[n]] = z
			p[pi[n]+dir[n]] = n
			pi[z] = pi[n]
			pi[n] += dir[n]
			// so when exchange numbers, move bigger number
			perm(n+1, p, pi, dir)
		}
		dir[n] = -dir[n]
	}
	perm(1, p, pi, dir)
	return res
}

// Heap's algorithm, same level as Backtracking, little speed-up
// This is an amazing algorithm which reduces the number of element exchanges down to 1.
// This means there is no backtracking (as you can compare it with Algorithm 1).
// It exchanges the last element with:
//
// The first element, if array length is odd.
// The i-th element, if array length is even.
//
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
// "The algorithm minimizes movement: it generates each permutation from the previous one
// by interchanging a single pair of elements; the other n−2 elements are not disturbed.
// In a 1977 review of permutation-generating algorithms, Robert Sedgewick concluded that
// it was at that time the most effective algorithm for generating permutations by computer."

// Details of the algorithm:
// Suppose we have a permutation containing n different elements.
// Heap found a systematic method for choosing at each step a pair of elements to switch,
// in order to produce every possible permutation of these elements exactly once.
// Let us describe Heap's method in a recursive way. First we set a counter i to 0.
// Now we perform the following steps repeatedly until i is equal to n.
// We use the algorithm to generate the (n−1)! permutations of the first n−1 elements,
// adjoining the last element to each of these. This generates all of the permutations
// that end with the last element. Then if n is odd, we switch the first element and the last one,
// while if n is even we can switch the ith element and the last one
// (there is no difference between n even and odd in the first iteration).
// We add one to the counter i and repeat. In each iteration, the algorithm will produce all of the
// permutations that end with the element that has just been moved to the "last" position.

// TODO: why it works?
func Heap(nums []int) [][]int {
	var res [][]int
	var heap func([]int, int)
	heap = func(l []int, n int) {
		if n == 1 {
			var tmp = make([]int, len(l))
			copy(tmp, l)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			heap(l, n-1)
			// key part
			pos := i
			if n%2 == 0 {
				pos = 0
			}
			l[pos], l[n-1] = l[n-1], l[pos]
		}
	}
	heap(nums, len(nums))
	return res
}

// LexicographicPermute: Introduction to The Design and Analysis of Algorithms 3ed, p.112
func Lexicographic(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	sort.Ints(nums)

	for {
		tmp := make([]int, n)
		copy(tmp, nums)
		ans = append(ans, tmp)

		pos := -1
		for i := n - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				pos = i
				break
			}
		}
		if pos == -1 {
			break
		}

		sort.Ints(nums[pos:])
		// swap with the smallest one which is bigger than nums[pos]
		for i := pos; i < n; i++ {
			if nums[i] > nums[pos-1] {
				nums[i], nums[pos-1] = nums[pos-1], nums[i]
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(Recursive([]int{1, 2, 3}))
	fmt.Println(Backtracking([]int{1, 2, 3}))
	fmt.Println("heap:")
	fmt.Println(Heap([]int{1, 2, 3}))
	fmt.Println(Lexicographic([]int{1, 2, 3}))
	fmt.Println(JohnsonTrotter(3))

	fmt.Println()
	fmt.Println(Lexicographic([]int{1, 2, 2, 3}))
}
