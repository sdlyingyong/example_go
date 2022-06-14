package main

import (
	"fmt"
	"math"
)

func main() {
	//测试最小栈
	ms := Constructor()
	ms.Push(1)
	ms.Pop()
	fmt.Println("ms.Top() :", ms.Top())
	fmt.Println("ms.GetMin() : ", ms.GetMin())
}

type MinStack struct {
	data Stack //保存数字
	min  Stack //保存最小数
}

//初始化
func Constructor() (newStack MinStack) {
	newStack = MinStack{}
	//将最大值int64压入min stack
	newStack.min.Push(math.MaxInt64)
	return
}

//进入栈
func (ms *MinStack) Push(val int) {
	//压入数据栈
	ms.data.Push(int64(val))
	//压入最小栈
	if int64(val) < ms.min.Peek() {
		//如果比当前最小值小,则更新最小值
		ms.min.Push(int64(val))
	} else {
		//如果不小于当前最小值,压入一个当前最小值,
		//此时栈到下个最小值之间都是当前最小值
		ms.min.Push(ms.min.Peek())
	}
}

//取出栈顶端元素
func (ms *MinStack) Pop() {
	//取出数据栈顶元素
	ms.data.Pop()
	//同步取出最小栈顶元素
	ms.min.Pop()
	return
}

//展示栈顶元素
func (ms *MinStack) Top() int {
	return int(ms.data.Peek())
}

//从最小栈获取
func (ms *MinStack) GetMin() (ret int64) {
	ret = ms.min.Peek()
	if ret == math.MaxInt64 {
		fmt.Println("min stack is empty")
		return
	}
	return
}

//----------自己实现的stack----------------
type Stack struct {
	data []int64
}

//stack支持的方法 入栈 出栈 栈顶元素  站的size 栈是否为空

//入栈
func (s *Stack) Push(input int64) {
	s.data = append(s.data, input)
}

//出栈
func (s *Stack) Pop() (value int64) {
	if s.IsEmpty() {
		return
	}
	value = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1] //[0...end) go的切片语法是不包含后面结尾元素的
	return
}

//查看栈顶元素
func (s *Stack) Peek() (value int64) {
	if s.IsEmpty() {
		return
	}
	value = s.data[len(s.data)-1]
	return
}

//栈的size
func (s *Stack) Size() (size int) {
	size = len(s.data)
	return
}

//栈是否为空
func (s *Stack) IsEmpty() (isEmpty bool) {
	isEmpty = len(s.data) == 0
	return
}
