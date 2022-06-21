package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testMerge()
	testMergeSort()
	testTimeO()
	TestTimeO2()
}

func testSort() {
	arr := generateArrayRandom(10)
	fmt.Println("arr :", arr)
	arr = mergeSort(arr)
	fmt.Println("ret :", arr)
	checkSort(arr)
}

//合并排序
func mergeSort(arr []int) (ret []int) {
	len := len(arr)
	if len < 2 {
		ret = arr
		return
	}
	mid := len / 2
	left := arr[0:mid] //不包含mid
	right := arr[mid:] //从mid到最后
	return merge(mergeSort(left), mergeSort(right))
}

//合并
func merge(left, right []int) (ret []int) {
	//两个数组都没取完,就继续取
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			ret = append(ret, left[0])
			left = left[1:] //从索引1开始到结尾
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	//left取完,就把right剩余的全部放入ret
	if len(left) == 0 {
		ret = append(ret, right...)
	}
	//right取完,就把left剩余的全部放入ret
	if len(right) == 0 {
		ret = append(ret, left...)
	}
	return ret
}

//testMergeSort
func testMergeSort() {
	arr := []int{6, 8, 3, 4, 9, 2, 7, 5, 10, 1}
	fmt.Println("testMergeSort success arr :", mergeSort(arr))
}

//testMerge
func testMerge() {
	arr := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	fmt.Println("testMerge success arr :", merge(arr[0:5], arr[5:]))
}

//测试时间复杂度
// 数组长度:  10000
// 运行时间: 3.5922ms
// ok
// 数组长度:  20000
// 运行时间: 6.1967ms
// ok
// 数组长度:  40000
// 运行时间: 12.8279ms
// ok
// 数组长度:  80000
// 运行时间: 25.9437ms
// ok
//可以看出数据规模变为2倍,耗时基本也是2倍 属于O(n)/O(nlogn)级别
func testTimeO() {
	//10的四次方
	testnum := 10000
	input := generateArrayRandom(testnum)
	//运行时间
	start := time.Now()
	ret := mergeSort(input)
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
	ret = mergeSort(input)
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
	ret = mergeSort(input)
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
	ret = mergeSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end = time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)
}

//测试10^7数据规模
// 数组长度:  10000000
// 运行时间: 4.7292423s
// ok
func TestTimeO2() {
	//八倍数组
	input := generateArrayRandom(1e7) //10^7
	//运行时间
	start := time.Now()
	ret := mergeSort(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end := time.Now()
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
