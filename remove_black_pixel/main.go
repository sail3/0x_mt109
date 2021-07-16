package main

import "fmt"

func main() {
	fmt.Println("init msg")
}

func RemoveBlackPixels(matrix [][]int) [][]int {
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if matrix[i-1][j] != 1 && matrix[i+1][j] != 1 && matrix[i][j-1] != 1 && matrix[i][j+1] != 1 {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
