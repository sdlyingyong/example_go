斯坦福链表问题

斯坦福大学的计算机系整理了一份和链表相关的18个问题，充分编程实践这18个问题，基本就可以说把“链表”这个知识点打通关了。

这18个问题如下：

* Count 计算链表的节点个数；

* GetNth 获得链表第n个节点的值；

* CreateList 根据数组（或者标准输入）创建链表；

* DeleteList 释放一个链表的所有节点空间（C++）；

* Push 向链表头插入一个新节点；

* Pop 删除链表的第一个节点并返回；

注意：我们可以将链表看做是一个队列，此时，向链表的头或者尾插入或者删除元素，就可以衍生出两个Push实现和两个Pop实现。大家也可以显示地将这四个方法命名为addFirst, addLast, removeFirst, removeLast（参考Java中的命名方式），并据此封装底层基于链表的栈，队列，双向队列等等线性数据结构：）

* InsertNth 在链表的第n个位置插入一个新节点；

* SortedInsert 给定一个有序链表，将一个新节点插入到有序链表的正确位置；

* InsertSort 使用插入排序法为链表排序；

提示：之前实现的SortedInsert有用了：）

* Append 挂接两个链表；

* FrontBackSplit 将一个链表分割成大小相等的两个链表（对于原链表大小为奇数的情况，分割为大小只相差1的两个链表）；

* RemoveDuplicates 给定一个有序链表，其中含有重复节点，删除链表中的重复节点，使得每个不同值的节点只有一个；

* MoveNode 给定两个链表，Pop出第二个链表的元素，Push进第一个链表；

注意：这里的Pop和Push可以根据实际场景使用4中的任意一组定义

* AlternatingSplit 给定一个链表，将他分割成两个链表，其中奇数位置的节点在一个链表，偶数位置的节点在另一个链表；

提示：之前实现的MoveNode有用了：）

* ShuffleMerge 给定两个链表，将这两个链表合并成一个链表，其中一个链表的元素在奇数位，另一个链表的元素在偶数位；

提示：之前实现的MoveNode有用了：）

* SortedMerge 给定两个有序链表，将他们合并成为一个有序链表；

提示：之前实现的MoveNode有用了：）

* MergeSort 对链表进行归并排序

提示：实现了FrontBackSplit和SortedMerge，是不是觉得很简单：）当然，也可以尝试一下自底向上的归并排序（非递归的归并排序）。另外，对MergeSort的一个经典优化，是递归到达小数据量的时候，转而使用插入排序法。此时，我们自己写的InsertSort也有用了：）

* SortedIntersect 给定两个有序链表，返回一个新链表，新链表中的元素是给定两个链表的公共元素。

提示：以这个方法为基础，可以封装基于链表的集合类。大家也可以思考一下如何实现其他集合操作，如Union(合并)，Diff(差集)等等。

Reverse 反转一个链表

RecursiveReverse 使用递归的方式反转一个链表

提示：最后特意将递归方式翻转链表列出来，是因为这个问题实在是太经典了，充分体现了和链表相关算法的美丽：）



链接地址

http://cslibrary.stanford.edu/105/