package main

import "fmt"

//remove-linked-list-elements
func main() {
	// TestRemoveElements()
	TestRemoveElementsDH()
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

//测试
func TestRemoveElementsDH() {
	//input
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head = removeElementsDH(head, 6)
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
