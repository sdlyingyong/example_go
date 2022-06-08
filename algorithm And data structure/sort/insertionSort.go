package main

import "fmt"

func main() {
	testInsertSort()
}

//插入排序法
func insertionSort(arr []int) []int {
	//一范围 从头开始,每次挑选一个元素,current = arr[i]
	for i := range arr {
		current := arr[i]
		//二范围 从arr[i-1]开始,直到arr[0],进行对比 index j
		preIndex := i - 1
		//如果arr[j]>current,就把这个元素位置往后一位放, arr[j+1] = arr[j]
		//随后j--,直到为0或者arr[j]<current,此时把current放在这里 arr[j] = current
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	//返回数组
	return arr
}

func testInsertSort() {
	input := []int{1, 2, 3, 4, 0}
	ret := insertionSort(input)
	fmt.Print(ret)
}
