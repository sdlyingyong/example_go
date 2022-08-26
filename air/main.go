package main

import (
	"fmt"
	"time"
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

	//时间
	now := time.Now()

	//后续时间
	// 	一、 复习点的确定（根据艾宾浩斯记忆曲线制定）：
	// 　　1. 第一个记忆周期：5分钟
	// 　　2. 第二个记忆周期：30分钟
	// 　　3. 第三个记忆周期：12小时
	// 　　4. 第四个记忆周期：1天
	// 　　5. 第五个记忆周期：2天
	// 　　6. 第六个记忆周期：4天
	// 　　7. 第七个记忆周期：7天
	// 　　8. 第八个记忆周期：15天
	timeArr := []time.Duration{
		time.Minute * 5,
		time.Minute * 30,
		time.Hour * 12,
		time.Hour * 24,
		time.Hour * 24 * 2,
		time.Hour * 24 * 4,
		time.Hour * 24 * 7,
		time.Hour * 24 * 15,
	}

	for _, v := range timeArr {
		fmt.Println(now.Add(v).Format("2006/01/02 15:04"))
	}

}
