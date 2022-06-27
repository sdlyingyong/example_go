package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// testMerge()
	// testMergeSort()
	// testTimeO()
	// TestTimeO2()
	// testMergeSortDep()
	// testMergeBetter()
	TestMergeSortUpward()
}

// 	数组长度:  4000000
// 运行时间: 1.847265s
// ok
// 数组长度:  4000000
// 运行时间: 1.4152578s
// ok
func testMergeBetter() {

	//测试归并排序
	var testnum int = 1e6 * 4
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

	//测试归并排序优化版 len<15使用归并
	input = generateArrayRandom(testnum)
	//运行时间
	start = time.Now()
	ret = mergeSortBetterTime(input)
	//数组长度
	fmt.Println("数组长度: ", len(ret))
	//结束时间
	end = time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)

}

//打印归并排序的过程
func testMergeSortDep() {
	arr := generateArrayRandom(10)
	arr = mergeSortDep(arr, 0)
	fmt.Println("testMergeSortDep ret :", arr)
	checkSort(arr)
}

func printLine(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("--")
	}
	fmt.Println()
}

//合并排序带深度
func mergeSortDep(arr []int, dep int) (ret []int) {
	printLine(dep)
	len := len(arr)
	if len < 2 {
		ret = arr
		fmt.Println("返回单个元素 :", ret)
		return
	}
	mid := len / 2
	left := arr[0:mid] //不包含mid
	right := arr[mid:] //从mid到最后
	fmt.Println("拆分为两个数组 left :", left, "right :", right)
	l := mergeSortDep(left, dep+1)
	r := mergeSortDep(right, dep+1)
	fmt.Printf("开始合并 left: %d, right: %d : \n", l, r)
	ret = merge(l, r)
	fmt.Println("合并两个数组为一个 ret: ", ret)
	return
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
	if len(arr) == 0 {
		fmt.Println("error: empty")
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("error: un sort")
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

//合并排序优化
//使用插入排序处理小于15的数组,减少
func mergeSortBetterTime(arr []int) (ret []int) {
	len := len(arr)
	//如果数量少的时候,使用插入排序法
	if len <= 15 {
		return insertionSort(arr)
	}
	if len < 2 {
		ret = arr
		return
	}
	mid := len / 2
	left := arr[0:mid] //不包含mid
	right := arr[mid:] //从mid到最后
	return merge(mergeSortBetterTime(left), mergeSortBetterTime(right))
}

//自底向上的归并排序
func mergeSortUpward(arr []int) (ret []int) {
	//1.最小情况的解决方案
	if len(arr) == 2 {
		if arr[0] < arr[1] {
			ret = []int{arr[0], arr[1]}
		} else {
			ret = []int{arr[1], arr[0]}
		}
	}
	//2.递归,当做调用子函数
	ret = merge([]int{1, 4}, []int{2, 3})
	return
}

// func mergeSortButton2Top(arr []int) {
// 	var lenth int = len(arr)
// 	for size := 1; size <= lenth; size += size {
// 		for i := 0; i+size < lenth; i += 2 * size { //对[i,i+size-1]和[i+size,i+2*size-1]进行归并
// 			merge2(arr, i, i+size-1, int(math.Min(float64(i+2*size-1), float64(lenth-1)))) // arr left mid right  如果i+2*size>n了，越界了，就取n-1
// 		}
// 	}
// }

// func merge2(arr []int, left, mid, right int) {
// 	// 将要合并的部分做个拷贝
// 	var tmp []int = make([]int, right-left+1)
// 	for i, j := left, 0; i <= right; i++ {
// 		tmp[j] = arr[i]
// 		j++
// 	}
// 	// i做为左半部分的指针   j作为右半部分的指针
// 	var i, j int = left, mid + 1
// 	for k := left; k <= right; k++ {
// 		if i > mid { // 左半部分 已经合入完了，将右半部分剩下的 全部合入
// 			arr[k] = tmp[j-left]
// 			j++
// 		} else if j > right { // 右半部分 已经合入完了，将左半部分剩下的 全部合入
// 			arr[k] = tmp[i-left]
// 			i++
// 		} else if tmp[i-left] > tmp[j-left] {
// 			arr[k] = tmp[j-left]
// 			j++
// 		} else {
// 			arr[k] = tmp[i-left]
// 			i++
// 		}
// 	}
// }

//自底向上的归并排序
//辅助函数
func mergeSortUpwardHelper(arr, left, right []int) (ret []int) {
	// //所有元素两两组成有序数组
	// for i := 0; i < len(arr); i += 2 {
	// 	//两个元素排序
	// 	if i+1 > len(arr) {
	// 		return []int{arr[i+1]}
	// 	}
	// 	if arr[i] < arr[i+1] {
	// 		ret = []int{arr[i], arr[i+1]}
	// 	} else {
	// 		ret = []int{arr[i+1], arr[i]}
	// 	}
	// 	return
	// }
	//1.最小情况
	if len(left) == 2 {
		if arr[0] < arr[1] {
			ret = []int{arr[0], arr[1]}
		} else {
			ret = []int{arr[1], arr[0]}
		}
	}
	//2.当做调用别的函数
	return
}

func TestMergeSortUpward() {
	//测试归并排序
	var testnum int = 1e6 * 4
	input := generateArrayRandom(testnum)
	input = []int{1, 4, 2, 3}
	//运行时间
	start := time.Now()
	ret := mergeSortUpward(input)
	fmt.Println("ret :", ret)
	//数组长度
	fmt.Println("数组长度: ", len(ret))
	//结束时间
	end := time.Now()
	fmt.Println("运行时间:", end.Sub(start))
	//检查ret是否排序
	checkSort(ret)
}
