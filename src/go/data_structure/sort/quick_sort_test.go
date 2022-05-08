package sort

import "testing"

func quickSort(arr *[]int, low int, high int) {
	if low < high {
		i := partition(arr, low, high)
		quickSort(arr, low, i-1)
		quickSort(arr, i+1, high)
	}
}

func partition(arr *[]int, low int, high int) int {
	numArray := *arr

	//将该数组第一个元素设置为比较元素
	x := numArray[low]
	//指向数组头的指针
	i := low
	//指向数组尾的指针
	j := high

	for i < j {
		//从右至左找到第一个小于比较元素的数
		for i < j && numArray[j] >= x {
			j--
		}

		//从左至右找到第一个大于比较元素的数
		for i < j && numArray[i] <= x {
			i++
		}

		//将大数与小数交换
		if i != j {
			swap(arr, i, j)
		}
	}
	//将比较元素交换到期望位置
	swap(arr, low, i)

	return i
}

func swap(arr *[]int, i int, j int) {
	numArray := *arr
	temp := numArray[i]
	numArray[i] = numArray[j]
	numArray[j] = temp
}

func Test1(t *testing.T) {
	t.Log("quick sort")
	arr := []int{5, 7, 1, 6, 4, 8, 3, 2}
	t.Log(arr)
	swap(&arr, 1, 2)
	t.Log(arr)
	quickSort(&arr, 0, len(arr)-1)
	t.Log(arr)
}
