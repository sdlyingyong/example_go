package main

import (
	"fmt"
	"time"
)

type LLStack struct {
	linkedList linkedList
}

//初始化
func (Stack *LLStack) Init() LLStack {
	Stack.linkedList = linkedList{}
	Stack.linkedList.Init()
	return *Stack
}

//入栈
func (stack *LLStack) Push(val interface{}) {
	stack.linkedList.AddFirst(val)
}

//出栈
func (stack *LLStack) Pop() (val interface{}) {
	return stack.linkedList.Remove(0)
}

//栈顶元素
func (stack *LLStack) Peek() (val interface{}) {
	return stack.linkedList.Get(0)
}

//遍历打印
func (stack *LLStack) String() (str string) {
	cur := stack.linkedList.dummyHead.Next
	str += fmt.Sprintf("stack size: %d \n", stack.linkedList.GetSize())
	for cur != nil {
		str += fmt.Sprint(" -> ", cur.Val)
		cur = cur.Next
	}
	return
}

//测试栈
func TestLLStack() {
	//初始化栈
	var s1 LLStack
	s1.Init()
	for i := 0; i < 5; i++ {
		s1.Push(i)
		fmt.Println(s1.String())
	}
	s1.Pop()
	fmt.Println(s1.String())
}

//测试循环操作
func TestCountLLStack(count int) {
	var s1 LLStack
	s1.Init()
	//开始时间
	start := time.Now()
	for i := 0; i < count; i++ {
		//随机数
		s1.Push(i)
	}
	for i := 0; i < count; i++ {
		//随机数
		s1.Pop()
	}
	//结束时间
	end := time.Now()
	//打印时间差
	fmt.Println("TestLLStack time: ", end.Sub(start))
}

//测试循环操作
func TestCountStack(count int) {
	var s1 Stack
	s1 = Stack{}
	//开始时间
	start := time.Now()
	for i := 0; i < count; i++ {
		//随机数
		s1.Push(i)
	}
	for i := 0; i < count; i++ {
		//随机数
		s1.Pop()
	}
	//结束时间
	end := time.Now()
	//打印时间差
	fmt.Println(" TestStack time: ", end.Sub(start))
}

func main() {
	// TestLLStack()
	count := 10000000
	TestCountLLStack(count)
	TestCountStack(count)
}

//-----stack 代码-------

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

//-------链表代码----------

//node结构设计
type Node struct {
	Val  interface{}
	Next *Node
}

//初始化
func (node *Node) Node() {
	node.Val = nil
	node.Next = nil
}

//初始化
func (node *Node) NodeWithVal(val interface{}) {
	node.Val = val
	node.Next = nil
}

//初始化
func (node *Node) NodeWithValAndNext(val interface{}, next *Node) {
	node.Val = val
	node.Next = next
}

//链表
type linkedList struct {
	dummyHead *Node
	size      int
}

//初始化linkedlist
func (linkedList *linkedList) Init() linkedList {
	linkedList.dummyHead = &Node{}
	linkedList.size = 0
	return *linkedList
}

//返回链表是否为空
func (linkedList *linkedList) IsEmpty() (isEmpty bool) {
	isEmpty = linkedList.size == 0
	return
}

//获取链表长度
func (linkedList *linkedList) GetSize() (size int) {
	size = linkedList.size
	return
}

//在链表的指定位置添加节点
func (linkedList *linkedList) Add(index int, val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.size {
		panic("wrong index " + string(index))
	}
	//如果索引为0,直接用addHead添加
	if index == 0 {
		linkedList.AddFirst(val)
		return
	}
	//找到索引位置的前一个节点
	pre := linkedList.dummyHead.Next
	for i := 0; i < index-1; i++ {
		pre = pre.Next
	}
	//创建一个新节点
	newNode := &Node{val, pre.Next}
	//新节点的next指向pre.Next
	//把pre.Next指向新节点
	pre.Next = newNode

	//更新链表长度+1
	linkedList.size++
}

//在指定位置插入节点(虚拟节点方式)
func (linkedList *linkedList) AddWithDummy(index int, val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.size {
		panic("wrong index " + string(index))
	}
	//找到索引位置的前一个节点
	pre := linkedList.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	//创建一个新节点
	newNode := &Node{val, pre.Next}
	//新节点的next指向pre.Next
	//把pre.Next指向新节点
	pre.Next = newNode
	//更新链表长度+1
	linkedList.size++
}

//在链表头部添加新的元素
func (linkedList *linkedList) AddFirst(val interface{}) {
	linkedList.AddWithDummy(0, val)
	return
}

//链表末尾增加元素
func (linkedList *linkedList) AddLast(val interface{}) {
	linkedList.AddWithDummy(linkedList.size, val)
	return
}

//遍历打印链表
func (linkedList *linkedList) String() (str string) {
	cur := linkedList.dummyHead.Next
	str += fmt.Sprintf("linkedList size: %d \n", linkedList.GetSize())
	for cur != nil {
		str += fmt.Sprint(" -> ", cur.Val)
		cur = cur.Next
	}
	return
}

//链表的遍历操作
func (linkedList *linkedList) Get(index int) (val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.size {
		panic("wrong index " + string(index))
	}
	//执行几次当前就是索引位置
	cur := linkedList.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

//获得链表的第一个元素
func (linkedList *linkedList) GetFirst() (val interface{}) {
	return linkedList.Get(0)
}

//获得链表的最后一个元素
func (linkedList *linkedList) GetLast() (val interface{}) {
	return linkedList.Get(linkedList.size - 1)
}

//设置链表指定位置的元素
func (linkedList *linkedList) Set(index int, val interface{}) {
	//索引合法性检查
	if index <= 0 || index > linkedList.size {
		panic("wrong index " + string(index))
	}
	//执行几次当前就是索引位置
	cur := linkedList.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	//节点操作
	cur.Val = val
}

//删除链表指定位置的元素
func (linkedList *linkedList) Remove(index int) (val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.size {
		panic("wrong index " + string(index))
	}
	//找到前一个节点
	pre := linkedList.dummyHead
	//循环n次找到n索引前面一个的节点
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	delNode := pre.Next
	pre.Next = delNode.Next
	delNode.Next = nil
	linkedList.size--
	return delNode.Val
}

//删除收个元素
func (linkedList *linkedList) RemoveFirst() (val interface{}) {
	val = linkedList.Remove(0)
	return
}

//删除末尾元素
func (linkedList *linkedList) RemoveLast() (val interface{}) {
	val = linkedList.Remove(linkedList.size - 1)
	return
}

//测试链表增加
func TestAddFirst() {
	var l1 linkedList
	l1 = l1.Init()
	for i := 0; i < 5; i++ {
		l1.AddFirst(i)
		fmt.Println(l1.String())
	}
	l1.AddWithDummy(2, 666)
	fmt.Println(l1.String())
}
