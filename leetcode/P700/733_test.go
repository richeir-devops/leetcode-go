package p700

import (
	"fmt"
	"testing"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	result := recursion(&image, sr, sc, newColor)
	return result
}

//递归算法
func recursion(imageRef *[][]int, x int, y int, newValue int) [][]int {
	//找到对应的4个数，如果数字和原版数字相同，则继续执行递归
	var image = *imageRef
	var orgValue = image[x][y]

	if image[x][y] == newValue {
		return image
	}

	image[x][y] = newValue

	top, left, bottom, right := findRonud(imageRef, x, y, orgValue)
	if top != -1 {
		recursion(imageRef, x-1, y, newValue)
	}

	if left != -1 {
		recursion(imageRef, x, y-1, newValue)
	}

	if bottom != -1 {
		recursion(imageRef, x+1, y, newValue)
	}

	if right != -1 {
		recursion(imageRef, x, y+1, newValue)
	}

	return image
}

func findRonud(imageRef *[][]int, x int, y int, orgValue int) (int, int, int, int) {
	var image = *imageRef
	lenX := len(image) - 1
	lenY := len(image[0]) - 1

	if x < 0 || y < 0 {
		return -1, -1, -1, -1
	}

	if x > lenX || y > lenY {
		return -1, -1, -1, -1
	}
	top := -1
	left := -1
	bottom := -1
	right := -1

	if x-1 < 0 || y < 0 {
		top = -1
	} else {
		top = image[x-1][y]
		if top != orgValue {
			top = -1
		}
	}

	if x < 0 || y-1 < 0 {
		left = -1
	} else {
		left = image[x][y-1]
		if left != orgValue {
			left = -1
		}
	}

	if x+1 > lenX {
		bottom = -1
	} else {
		bottom = image[x+1][y]
		if bottom != orgValue {
			bottom = -1
		}
	}

	if y+1 > lenY {
		right = -1
	} else {
		right = image[x][y+1]
		if right != orgValue {
			right = -1
		}
	}

	return top, left, bottom, right
}

func printArray(image *[][]int) {

	var image2 [][]int
	image2 = *image

	fmt.Printf("%v \n", image)
	fmt.Printf("%v \n", image2)
}

func Test1(t *testing.T) {
	var image [][]int
	image = append(image, []int{1, 1, 1})
	image = append(image, []int{1, 1, 0})
	image = append(image, []int{1, 0, 1})

	// sr := 1
	// sc := 1
	// orginalValue = image[sr][sc]
	// newValue = 2

	t.Log("733")
	t.Log(len(image[0]))
	t.Log(len(image))

	result := floodFill(image, 1, 1, 2)

	printArray(&result)

	// printArray(&image)

	// image[0][0] = 5

	// printArray(&image)
}
