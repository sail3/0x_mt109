package main

import (
	"fmt"
)

func main() {
	fmt.Println("init msg")
}

func RemoveBlackPixels(matrix [][]int) [][]int {
	var result = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		result[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		result = drawPixel(result, buscarPixel(i, 0, matrix, make(map[string][]int)))
		result = drawPixel(result, buscarPixel(0, i, matrix, make(map[string][]int)))
		result = drawPixel(result, buscarPixel(i, len(matrix)-1, matrix, make(map[string][]int)))
		result = drawPixel(result, buscarPixel(len(matrix)-1, i, matrix, make(map[string][]int)))
	}
	return result
}

func drawPixel(matrix [][]int, pixels map[string][]int) [][]int {
	for _, v := range pixels {
		matrix[v[1]][v[0]] = 1
	}
	return matrix
}

func buscarPixel(posX int, posY int, matrix [][]int, pixels map[string][]int) map[string][]int {
	if posX < len(matrix[0]) && posY < len(matrix) && posY > -1 && posX > -1 && matrix[posY][posX] == 1 && pixels[fmt.Sprintf("%d%d", posX, posY)] == nil {
		for _, v := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			pixels[fmt.Sprintf("%d%d", posX, posY)] = []int{posX, posY}
			buscarPixel(posX+v[0], posY+v[1], matrix, pixels)
		}
	}
	return pixels
}
