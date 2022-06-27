package main

import "fmt"

func main() {
	//测试二分排序
	testBinarySearch()
}

//测试二分搜索
func testBinarySearch() {
	input := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println("binarySearch(input, 2)", search2(input, 13))
}

//二分搜索,返回索引位置
//不存在则返回-1
func search(nums []int, target int) int {
	return binarySearchHelper(nums, target, 0, len(nums)-1)
}

//二分搜索递归辅助函数
func binarySearchHelper(nums []int, target, left, right int) (index int) {
	if left > right {
		return -1
	}
	//1.最小情况的解决方案
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}
	//2.当做调用子函数
	if target > nums[mid] {
		//搜索范围 (mid...right]
		return binarySearchHelper(nums, target, mid+1, right)
	}
	//搜索范围 [left...mid)
	return binarySearchHelper(nums, target, left, mid-1)
}

//非递归实现的二分查找
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//范围 [mid+1...right]
			left = mid + 1
		} else {
			//范围 [left...mid-1]
			right = mid - 1
		}
	}
	return -1
}
