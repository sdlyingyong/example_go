package main

import (
	"fmt"
)

func main() {
	// //测试二分排序
	// testBinarySearch()

	// //测试大于target的最小值
	// testSearchBigger()

	// testCeil()

	// //testLower
	// testLower()

	//
	// testLowerFloor()

	testMyUpper()
}

//测试二分搜索
func testBinarySearch() {
	input := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println("binarySearch(input, 2)", search3(input, 13))
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

//非递归实现的二分查找
func search3(nums []int, target int) int {
	left, right := 0, len(nums)
	//范围 [left...right)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//范围[mid+1...right]
			left = mid + 1
		} else {
			//包含了等于
			//范围[left...mid]
			right = mid
		}
	}
	return -1
}

//测试数组查找大于的元素
func testSearchBigger() {
	input := []int{1, 1, 3, 3, 5, 5}
	for i := 0; i <= 6; i++ {
		fmt.Printf("SearchBigger(input, %d) : %d \n", i, SearchBigger(input, i))
	}
}

//搜索有序数组中大于60的元素位置
func SearchBigger(arr []int, target int) int {
	//设定左右范围
	left, right := 0, len(arr)-1
	//包含只有一个元素的情况
	for left <= right {
		mid := left + (right-left)/2
		//返回大于他的第一个值
		if arr[mid] > target {
			return mid
		}
		//二分查找缩小范围
		if arr[mid] > target {
			//范围[left...mid]
			right = mid
		} else {
			//包含arr[mid] == target情况
			//范围[mid+1...right]
			left = mid + 1
		}
	}
	return len(arr)
}

//测试ceil方法
func testCeil() {
	input := []int{1, 1, 3, 3, 5, 5}
	for i := 0; i <= 6; i++ {
		fmt.Printf("Ceil(input, %d) : %d \n", i, Ceil(input, i))
	}
}

//搜索数组中元素的最大索引
//如果不存在的话,返回大于这个元素的最小值索引
func Ceil(arr []int, target int) (index int) {
	ret := SearchBigger(arr, target)
	if ret-1 >= 0 && arr[ret-1] == target {
		return ret - 1
	}
	return ret
}

//lower_ceil
//如果数组存在元素,返回最小索引
//如果没有,就返回元素的upper
//todo
func lower_ceil(arr []int) (index int) {
	return
}

//测试Lower方法
func testLower() {
	input := []int{1, 1, 3, 3, 5, 5}
	for i := 0; i <= 6; i++ {
		fmt.Printf("Lower(input, %d) : %d \n", i, Lower(input, i))
	}
}

//mid计算中,如果有0.5会被设置为0,会用下取整方式处理
func Lower(arr []int, target int) (index int) {
	//设定左右边界
	left, right := -1, len(arr)-1
	//范围 [left...right]
	for left < right {
		//特殊案例[0,1] mid=0.5但是int类型设置为0了
		mid := left + (right-left+1)/2
		// fmt.Println("left", left, "right ", right, "mid ", mid)
		//二分查找
		if arr[mid] < target {
			//范围 [mid...right]
			left = mid
		} else {
			//范围 [left...mid-1]
			right = mid - 1
		}
	}
	return left
}

func testLowerFloor() {
	input := []int{1, 1, 3, 3, 5, 5}
	for i := 0; i <= 6; i++ {
		fmt.Printf("lowerFloor(input, %d) : %d \n", i, lowerFloor(input, i))
	}
	return
}

//下层楼
func lowerFloor(arr []int, target int) (index int) {
	left, right := -1, len(arr)-1
	for left < right { //left right 不为同一个数
		mid := left + (right-left+1)/2 //采用向上取整的方式
		if arr[mid] > target {
			//范围 [left...mid-1]
			right = mid - 1
		} else {
			//范围 [mid...right]
			left = mid
		}
	}
	return left - 1
}

func myUpper(arr []int, target int) (index int) {
	//涉及到
	left, right := 0, len(arr)
	for left < right {
		//left right 不相同
		mid := left + (right-left)/2
		if arr[mid] < target {
			//范围 [mid+1...right]
			left = mid + 1
		} else {
			//范围 [left...mid]
			right = mid
		}
	}
	return left
}

func testMyUpper() {
	input := []int{50, 56, 65, 69, 72, 89, 96, 99}
	fmt.Println("myUpper(input, 60) : ", myUpper(input, 60))
	//预期是4
}
