package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	TestReverseList()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) (retHead *ListNode) {
	//过滤空和异常
	if head == nil || head.Next == nil {
		retHead = head
		return
	}
	retHead = reverseListHelper(nil, head, head.Next)
	return
}

//反装链表的辅助函数
func reverseListHelper(pre, cur, next *ListNode) (retHead *ListNode) {
	//base问题的解决方案
	if next == nil {
		cur.Next = pre
		return cur
	}
	//当做调用子函数
	retHead = reverseListHelper(cur, next, next.Next)
	//节点操作
	cur.Next = pre
	return
}

//测试反转链表
func TestReverseList() {
	//input
	ll := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(ll.String())
	//output
	fmt.Println(reverseList(ll))
}

//string
func (l *ListNode) String() string {
	var ret string
	for l != nil {
		ret += fmt.Sprintf("%d->", l.Val)
		l = l.Next
	}
	return ret
}
