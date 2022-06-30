package main

import "fmt"

func main() {
	//吃完耗费的时间
	testEatingTime()
}

func testEatingTime() {
	fmt.Println("minEatingSpeed([]int{3,6,7,11},8) ", minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

//找到h小时内吃完所需要的最小k
func minEatingSpeed(piles []int, h int) (k int) {
	//范围 [1...1e9]
	left, right := 1, int(1e9)
	for left < right {
		mid := left + (right-left)/2
		if eatingTime(piles, mid) <= h {
			//范围 [left...mid]
			right = mid
		} else {
			//范围[mid+1...right]
			left = mid + 1
		}
	}
	k = left
	return
}

//每一堆吃完花掉的间
func eatingTime(piles []int, k int) (useHour int) {
	for _, pie := range piles {
		if pie%k == 0 {
			useHour += pie / k
		} else {
			useHour += pie/k + 1
		}
	}
	return
}
