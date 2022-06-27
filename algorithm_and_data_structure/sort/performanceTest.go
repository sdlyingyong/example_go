package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	testLen int = 1e5 //测试数量级
)

func main() {

	fmt.Println("quickSortBase")
	TestTimeVs(testLen, quickSortBase)

	// fmt.Println("quickSortMe")
	// TestTimeVs(testLen, quickSortMe)

	fmt.Println("mergeSort")
	TestTimeVs(testLen, mergeSort) //多两个数量级,都依旧小于section的耗时

	fmt.Println("insertionSort")
	TestTimeVs(testLen, insertionSort)

	fmt.Println("sectionSort")
	TestTimeVs(testLen, sectionSort)
}

//测试同样数据,归并排序和插入排序,选择排序的耗时
func TestTimeVs(count int, funcName func([]int) []int) {
	//归并排序 merge sort
	input := generateArrayRandom(count)
	//运行时间
	start := time.Now()
	ret := funcName(input)
	//数组长度
	fmt.Println("数组长度: ", len(input))
	//结束时间
	end := time.Now()
	fmt.Println(" 运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)
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

//快速排序 标准版
func quickSortBase(arr []int) []int {
	return quickSortBaseHelp(arr, 0, len(arr)-1)
}

//快速排序 标准版
//辅助函数
func quickSortBaseHelp(arr []int, low, high int) []int {
	if low >= high {
		return arr
	}
	arr, p := partition(arr, low, high)
	arr = quickSortBaseHelp(arr, low, p-1)
	arr = quickSortBaseHelp(arr, p+1, high)
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

//------------辅助函数-----------------

//生成数组长度n,乱序不重复的数组
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
