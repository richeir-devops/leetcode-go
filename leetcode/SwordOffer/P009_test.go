package SwordOffer

import (
	"container/list"
	"testing"
)

///////////////////////////////////////////
// 解法1（普通逻辑）：
// 1.使用标准库 container/list
// 2.list1 模拟队列的队尾，一个一个进
// 3.从队首删除元素要把 list1 倒着压入 list2，这样 list2 的弹展操作就相当于删除了队首元素

type CQueue struct {
	list1 *list.List
	list2 *list.List
}

func Constructor() CQueue {
	return CQueue{list1: &list.List{}, list2: &list.List{}}
}

func (this *CQueue) AppendTail(value int) {
	this.list1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.list2.Len() != 0 {
		return deleteListLastElement(this.list2)
	}

	if this.list1.Len() == 0 {
		return -1
	}

	for this.list1.Len() != 0 {
		this.list2.PushBack(deleteListLastElement(this.list1))
	}

	return deleteListLastElement(this.list2)
}

func deleteListLastElement(target *list.List) int {
	returnValue := target.Back()
	target.Remove(returnValue)
	return returnValue.Value.(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func Test_009_01(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(3)
	obj.AppendTail(4)
	d1 := obj.DeleteHead()
	d1 = obj.DeleteHead()
	d1 = obj.DeleteHead()

	t.Log(d1)
}
