package main

import (
	"fmt"
)

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

//---------------链表定义-----------------

//链表
type linkedList struct {
	dummyHead *Node
	size      int
}

//初始化linkedlist
func (linkedList *linkedList) Init() linkedList {
	linkedList.dummyHead = &Node{}
	linkedList.size = 0
	fmt.Print("linkedList init: ")
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
	//处理函数边界条件
	if index < 0 || index > linkedList.size {
		panic("Add failed. Illegal index.")
	}
	linkedList.addRecursion(index, val, 0, linkedList.dummyHead)
	//节点操作
	linkedList.size++
	return
}

//链表添加节点的递归处理
//使用虚拟头结点
//递归是小问题的累加
func (linkedList *linkedList) addRecursion(index int, val interface{}, depth int, head *Node) {
	//base问题 在index node后插入新的节点
	if depth == index {
		//处理节点操作:插入新节点
		newNode := &Node{val, nil}
		newNode.Next = head.Next
		//先操作,最后断开连接
		head.Next = newNode
		return
	}
	linkedList.addRecursion(index, val, depth+1, head.Next)
	return
}

//在链表头部添加新的元素
func (linkedList *linkedList) AddFirst(val interface{}) {
	linkedList.Add(0, val)
	return
}

//链表末尾增加元素
func (linkedList *linkedList) AddLast(val interface{}) {
	linkedList.Add(linkedList.size, val)
	return
}

//遍历打印链表
func (linkedList *linkedList) String() (str string) {
	fmt.Println("linkedList size: ", linkedList.size)
	str = StrRecursion(linkedList.dummyHead.Next)
	return
}

//打印链表的递归函数-递归版
func StrRecursion(head *Node) (str string) {
	//base问题
	if head.Next == nil {
		return fmt.Sprintf("-> %v", head.Val)
	}
	//子函数的结果
	ret := StrRecursion(head.Next)
	str += fmt.Sprintf("-> %v %s", head.Val, ret)
	return
}

//获取指定位置的val-递归版
func (linkedList *linkedList) Get(index int) (val interface{}) {
	//合法性检查
	if index < 0 || index >= linkedList.size {
		panic("Get failed. Illegal index.")
	}
	val = linkedList.getRecursion(index, 0, linkedList.dummyHead.Next)
	return
}

//get操作的递归
func (linkedList *linkedList) getRecursion(index, depth int, head *Node) (val interface{}) {
	//最小问题的解决
	if depth == index {
		return head.Val
	}
	//递归
	val = linkedList.getRecursion(index, depth+1, head.Next)
	return
}

//获得链表的第一个元素
func (linkedList *linkedList) GetFirst() (val interface{}) {
	return linkedList.Get(0)
}

//获得链表的最后一个元素
func (linkedList *linkedList) GetLast() (val interface{}) {
	return linkedList.Get(linkedList.size - 1)
}

//设置链表指定位置的元素-递归版
func (linkedList *linkedList) Set(index int, val interface{}) {
	//合法性检查
	if index < 0 || index >= linkedList.size {
		panic("Set failed. Illegal index.")
	}
	linkedList.setRecursion(index, val, 0, linkedList.dummyHead.Next)
	return
}

//设置链表指定位置的元素递归函数
func (linkedList *linkedList) setRecursion(index int, val interface{}, depth int, head *Node) {
	//最小问题的解决
	if depth == index-1 {
		replaceNode := head.Next
		newNode := &Node{val, nil}
		newNode.Next = replaceNode.Next
		head.Next = newNode
		return
	}
	//递归
	linkedList.setRecursion(index, val, depth+1, head.Next)
	return
}

//删除链表指定位置的元素-递归版
func (linkedList *linkedList) Remove(index int) (val interface{}) {
	//合法性检查
	if index < 0 || index > linkedList.size-1 {
		panic("Remove failed. Illegal index.")
	}
	//递归函数
	val = linkedList.removeRecursion(index, 0, linkedList.dummyHead)
	return
}

//删除元素递归版
func (linkedList *linkedList) removeRecursion(index int, depth int, head *Node) (val interface{}) {
	//最小问题的解决
	if depth == index {
		delNode := head.Next
		head.Next = delNode.Next
		delNode.Next = nil
		val = delNode.Val
		linkedList.size--
		return
	}
	//递归,当做调用子函数
	val = linkedList.removeRecursion(index, depth+1, head.Next)
	return
}

//删除首个元素
//时间复杂度 O(1)
func (linkedList *linkedList) RemoveFirst() (val interface{}) {
	val = linkedList.Remove(0)
	return
}

//删除末尾元素
//时间复杂度 O(n)
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
	l1.Add(2, 666)
	fmt.Println(l1.String())
	fmt.Println("l1.GetFirst() :", l1.GetFirst())
	fmt.Println("l1.GetLast() :", l1.GetLast())
	fmt.Println("l1.Get(l1.size-1) : ", l1.Get(l1.size-1))
	l1.Set(2, 665)
	fmt.Println(l1.String())
}

func TestRemove() {
	var l1 linkedList
	l1 = l1.Init()
	for i := 0; i < 5; i++ {
		l1.AddFirst(i)
	}
	fmt.Println(l1.String())
	fmt.Println("l1.Remove(0) :", l1.Remove(0))
	fmt.Println(l1.String())
	fmt.Println("l1.RemoveFirst() :", l1.RemoveFirst())
	fmt.Println(l1.String())
	fmt.Println("l1.RemoveLast() :", l1.RemoveLast())
	fmt.Println(l1.String())
}

func main() {
	// TestAddFirst()
	TestRemove()
}
