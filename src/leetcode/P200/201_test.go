package p200

import (
	"testing"
)

//直接计算会超时
// func rangeBitwiseAnd(m int, n int) int {
// 	sum := m
// 	for i := m; i < n; i++ {
// 		sum = sum & (i + 1)
// 	}

// 	return sum
// }

//先把所有数字转二进制然后存绕数组里面，之后统一计算
//二维数组中，某一列所有数字都是1，最终的 AND 计算的这一位才是 1
//仍然超时
// func rangeBitwiseAnd(m int, n int) int {
// 	finalBinArray := make([]int, 33)
// 	result := 0

// 	var binArray [][]int
// 	//转换数字到二进制数组
// 	for i := m; i <= n; i++ {
// 		binArray = append(binArray, convertIntToBinary(i))
// 	}

// 	//计算最终结果的二进制的数组表示
// 	for i := 32; i >= 0; i-- {
// 		isAllOne := true
// 		for j := 0; j < len(binArray); j++ {
// 			if binArray[j][i] == 0 {
// 				isAllOne = false
// 				break
// 			}
// 		}
// 		if isAllOne {
// 			finalBinArray[i] = 1
// 		}
// 	}

// 	num2 := 1
// 	for i := 32; i >= 0; i-- {
// 		if finalBinArray[i] == 1 {
// 			result += num2
// 		}
// 		num2 *= 2
// 	}

// 	return result
// }

func convertIntToBinary(num int) []int {
	var result = make([]int, 33)
	for i := 32; i > 0; i-- {
		result[i] = num % 2
		num /= 2
		if num == 0 {
			break
		}
	}
	return result
}

//先计算首尾的两个数字 AND 运算的二进制数组
//之后进行再和中间区域的数字一个一个比较（从大到小），如果两个位上面有不同的情况，就设置为0
//如果比较中，数字已经是都是0了，就直接返回0
func rangeBitwiseAnd(m int, n int) int {
	var compareResult = m & n

	for i := n - 1; i >= m+1; i-- {
		compareResult = compareResult & i
		if compareResult == 0 {
			return compareResult
		}
	}

	return compareResult
}

func Test_201_1(t *testing.T) {
	t.Log("201")
	a := 5
	b := 7
	c := 6
	t.Log(a & b & c)

	sum := 5
	for i := 5; i < 7; i++ {
		sum = sum & (i + 1)
	}
	t.Log(sum)

	t.Log(rangeBitwiseAnd(0, 2147483647))

	var bins [][]int
	// binA := make([]int, 32)
	// binA[0] = 1
	// bins = append(bins, binA)
	// t.Log(bins[0])
	// t.Log(bins)

	t.Log(convertIntToBinary(5))
	t.Log(convertIntToBinary(6))
	t.Log(convertIntToBinary(7))

	bins = append(bins, convertIntToBinary(5))
	bins = append(bins, convertIntToBinary(6))
	bins = append(bins, convertIntToBinary(7))

	t.Log(len(bins))

	finalBinArray := make([]int, 33)
	for i := 32; i >= 0; i-- {
		isAllOne := true
		for j := 0; j < len(bins); j++ {
			if bins[j][i] == 0 {
				isAllOne = false
				break
			}
		}
		if isAllOne {
			finalBinArray[i] = 1
		}
	}

	t.Log(bins)
	t.Log(finalBinArray)

	t.Log(rangeBitwiseAnd(5, 7))
}
