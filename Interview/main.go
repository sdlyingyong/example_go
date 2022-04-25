package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//配套问题 极客时间go面试题2021
	//11、Go 语言中 cap 函数可以作用于那些内容？
	showCap()

	//15、PRINTF(),SPRINTF(),FPRINTF() 都是格式化输，有什么不同？
	showSprintf()

	//19、GO 语言是如何实现切片扩容的？
	showSliceAppend()

	//23、扩容前后的 Slice 是否相同？
	showSliceDiff()

	//v2ex 循环中操作数组元素
	showForRange()

	//v2ex,读取文件
	showReadFile()

}

func showReadFile() {
	file, err := os.Open("")
	bufio.NewScanner(file)
}

func showForRange() {
	items := [...]int{10, 20, 30, 40, 50}
	fmt.Println("items :", items)
	for k, item := range items {
		fmt.Println("k: ", k, "&item: ", &item, "&items[k] :", &items[k])
	}
	fmt.Println("items :", items)
}

//11、Go 语言中 cap 函数可以作用于那些内容？
func showCap() {
	slice1 := make([]string, 0, 10)
	fmt.Println("cap(slice1) : ", cap(slice1))

	arr1 := [...]int{1, 2, 3}
	fmt.Println("cap(arr1) : ", cap(arr1))

	chan1 := make(chan int)
	fmt.Println("cap(chan1) : ", cap(chan1))

	chan2 := make(chan int, 10)
	fmt.Println("cap(chan2) : ", cap(chan2))
}

//15、PRINTF(),SPRINTF(),FPRINTF() 都是格式化输，有什么不同？
func showSprintf() {
	fmt.Printf("hello %s \n", "world")
	fmt.Fprintf(os.Stdout, "hello %s \n", "world")
	str := fmt.Sprintf("hello %s \n", "world")
	fmt.Println(str)
}

//19、GO 语言是如何实现切片扩容的？
func showSliceAppend() {
	sl1 := make([]int, 0, 0)
	for i := 0; i < 3000; i++ {
		sl1 = append(sl1, i)
		fmt.Printf("slice cap: %d len: %d \n", cap(sl1), len(sl1))
	}
}

//23、扩容前后的 Slice 是否相同？
func showSliceDiff() {
	//拷贝的切片底层是同一个引用数组,修改会互相影响
	var slice3 = []int{1, 2, 3}
	slice4 := slice3
	slice4[0] = 10
	fmt.Printf("slice3 = %v slice4 = %v \n", slice3, slice4)
}
