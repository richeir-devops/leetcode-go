package p000

import (
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	//1、负数肯定不是回文数
	if x < 0 {
		return false
	}

	//2、从整数转换成字符串再进行比对
	//对普通整数先转 int64 ，然后再使用 strconv.FormatInt(b, 10) 进行转换
	x64 := int64(x)
	strx64 := strconv.FormatInt(x64, 10)
	runeList := []rune(strx64)

	//3、从两边开始对比，判断是否完全满足回文数条件
	lenList := len(runeList)
	for i := 0; i < lenList/2; i++ {
		if runeList[i] != runeList[lenList-i-1] {
			return false
		}
	}
	return true
}

func Test_009_1(t *testing.T) {
	t.Log("009")
	a := 1234543212
	b := int64(a)
	c := []rune(strconv.FormatInt(b, 10))
	t.Log(c)
	t.Log(isPalindrome(a))
}
