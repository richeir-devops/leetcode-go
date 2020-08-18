package daily202008

import (
	"fmt"
	"testing"
)

// Box 类型
type BoxList struct {
	boxes          []int
	segmentsBegins []int
	segmentsEnds   []int
	segmentCount   int
	boxesLength    int
}

// 初始化方法
func (b BoxList) Init(gavingBoxes []int) BoxList {
	b.boxes = gavingBoxes
	b.segmentsBegins = make([]int, len(gavingBoxes)+1)
	b.segmentsEnds = make([]int, len(gavingBoxes)+1)
	b.segmentCount = 0
	b.boxesLength = len(gavingBoxes)
	return b
}

// 把现有数组划分好分组
func (b BoxList) getSegment() int {
	v1 := 0
	v2 := 0
	segmentIndex := 1
	for i := 0; i < b.boxesLength; i++ {
		if b.boxes[i] != 0 {
			v1 = b.boxes[i]
			b.segmentsBegins[segmentIndex] = i
			for j := i; j < b.boxesLength; j++ {
				v2 = b.boxes[j]
				if v1 == v2 {
					b.segmentsEnds[segmentIndex] = j
					i = j
				} else {
					segmentIndex++
					break
				}
			}
		}
	}

	return segmentIndex
}

//移除一个片段
func removeSegment(boxes []int, segmentIndex int) ([]int, int) {
	returnValue := 0
	segment := 0
	segmentLength := 1
	// currentNum := boxes[0]
	for i := 0; i < len(boxes)-1; i++ {
		var v1 = boxes[i]
		var v2 = boxes[i+1]

		if v1 != 0 {
			segment++

			if i == len(boxes)-2 && v2 != 0 && v1 != v2 {
				segment++
				if segment == segmentIndex {
					boxes[len(boxes)-1] = 0
					return boxes, 1
				}
			}

			if segment == segmentIndex {
				for j := i; j < len(boxes)-1; j++ {
					if boxes[j] == boxes[j+1] {
						segmentLength++
					} else {
						break
					}
				}

				returnValue = segmentLength * segmentLength
				for j := i; j < i+segmentLength; j++ {
					boxes[j] = 0
				}

				return boxes, returnValue
			}
			// if v1 != v2 && v2 != 0 {

			// }
		}
	}

	if segment == 0 {
		for i := 0; i < len(boxes); i++ {
			boxes[i] = 0
		}
		return boxes, 1
	}

	return boxes, returnValue
}

func removeBoxes(boxes []int) int {
	//计算传入数组的分段
	segment := 0
	currentNum := boxes[0]
	for i := 0; i < len(boxes)-1; i++ {
		if currentNum != boxes[i+1] {
			segment++
		}
	}
	fmt.Printf("segment: %v \n", segment)

	row, column := segment, segment
	var answer [][]int
	for i := 0; i < row; i++ {
		inline := make([]int, column)
		answer = append(answer, inline)
	}
	fmt.Println(answer)

	//记录每一步的答案
	//i 表示刚开始消除的第 i 个数字
	//j 表示接下来的方案
	for i := 0; i < segment; i++ {
		for j := 0; j < segment-i-1; j++ {
			//nswer[]
		}
	}

	return 1
}

func Test_1(t *testing.T) {
	t.Log("go")
	boxes := []int{1, 3, 2, 2, 2, 3, 4, 3, 1}
	result := removeBoxes(boxes)
	t.Log(result)
	t.Log(result == 23)
}

func Test_2(t *testing.T) {
	boxes := []int{1, 3, 2, 2, 2, 3, 4, 3, 1}
	val := 0
	boxes, val = removeSegment(boxes, 1)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 1)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 1)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 1)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 2)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 2)
	t.Log(boxes, val)

	boxes, val = removeSegment(boxes, 1)
	t.Log(boxes, val)
}

func Test_3(t *testing.T) {
	boxes := []int{1, 3, 0, 0, 0, 3, 4, 3, 1}
	mybox := new(BoxList).Init(boxes)
	mybox.segmentCount = mybox.getSegment()
	t.Log(mybox.boxesLength)
	t.Log(mybox.segmentsBegins)
	t.Log(mybox.segmentsEnds)
	t.Log(mybox.segmentCount)
}
