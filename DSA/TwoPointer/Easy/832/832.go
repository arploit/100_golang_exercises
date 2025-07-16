package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	flipImage(image)
	InverImage(image)
	return image
}

func flipImage(image [][]int) [][]int {
	for index := range image {
		left := 0
		right := len(image[index]) - 1
		for left < right {
			image[index][left], image[index][right] = image[index][right], image[index][left]
			left++
			right--
		}
	}
	return image
}

func InverImage(image [][]int) [][]int {
	for index := range image {
		for i := range image[index] {
			image[index][i] ^= 1
		}
	}

	return image
}

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	flipAndInvertImage(image)
	fmt.Println("FlipAndInvertedImage", image)
}
