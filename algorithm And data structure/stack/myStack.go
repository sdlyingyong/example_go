package main

import "fmt"

type Stack struct {
	data []interface{}
}

//stack支持的方法 入栈 出栈 栈顶元素  站的size 栈是否为空

//入栈
func (s *Stack) Push(input interface{}) {
	s.data = append(s.data, input)
}

//出栈
func (s *Stack) Pop() (value interface{}) {
	if s.IsEmpty() {
		return nil
	}
	value = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1] //[0...end) go的切片语法是不包含后面结尾元素的
	return
}

//查看栈顶元素
func (s *Stack) Peek() (value interface{}) {
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

//使用stack完成LeetCode 20. Valid Parentheses
func LeetCode20(s string) (isValid bool) {
	//使用stack的结构来解决问题
	stk := &Stack{}
	//遍历字符串
	for _, v := range s {
		//如果遇到左括号,就入栈
		if v == '(' || v == '[' || v == '{' {
			stk.Push(v)
		} else if v == ')' || v == ']' || v == '}' {
			//右括号就出栈对比
			left := stk.Pop()
			if left == nil {
				return false
			}
			switch left {
			case '(':
				if v != ')' {
					return false
				}
			case '[':
				if v != ']' {
					return false
				}
			case '{':
				if v != '}' {
					return false
				}
			}
		} else {
			panic("非法字符")
		}
	}
	//如果遍历完,没有未闭合的标签(栈中镁元素),就判定为已闭合
	if stk.IsEmpty() {
		isValid = true
	}
	return
}
func main() {
	//测试栈push
	s := &Stack{}
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

	// ret := LeetCode20("({{{{}}}))")
	// ret := LeetCode20("(]")
	ret := LeetCode20("(]")
	fmt.Println("ret :", ret)
}
