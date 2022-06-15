package main

import (
	"fmt"
	"time"
)

//定义队列,底层基于链表实现
type LLQueue struct {
	head *Node
	tail *Node
	size int
}

//初始化队列
func (queue *LLQueue) Init() LLQueue {
	queue.size = 0
	return *queue
}

//队列是否为空
func (queue *LLQueue) IsEmpty() (isEmpty bool) {
	isEmpty = queue.size == 0
	return
}

//队列长度
func (queue *LLQueue) GetLen() (len int) {
	len = queue.size
	return
}

//入队
func (queue *LLQueue) Enqueue(val interface{}) {
	//队列为空
	if queue.size == 0 {
		//初始化
		newNode := &Node{val, nil}
		queue.head = newNode
		queue.tail = newNode
		queue.size++
		return
	}
	//队列有值
	newNode := &Node{val, nil}
	queue.tail.Next = newNode
	queue.tail = newNode
	queue.size++
}

//出列
func (queue *LLQueue) Dequeue() (val interface{}) {
	//如果队列为空
	if queue.size == 0 {
		return
	}
	//队列不为空
	val = queue.head.Val
	queue.head = queue.head.Next
	queue.size--
	return
}

//查看队列首个元素
func (queue *LLQueue) GetFront() (val interface{}) {
	//如果队列为空
	if queue.size == 0 {
		return
	}
	//队列不为空
	val = queue.head.Val
	return
}

//打印队列
func (queue *LLQueue) Print() {
	ret := fmt.Sprintf("Queue len = %d  \n front ", queue.GetLen())
	//遍历队列长度次
	cur := queue.head
	for i := 0; i < queue.size; i++ {
		ret += fmt.Sprintf("%d <- ", cur.Val)
		cur = cur.Next
	}
	fmt.Println(ret)
}

//基础测试
func testQueue() {
	mq := LLQueue{}
	mq.Init()
	for i := 0; i < 10; i++ {
		mq.Enqueue(i)
		mq.Print()
		if i%3 == 2 {
			mq.Dequeue()
			mq.Print()
		}
	}
}

//性能测试 linkedList queue
// TestCountLLQueue time:  1.2327185s
func TestCountLLQueue(count int) {
	var mq LLQueue
	mq.Init()
	//开始时间
	start := time.Now()
	for i := 0; i < count; i++ {
		mq.Enqueue(i)
		if i%3 == 2 {
			mq.Dequeue()
		}
	}
	//结束时间
	end := time.Now()
	//打印时间差
	fmt.Println("TestCountLLQueue time: ", end.Sub(start))
}

//性能测试 loop queue
// TestCountLPQueue time:  207.9507ms
//测试循环操作
func TestCountLPQueue(count int) {
	mq := Constructor()
	//开始时间
	start := time.Now()
	for i := 0; i < count; i++ {
		mq.Push(i)
		if i%3 == 2 {
			mq.Pop()
		}
	}
	//结束时间
	end := time.Now()
	//打印时间差
	fmt.Println("TestCountLPQueue time: ", end.Sub(start))
}

func main() {
	count := 10000000
	// testQueue()
	TestCountLLQueue(count)
	TestCountLPQueue(count)
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
//时间复杂度 O(2/n)
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

//--------循环队列--------

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
