package main

import "fmt"

// 8.1 12
func cal(n int) float64 {
	row, col := (n+1)/2, (n+1)/2
	var pro = make([][]float64, row)
	for i := 0; i < row; i++ {
		pro[i] = make([]float64, col)
		pro[i][0] = 1.0
	}

	for i := 1; i < row; i++ {
		for j := 1; j < row; j++ {
			pro[i][j] = 0.6*pro[i-1][j] + 0.4*pro[i][j-1]
		}
	}
	return pro[row-1][col-1]
}

func main() {
	fmt.Println(cal(7))
}
