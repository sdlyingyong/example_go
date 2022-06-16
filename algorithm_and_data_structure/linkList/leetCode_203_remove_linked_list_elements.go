package main

import (
	"fmt"
)

//remove-linked-list-elements
func main() {
	// TestRemoveElements()
	// TestRemoveElementsDH()        //dummy head 方式解决
	// TestRemoveElementsRecursion() //递归方式解决 recursion
	TestSumRecursion()

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//链表删除节点
func removeElements(head *ListNode, val int) *ListNode {
	//遍历头部节点检查是否要删除
	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}
	//如果头结点为空,直接返回
	if head == nil {
		return nil
	}
	pre := head
	//找到最后一个节点,检查pre后一个是否要删除
	for pre.Next != nil {
		if pre.Next.Val == val {
			//删除节点操作
			delNode := pre.Next
			pre.Next = delNode.Next
			delNode.Next = nil
		} else {
			pre = pre.Next
		}
	}
	return head
}

//测试
func TestRemoveElements() {
	//input
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head = removeElements(head, 6)
	//output
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
}

//链表删除节点 使用递归方式
//增加调试打印
func removeElementsRecursion(head *ListNode, val int, depth int) *ListNode {
	depthString := generateDepthString(depth)
	fmt.Println(depthString)
	fmt.Println("Call : remove ", val, " in ", head.String())
	//问题最小情况的处理
	if head == nil {
		fmt.Print(depthString)
		fmt.Println("Return : nil")
		return nil
	}
	//当做调用子函数,返回后续删除完指定元素后的一条链表
	head.Next = removeElementsRecursion(head.Next, val, depth+1)
	fmt.Println(depthString)
	fmt.Println("After  remove ", val, " : ", head.Next.String())
	//当前节点的操作
	if head.Val == val {
		fmt.Println(depthString)
		fmt.Println("Return   ", " : ", head.Next.String())
		//当前要删除,就返回下个节点开始的一条链表
		return head.Next
	}
	fmt.Println(depthString)
	fmt.Println("Return   ", " : ", head.String())
	//不属于要删除的,就返回当前节点后续的一条链表
	return head
}

//node to string
func (head *ListNode) String() string {
	str := ""
	for head != nil {
		str += fmt.Sprintf("%d -> ", head.Val)
		head = head.Next
	}
	return str
}

//深度字符串
func generateDepthString(depth int) (str string) {
	for i := 0; i < depth; i++ {
		str += "--"
	}
	return str
}

//测试
func TestRemoveElementsRecursion() {
	//input
	input := []int{1, 2, 6, 3, 4, 5, 6}
	inputLL := NewListNode(input)
	head := removeElementsRecursion(inputLL, 6, 0)
	//output
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
}

//链表删除节点(使用虚拟头结点)
func removeElementsDH(head *ListNode, val int) *ListNode {
	//处理空情况
	if head == nil {
		return nil
	}
	//设置虚拟头结点
	dummyHead := &ListNode{Val: 0, Next: head}
	pre := dummyHead
	//遍历每个节点,pre指向待删除值的前一个节点
	for pre.Next != nil {
		if pre.Next.Val == val {
			//删除节点操作
			delNode := pre.Next
			pre.Next = delNode.Next
			delNode.Next = nil
		} else {
			pre = pre.Next
		}
	}
	//隐藏虚拟节点
	return dummyHead.Next
}

//测试
func TestRemoveElementsDH() {
	//input
	input := []int{1, 2, 6, 3, 4, 5, 6}
	inputLL := NewListNode(input)
	head := removeElementsDH(inputLL, 6)
	//output
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
}

//链表节点的构造函数
//传入数组
func NewListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

//递归求和
func sumRecursion(arr []int, depth int) int {
	depthStr := generateDepthString(depth)
	//问题最小情况的处理
	fmt.Println(depthStr)
	fmt.Println("Call : sum ", arr)
	if len(arr) == 0 {
		fmt.Println(depthStr)
		fmt.Println("Return : 0")
		return 0
	}
	//当做调用子函数,拿到后续数组的和
	ret := sumRecursion(arr[1:], depth+1)
	fmt.Println(depthStr)
	fmt.Println("After  sum : ", arr[1:], " ret : ", ret)
	//当前节点的操作
	return ret + arr[0]
}

//测试递归求sum
func TestSumRecursion() {
	//input
	input := []int{1, 2, 3, 4}
	// input = generateArray(1e2) //10^2 科学计数法表示为 1e2
	//output
	fmt.Println(sumRecursion(input, 0))
}

//生成数组
func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}
