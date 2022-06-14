package main

import "fmt"

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
	Size      int
}

//初始化linkedlist
func (linkedList *linkedList) Init() linkedList {
	linkedList.dummyHead = &Node{}
	linkedList.Size = 0
	fmt.Print("linkedList init: ")
	return *linkedList
}

//返回链表是否为空
func (linkedList *linkedList) IsEmpty() (isEmpty bool) {
	isEmpty = linkedList.Size == 0
	return
}

//获取链表长度
func (linkedList *linkedList) GetSize() (size int) {
	size = linkedList.Size
	return
}

//在链表的指定位置添加节点
func (linkedList *linkedList) Add(index int, val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.Size {
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
	linkedList.Size++
}

//在指定位置插入节点(虚拟节点方式)
func (linkedList *linkedList) AddWithDummy(index int, val interface{}) {
	//检查索引合法性
	if index < 0 || index > linkedList.Size {
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
	linkedList.Size++
}

//在链表头部添加新的元素
func (linkedList *linkedList) AddFirst(val interface{}) {
	linkedList.AddWithDummy(0, val)
	return
}

//链表末尾增加元素
func (linkedList *linkedList) AddLast(val interface{}) {
	linkedList.AddWithDummy(linkedList.Size, val)
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
	if index < 0 || index > linkedList.Size {
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
	return linkedList.Get(linkedList.Size - 1)
}

//设置链表指定位置的元素
func (linkedList *linkedList) Set(index int, val interface{}) {
	//索引合法性检查
	if index <= 0 || index > linkedList.Size {
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
	if index < 0 || index > linkedList.Size {
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
	linkedList.Size--
	return delNode.Val
}

//删除收个元素
func (linkedList *linkedList) RemoveFirst() (val interface{}) {
	val = linkedList.Remove(0)
	return
}

//删除末尾元素
func (linkedList *linkedList) RemoveLast() (val interface{}) {
	val = linkedList.Remove(linkedList.Size - 1)
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

func main() {
	TestAddFirst()
}
