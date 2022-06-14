package main

import (
	"fmt"
)

type MyArr struct {
	data [10]interface{} //存储数组元素,外界不能直接访问
	size int             //数组固定容量
}

//数组操作 插入指定位置 删除 访问元素 检查数组是否为空 获取数组容量

//初始化数组 暂无

//无参数初始化
func NewMa() *MyArr {
	newArr := &MyArr{
		data: [10]interface{}{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		size: 10,
	}
	return newArr
}

//获取数组当前的容量
func (ma *MyArr) GetSize() int {
	return ma.size
}

//获取数组当前元素个数
func (ma *MyArr) GetLen() int {
	len := 0
	for i := range ma.data {
		if ma.data[i] != nil {
			len++
		}
	}
	return len
}

//是否为空
func (ma *MyArr) IsEmpty() bool {
	return ma.size == 0
}

//向头部添加元素
//往头部添加,需要把数组每个元素往右搬一位,所以 时间复杂度 O(n)
func (ma *MyArr) AddFirst(input interface{}) {
	ma.Add(0, input)
	return
}

//测试头部添加元素
func testAddFirst() {
	ma := NewMa()
	ma.AddFirst(10)
	ma.Print()
	ma.AddFirst(9)
	ma.Print()
}

//向尾部添加元素
//末尾添加不需要搬动任何元素,所以 时间复杂度 O(1)
func (ma *MyArr) AddLast(input int) (err error) {
	ma.Add(ma.GetLen(), input)
	return
}

func testAddLast() {
	nArr := NewMa()
	nArr.Print()
	nArr.AddLast(1)
	nArr.AddLast(2)
	nArr.Print()
}

//向index位置插入元素
//由于index=0时,O(n);index=size时,O(1)
//从概率论角度 给出意见是O(2/n)
//从代码实践忽略常数,所以时间复杂度 O(n)
func (ma *MyArr) Add(index int, input interface{}) {
	//检查是否越界
	if index < 0 || index > ma.size {
		fmt.Println("插入位置越界")
		return
	}
	//检查是否用完容量
	if ma.GetSize() == ma.GetLen() {
		//如果满了,扩容为两倍容量的数组
		ma.resize(2 * ma.GetLen())
		return
		// fmt.Println("数组已经满了")
		// return
	}
	//范围 (x...end)]
	//把插入位置后面的数组往后挪动一位,腾出空间
	//从后往前挪,否则会覆盖掉原来的数据
	for i := ma.GetLen() - 1; i >= index; i-- {
		ma.data[i+1] = ma.data[i]
	}
	//插入指定位置
	ma.data[index] = input
	return
}

func testAdd() {
	nArr := NewMa()
	fmt.Println(nArr.data)
	nArr.Add(0, 1)
	fmt.Println(nArr.data)
}

func (ma *MyArr) Print() {
	ret := fmt.Sprintf("Array len = %d cap = %d \n [ ", ma.GetLen(), ma.GetSize())
	for i := 0; i < len(ma.data); i++ {
		if ma.data[i] != nil {
			ret += fmt.Sprintf("%d, ", ma.data[i])
		}
	}
	ret += "]"
	fmt.Println(ret)
}

func testPrint() {
	nArr := NewMa()
	fmt.Println(nArr.data)
	nArr.AddLast(1)
	nArr.AddLast(2)
	nArr.Print()
}

//取出元素
func (ma *MyArr) Get(index int) (value interface{}) {
	//只允许访问有效元素
	if index < 0 || index >= ma.GetLen() {
		fmt.Println("out of index")
		return
	}
	value = ma.data[index]
	return
}

//测试取出元素
func testGet() {
	nArr := NewMa()
	nArr.AddLast(1)
	nArr.AddLast(2)
	fmt.Println("nArr.Get(0) :", nArr.Get(0))
	fmt.Println("nArr.Get(1) :", nArr.Get(1))
	//异常情况
	fmt.Println("nArr.Get(2) :", nArr.Get(2))
	fmt.Println("nArr.Get(2) :", nArr.Get(-1))
}

//修改元素
//支持随机访问
//直接修改元素,所以 时间复杂度是 O(1)
func (ma *MyArr) Set(index int, value interface{}) {
	//检查越界情况
	if index < 0 || index >= ma.GetLen() {
		fmt.Println("out of index")
		return
	}
	ma.data[index] = value
	return
}

func testSet() {
	nArr := NewMa()
	nArr.AddLast(1)
	nArr.AddLast(2)
	nArr.Set(0, 3)
	nArr.Print()
	//异常情况
	nArr.Set(2, 3)
	nArr.Set(-1, 3)
}

//数组中是否有某个元素
//需要遍历整个数组,来判断是否有这个元素,所以 时间复杂度是 O(n)
func (ma *MyArr) contains(value interface{}) (has bool) {
	for i := 0; i < ma.GetLen(); i++ {
		if ma.data[i] == value {
			return true
		}
	}
	return false
}

func testContains() {
	nArr := NewMa()
	nArr.AddLast(1)
	nArr.AddLast(2)
	fmt.Println("nArr.contains(1) :", nArr.contains(1))
	fmt.Println("nArr.contains(3) :", nArr.contains(3))
}

//数组中搜索元素
//时间复杂度 O(n)
func (ma *MyArr) Search(value interface{}) (index int) {
	for i := 0; i < ma.GetLen(); i++ {
		if ma.data[i] == value {
			return i
		}
	}
	return -1
}

func testSearch() {
	nArr := NewMa()
	nArr.AddLast(1)
	nArr.AddLast(2)
	fmt.Println("nArr.Search(1) :", nArr.Search(1))
	fmt.Println("nArr.Search(3) :", nArr.Search(3))
}

//删除元素
//从删除index后所有的元素都需要往左移动
//[index...end]都需要设置为i+1
//size需要更新--
//复杂度范围从O(1)到O(n)都有可能 概率上认为是O(2/N)
//代码中认为是O(n)
func (ma *MyArr) Remove(index int) (value interface{}) {
	//检查删除元素索引合规
	if index < 0 || index >= ma.GetLen() {
		fmt.Println("out of index")
		return
	}
	//存储删除的元素
	value = ma.data[index]
	//[index...end)需要设置为之后的元素
	for i := index; i < ma.GetLen()-1; i++ {
		ma.data[i] = ma.data[i+1]
	}
	//最后一个元素设置为nil
	ma.data[ma.GetLen()-1] = nil
	//缩容操作,如果当前元素长度为容量的1/4,则缩容
	if ma.GetLen() <= ma.GetSize()/4 && ma.GetLen()/2 != 0 {
		//制造一个1/4长度的数组,并遍历拷贝当前所有元素
		//code...
	}
	return
}

//删除最后一个元素
//不需要搬动任何一个元素,所以时间复杂度是 O(1)
func (ma *MyArr) RemoveLast() (value interface{}) {
	return ma.Remove(ma.GetLen() - 1)
}

//删除第一个元素
//需要搬动数组中所有元素,所以时间复杂度是 O(n)
func (ma *MyArr) RemoveFirst() (value interface{}) {
	return ma.Remove(0)
}

//从数组中删除某个元素
//需要遍历搜索数组再删除,所以时间复杂度是 O(n)
func (ma *MyArr) RemoveElement(value interface{}) bool {
	index := ma.Search(value)
	if index != -1 {
		val := ma.Remove(index)
		if val != nil {
			return true
		}
	}
	return false
}

func testRemove() {
	nArr := NewMa()
	nArr.AddLast(1)
	nArr.AddLast(2)
	nArr.AddLast(3)
	nArr.AddLast(4)
	nArr.AddLast(5)
	nArr.Print()
	nArr.Remove(0)
	nArr.Print()
	nArr.Remove(2)
	nArr.Print()
	nArr.Remove(4)
	nArr.Print()
	nArr.Remove(4)
	nArr.Print()
	nArr.AddLast(1)
	nArr.Print()
	nArr.RemoveFirst()
	nArr.Print()
	nArr.RemoveLast()
	nArr.Print()
	fmt.Println("nArr.RemoveElement(3) :", nArr.RemoveElement(3))
	nArr.Print()
	fmt.Println("nArr.RemoveElement(3) :", nArr.RemoveElement(3))
}

//resize扩容 不允许用户访问
//go中1024之前是两倍扩容,之后是每次+1/4
func (ma *MyArr) resize(capacity int) (newArr *MyArr) {
	fmt.Println("resize is not allowed")
	return
	// //新建一个数组
	// newArr = NewMa()
	// //把原来的数组元素拷贝到新数组
	// for i := 0; i < capacity; i++ {
	// 	newArr.data[i] = ma.data[i]
	// }
	// //把新数组的容量设置为原来的容量
	// newArr.size = capacity
	// //返回数组
	// return
}

func testResize() {
	nArr := NewMa()
	nArr.Add(0, 1)
	nArr.Add(1, 2)
	nArr.Add(2, 3)
	nArr.Add(3, 4)
	nArr.Add(4, 5)
	nArr.Add(5, 6)
	nArr.Add(6, 7)
	nArr.Add(7, 8)
	nArr.Add(8, 9)
	nArr.Add(9, 10)
	nArr.Add(10, 11)
	nArr.Add(11, 12)
	nArr.Add(12, 13)
	nArr.Print()
}

func main() {
	//设计自己的数组类型
	// testAddLast()
	// testAddFirst()
	// testInsert()
	// testPrint()
	// testGet()
	// testSet()
	// testContains()
	// testSearch()
	testRemove()
	// testResize()
}
