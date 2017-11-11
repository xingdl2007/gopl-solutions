// Tower of hanoi
package main

import "fmt"

var cnt int

// classic recursive solution, slice is trick
// if a slice is changed in length, it must be returned
func hanoi(from []int, f string, aux []int, a string, to []int, t string) []int {
	// for verification: must be 2^len(from) -1
	cnt++
	if len(from) == 0 {
		return to
	}
	if len(from) == 1 && len(to) == 0 {
		to = append(to, from[0])
		fmt.Printf("Move disk %d from %s to %s\n", from[0], f, t)
	} else {
		aux = hanoi(from[1:], f, to[len(to):], t, aux[len(aux):], a)
		to = append(to, from[0])
		fmt.Printf("Move disk %d from %s to %s\n", from[0], f, t)
		to = append(to, hanoi(aux, a, from[:0], f, to[len(to):], t)...)
	}
	return to
}

// Only print specific move process
// since the result is obvious
func hanoi3(num int, from, aux, to string) {
	if num > 0 {
		hanoi3(num-1, from, to, aux)
		fmt.Printf("Move disk %d from %s to %s\n", num, from, to)
		hanoi3(num-1, aux, from, to)
	}
}

// iterative solution
// ref: https://en.wikipedia.org/wiki/Tower_of_Hanoi#Iterative_solution
// Use slice pointer (*[]int) instead of slice in hanoi2 , add a layer
// of indirection, eliminate similar code
func hanoi4(from, aux, to *[]int) {
	var iter int
	var direction []*[]int

	// for printing
	tower := map[*[]int]string{from: "A", aux: "B", to: "C"}

	if len(*from)%2 == 0 {
		direction = []*[]int{from, aux, to} // means from->aux->to->from
	} else {
		direction = []*[]int{from, to, aux} // means from->to->aux->from
	}

	// when from&aux are empty, we'are done
	for len(*from) > 0 || len(*aux) > 0 {
		// smallest at peg1, move it to peg2 according `direction`
		peg1 := direction[iter%3]
		peg2 := direction[(iter+1)%3]
		peg3 := direction[(iter+2)%3]
		iter++

		// move smallest disk from peg1->peg2
		fmt.Printf("Move disk %d from %s to %s.\n", (*peg1)[len(*peg1)-1], tower[peg1], tower[peg2])
		*peg2 = append(*peg2, (*peg1)[len(*peg1)-1])
		*peg1 = (*peg1)[:len(*peg1)-1]

		// move non-smallest disk either peg1->peg3 or peg3->peg1
		l, r := len(*peg1), len(*peg3)
		if l > 0 || r > 0 {
			if l == 0 || (l != 0 && r != 0 && (*peg1)[l-1] > (*peg3)[r-1]) {
				fmt.Printf("Move disk %d from %s to %s.\n", (*peg3)[r-1], tower[peg3], tower[peg1])
				*peg1 = append(*peg1, (*peg3)[r-1])
				*peg3 = (*peg3)[:r-1]
			} else {
				fmt.Printf("Move disk %d from %s to %s.\n", (*peg1)[l-1], tower[peg1], tower[peg3])
				*peg3 = append(*peg3, (*peg1)[l-1])
				*peg1 = (*peg1)[:l-1]
			}
		}
	}
}

// iterative solution
// ref: https://en.wikipedia.org/wiki/Tower_of_Hanoi#Iterative_solution
// straightforward but in such a situation, slice is kind of cumbersome
// need another layer of indirect, see hanoi4 with slice pointer
func hanoi2(from []int, f string, aux []int, a string, to []int, t string) []int {
	var iter int
	var direction []string

	if len(from)%2 == 0 {
		direction = []string{f, a, t} // means from->aux->to->from
	} else {
		direction = []string{f, t, a} // means from->to->aux->from
	}

	// first move smallest disk, according to direction
	// then move non-smallest disk, only one legal move
	// repeat until complete
	for len(from) > 0 || len(aux) > 0 {
		// smallest at peg1, move to peg2, must exist
		peg1 := direction[iter%3]
		peg2 := direction[(iter+1)%3]
		iter++

		if peg1 == f {
			// from->aux or from->to
			if peg2 == a {
				aux = append(aux, from[len(from)-1])
			} else {
				to = append(to, from[len(from)-1])
			}
			from = from[:len(from)-1]
		} else if peg1 == a {
			// aux->to or aux->from
			if peg2 == t {
				to = append(to, aux[len(aux)-1])
			} else {
				from = append(from, aux[len(aux)-1])
			}
			aux = aux[:len(aux)-1]
		} else {
			// to->from or to->aux
			if peg2 == f {
				from = append(from, to[len(to)-1])
			} else {
				aux = append(aux, to[len(to)-1])
			}
			to = to[:len(to)-1]
		}

		// then move non-smallest disk
		if peg2 == f && (len(aux) > 0 || len(to) > 0) {
			// aux->to or to->aux
			if len(aux) == 0 {
				aux = append(aux, to[len(to)-1])
				to = to[:len(to)-1]
			} else if len(to) == 0 {
				to = append(to, aux[len(aux)-1])
				aux = aux[:len(aux)-1]
			} else {
				if aux[len(aux)-1] < to[len(to)-1] {
					to = append(to, aux[len(aux)-1])
					aux = aux[:len(aux)-1]
				} else {
					aux = append(aux, to[len(to)-1])
					to = to[:len(to)-1]
				}
			}
		} else if peg2 == a && (len(from) > 0 || len(to) > 0) {
			// from->to or to->from
			if len(from) == 0 {
				from = append(from, to[len(to)-1])
				to = to[:len(to)-1]
			} else if len(to) == 0 {
				to = append(to, from[len(from)-1])
				from = from[:len(from)-1]
			} else {
				if from[len(from)-1] < to[len(to)-1] {
					to = append(to, from[len(from)-1])
					from = from[:len(from)-1]
				} else {
					from = append(from, to[len(to)-1])
					to = to[:len(to)-1]
				}
			}
		} else if len(from) > 0 || len(aux) > 0 {
			// from->aux or aux->from
			if len(from) == 0 {
				from = append(from, aux[len(aux)-1])
				aux = aux[:len(aux)-1]
			} else if len(aux) == 0 {
				aux = append(aux, from[len(from)-1])
				from = from[:len(from)-1]
			} else {
				if from[len(from)-1] < aux[len(aux)-1] {
					aux = append(aux, from[len(from)-1])
					from = from[:len(from)-1]
				} else {
					from = append(from, aux[len(aux)-1])
					aux = aux[:len(aux)-1]
				}
			}
		}
	}
	return to
}

func main() {
	from := []int{5, 4, 3, 2, 1}
	aux := make([]int, 0, 5)
	to := make([]int, 0, 5)

	// offer tower name for print
	//to = hanoi2(from, "A", aux, "B", to, "C")

	hanoi3(0, "A", "B", "C")
	hanoi4(&from, &aux, &to)

	fmt.Println(from, aux, to)
}
