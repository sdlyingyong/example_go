package main

import "fmt"

type LoopQueue struct {
	data  *MyArr
	front int
	tail  int
	size  int
}

//循环队列
//支持的方法 初始化 是否空 获取容量 获取长度
//扩容

//初始化
func (LQ *LoopQueue) Init() {
	LQ.data = NewMa()
	LQ.front = 0
	LQ.tail = 0
	LQ.size = 0
}

//是否为空
func (lq *LoopQueue) IsEmpty() bool {
	//循环头和尾在同一个位置,数据空间为0
	return lq.front == lq.tail
}

//获取长度
func (lq *LoopQueue) GetLen() int {
	return lq.data.GetLen()
}

//获取容量
func (lq *LoopQueue) GetCap() int {
	return lq.data.GetSize()
}

//入队
func (lq *LoopQueue) EnQueue(input interface{}) {
	//检查队列是否是满的
	if (lq.tail+1)%lq.data.size == lq.front {
		//如果满了,扩容为两倍容量的数组
		lq.resize(2 * lq.data.size)
	}
	//需要存入元素,并把tail++
	lq.data.Add(lq.tail, input)
	lq.tail++
}

//扩容
func (lq *LoopQueue) resize(newSize int) {
	panic("暂未实现:由于[cap]int类型数组不让创建,应该替换成slice")
	//创建新的数组
	newArr := NewMa()
	//把原来的数组的数据拷贝到新数组
	for i := 0; i < lq.data.GetLen(); i++ {
		newArr.Add(i, lq.data.Get((i+lq.front)%lq.data.GetSize()))
	}
	//把新数组赋值给原来的数组
	lq.data = newArr
	//把front和tail重新赋值
	lq.front = 0
	lq.tail = lq.data.GetLen()
}

//出队
func (lq *LoopQueue) DeQueue() interface{} {
	if lq.IsEmpty() {
		panic("队列为空")
	}
	//获取出队的元素
	ret := lq.data.Get(lq.front)
	//把front指向下一个元素
	lq.front = (lq.front + 1) % lq.data.GetSize()
	//把长度减一
	lq.size--
	if lq.size == lq.data.GetSize()/4 && lq.data.GetSize()/2 != 0 {
		//如果长度小于容量的一半,并且容量不为0,那么需要缩容
		lq.resize(lq.data.GetSize() / 2)
	}
	return ret
}

//获取最前一个元素
func (lq *LoopQueue) GetFirst() interface{} {
	if lq.IsEmpty() {
		panic("队列为空")
	}
	return lq.data.Get(lq.front)
}

//打印队列
func (lq *LoopQueue) Print() {
	ret := fmt.Sprintf("Queue len = %d cap = %d \n [ ", lq.GetLen(), lq.GetCap())
	//显示从front到tail的元素
	for i := lq.front; i != lq.tail; i = (i + 1) % lq.data.GetSize() {
		ret += fmt.Sprintf("%v ,", lq.data.Get(i))
	}
	ret += "]"
	fmt.Println(ret)
}

//测试
func testLoopQueue() {
	LoopQueue := &LoopQueue{}
	LoopQueue.Init()
	LoopQueue.EnQueue(1)
	LoopQueue.EnQueue(2)
	LoopQueue.EnQueue(3)
	LoopQueue.Print()
	LoopQueue.DeQueue()
	LoopQueue.Print()
}

func main() {
	testLoopQueue()
}

//---------------自己的动态数组----------------------
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
