package main

import "fmt"

type stack struct {
	data []interface{}
}

//stack支持的方法 入栈 出栈 栈顶元素  站的size 栈是否为空

//入栈
func (s *stack) Push(input interface{}) {
	s.data = append(s.data, input)
}

//出栈
func (s *stack) Pop() (value interface{}) {
	value = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1] //[0...end) go的切片语法是不包含后面结尾元素的
	return
}

//查看栈顶元素
func (s *stack) Peek() (value interface{}) {
	value = s.data[len(s.data)-1]
	return
}

//栈的size
func (s *stack) Size() (size int) {
	size = len(s.data)
	return
}

//栈是否为空
func (s *stack) IsEmpty() (isEmpty bool) {
	isEmpty = len(s.data) == 0
	return
}

func main() {
	//测试栈push
	s := &stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	//测试栈顶元素
	fmt.Println("s.Peek() :", s.Peek())
	//测试pop
	fmt.Println("s.Pop() :", s.Pop())
	//测试栈的size
	fmt.Println("s.Size() :", s.Size())
	//测试栈是否为空
	fmt.Println("s.IsEmpty() :", s.IsEmpty())
}
