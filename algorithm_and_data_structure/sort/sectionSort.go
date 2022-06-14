package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testSectionSort()
}

//基础测试
func testSectionSort2() {
	input := []int{1, 2, 3, 4, 0}
	ret := sectionSort(input)
	checkSort(ret)
	fmt.Println(ret)
}

//选择排序
//传入一个数组,返回按照从小到大排列的
func sectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		tmpIdx := i
		for j := i + 1; j < length; j++ {
			//每个节点把最小的放在前面
			if arr[j] < arr[tmpIdx] {
				tmpIdx = j
			}
		}
		//交换
		arr[i], arr[tmpIdx] = arr[tmpIdx], arr[i]
	}
	return arr
}

//选择排序2
//从大到小排序
//todo 目前不正确
func sectionSort2(arr []int) (ret []int) {
	//每次选择把最大的放在最后
	length := len(arr)
	for i := length - 1; i > 1; i-- { //范围 [1...i]
		tmpIdx := i
		for j := i - 1; j > 0; j-- { //范围 [0...i-1]
			if arr[j] > arr[tmpIdx] {
				tmpIdx = j
			}
		}
		//交换
		fmt.Println("tmpIdx ", tmpIdx)
		arr[tmpIdx], arr[i] = arr[i], arr[tmpIdx]
	}
	return arr
}

//测试选择排序
func testSectionSort() {
	//10的四次方
	testnum := 10000
	input := generateArrayRandom(testnum)
	//运行时间
	start := time.Now()
	ret := sectionSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(ret))
	//结束时间
	end := time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)

	//两倍数组
	input = generateArrayRandom(testnum * 2)
	//运行时间
	start = time.Now()
	ret = sectionSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end = time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)

	//四倍数组
	input = generateArrayRandom(testnum * 2 * 2)
	//运行时间
	start = time.Now()
	ret = sectionSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end = time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)

	//八倍数组
	input = generateArrayRandom(testnum * 2 * 2 * 2)
	//运行时间
	start = time.Now()
	ret = sectionSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end = time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)
}

//检查排序数组
func checkSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("error")
			return
		}
	}
	fmt.Println("ok")
}

//生成数组长度n,乱序的数组
func generateArrayRandom(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
