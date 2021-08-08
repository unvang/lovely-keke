/*
最小栈题目
from: https://leetcode-cn.com/problems/min-stack/solution/zui-xiao-zhan-by-leetcode-solution/
*/

package week01

//solution 1
//one stack for loop get min
type minStack struct {
	top    *node
	length int
}
type node struct {
	value int
	prev  *node
}

/** initialize your data structure here. */
func NewStack() minStack {
	return minStack{nil, 0}
}

func (this *minStack) Push(val int) {
	n := &node{val, this.top}
	this.top = n
	this.length++
}

func (this *minStack) Pop() {
	n := this.top
	this.top = n.prev
	this.length--
}

func (this *minStack) Top() int {
	return this.top.value
}

func (this *minStack) GetMin() int {
	n := this.top
	minNum := this.top.value
	for n != nil {
		if n.value < minNum {
			minNum = n.value
		}
		n = n.prev
	}
	return minNum
}

//solution 2
//two stack,one store normal ,one store min val
type stack struct {
	top    *inode
	length int
}
type inode struct {
	val  int
	prev *inode
}
type MinStack struct {
	s    *stack
	minS *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var st MinStack
	if st.s == nil {
		st.s = &stack{nil, 0}
	}
	if st.minS == nil {
		st.minS = &stack{nil, 0}
	}
	return st
}

func (this *MinStack) Push(val int) {
	n1 := &inode{val, this.s.top}
	this.s.top = n1
	this.s.length++

	var curMinVal int
	if this.minS.length == 0 {
		curMinVal = val
	} else {
		curMinVal = this.GetMin()
	}
	if val < curMinVal {
		curMinVal = val
	}
	n2 := &inode{curMinVal, this.minS.top}
	this.minS.top = n2
	this.minS.length++
}

func (this *MinStack) Pop() {
	n1 := this.s.top
	this.s.top = n1.prev
	this.s.length--
	n2 := this.minS.top
	this.minS.top = n2.prev
	this.minS.length--
}

func (this *MinStack) Top() int {
	return this.s.top.val
}

func (this *MinStack) GetMin() int {
	return this.minS.top.val
}
