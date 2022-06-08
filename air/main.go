package main

import (
	"fmt"
)

// //leetcode219
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	m := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		if _, ok := m[nums[i]]; ok {
// 			if i-m[nums[i]] <= k {
// 				return true
// 			}
// 		}
// 		m[nums[i]] = i
// 	}
// 	return false
// }

// func reverseList(head *ListNode) *ListNode {
// 	//需要记录当前链表的前后链表
// 	pre := nil
// 	cur := head
// 	//终止条件 遍历完链表,cur指向空
// 	for cur != nil {
// 		nex := cur //挪到这,防止cur空导致next异常
// 		//单个节点的操作
// 		//把当前节点的next设为pre,cur,nex,pre都往后挪1
// 		cur.Next = pre
// 		cur = nex
// 		nex = nex.Next
// 		pre = cur
// 	}
// 	//最后cur为空,pre是最后一个节点
// 	return pre
// }

// //leetcode144
// func preorderTraversal(root *TreeNode) (vals []int) {
// 	stack := []*TreeNode{}
// 	node := root
// 	for node != nil || len(stack) > 0 {
// 		for node != nil {
// 			vals = append(vals, node.Val)
// 			stack = append(stack, node)
// 			node = node.Left
// 		}
// 		node = stack[len(stack)-1].Right
// 		stack = stack[:len(stack)-1]
// 	}
// 	return
// }

//程序入口
func main() {
	str := "_"
	r := []rune(str)
	fmt.Println(r)
}
