package shell

func InsertSort(array []int) {
	length := len(array)

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				// swap
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
}

// Sort sort array with shell algorithm
// ascending array: 1,4,13,40,121,364...
func Sort(array []int) {
	length := len(array)
	step := 1
	for step < length/3 {
		step = step*3 + 1
	}
	for step > 0 {
		for i := step; i < length; i ++ {
			for j := i; j >= step; j -= step {
				if array[j] < array[j-step] {
					// swap
					array[j-step], array[j] = array[j], array[j-step]
				}
			}
		}
		step /= 3
	}
}
