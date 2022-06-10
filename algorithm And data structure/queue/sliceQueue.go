package main

import "fmt"

func main() {
	//leetCode232. 用栈实现队列
	// implement-queue-using-stacks
	TestMyQueue()
}

//测试MyQueue
func TestMyQueue() {
	mq := Constructor()
	mq.Push(1)
	mq.Push(2)
	fmt.Println("mq.Peek(): ", mq.Peek())
	fmt.Println("mq.Pop(): ", mq.Pop())
	fmt.Println("mq.Empty(): ", mq.Empty())
}

//定义队列
type MyQueue struct {
	out, in Stack
}

//初始化
func Constructor() MyQueue {
	return MyQueue{}
}

//末尾插入元素
func (mq *MyQueue) Push(x int) {
	mq.in.Push(x)
}

//队首取出元素
func (mq *MyQueue) Pop() (value int) {
	if mq.Empty() {
		return
	}
	//out为空,把in倒入out,再查看out的队首元素
	if mq.out.IsEmpty() {
		mq.Move(&mq.in, &mq.out)
		return mq.out.Pop()
	} else {
		//out不为空
		//in为空,就直接返回out顶端
		if !mq.in.IsEmpty() {
			return mq.out.Pop()
		} else {
			//in不为空,把out倒入in,再倒回out,再返回out顶端
			mq.Move(&mq.out, &mq.in)
			mq.Move(&mq.in, &mq.out)
			return mq.out.Pop()
		}
	}
}

//遍历一个stack,导入到另一个stack
func (mq *MyQueue) Move(s1, s2 *Stack) {
	//如果s1为空,则不需要做任何操作
	if s1.IsEmpty() {
		return
	}
	//遍历,把所有s1的元素倒进s2
	for !s1.IsEmpty() {
		s2.Push(s1.Pop())
	}
	//搬家操作没搬掉,这里传递的是复制的变量
	// fmt.Println("S1", s1, "S2", s2)
	// fmt.Println("mq.in", mq.in, "mq.out", mq.out)
}

//查看队首元素
func (mq *MyQueue) Peek() (value int) {
	if mq.Empty() {
		return
	}
	//out为空,把in倒入out,再查看out的队首元素
	if mq.out.IsEmpty() {
		mq.Move(&mq.in, &mq.out)
		return mq.out.Peek()
	} else {
		//out不为空
		//in为空,就直接返回out顶端
		if !mq.in.IsEmpty() {
			return mq.out.Peek()
		} else {
			//in不为空,把out倒入in,再倒回out,再返回out顶端
			mq.Move(&mq.out, &mq.in)
			mq.Move(&mq.in, &mq.out)
			return mq.out.Peek()
		}
	}
}

//是否为空
func (mq *MyQueue) Empty() (isEmpty bool) {
	isEmpty = mq.in.IsEmpty() && mq.out.IsEmpty()
	return
}

//-------stack-----

type Stack struct {
	data []int
}

//stack支持的方法 入栈 出栈 栈顶元素  站的size 栈是否为空

//入栈
func (s *Stack) Push(input int) {
	s.data = append(s.data, input)
}

//出栈
func (s *Stack) Pop() (value int) {
	if s.IsEmpty() {
		return
	}
	value = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1] //[0...end) go的切片语法是不包含后面结尾元素的
	return
}

//查看栈顶元素
func (s *Stack) Peek() (value int) {
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
