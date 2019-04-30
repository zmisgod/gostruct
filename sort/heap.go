package sort

//HeapSort S
func HeapSort(array []int) {
	m := len(array)
	//找出所有叶子节点
	s := m / 2
	//先将数组组成二叉树（符合堆的特性：父节点比子节点大，左子树比右子树大）
	//所以此处仅比较叶子节点即可（也就是为什么m/2）
	for i := s; i > -1; i-- {
		minHeap(array, i, m-1)
	}
	//然后在进行排序
	for i := m - 1; i > 0; i-- {
		//每一次将i与第一个值进行互换，然后此时i位之后的数都是排序好的数，前面的数会经过heap函数处理，每次都会将最大的值放在第一个，
		//然后重复将i为与第一个值位置互换
		array[i], array[0] = array[0], array[i]
		minHeap(array, 0, i-1)
	}
}

//maxHeap 最大堆
func maxHeap(array []int, father, end int) {
	//左子节点索引
	l := 2*father + 1
	//如果左子节点索引大于堆的索引长度（父节点没有子节点了），终止
	if l > end {
		return
	}
	//n为最大值的索引
	largest := l
	//右子节点索引
	r := 2*father + 2
	//如果右子节点比左子节点大，那么largest为右子节点索引
	if r <= end && array[r] > array[l] {
		largest = r
	}
	//如果父节点值比largest索引的值大，退出
	if array[father] > array[largest] {
		return
	}
	//否则largest与father的位置调换
	array[largest], array[father] = array[father], array[largest]
	//递归对比largest后续的值是否也是下面子节点最大的
	maxHeap(array, largest, end)
}

//minHeap 最小堆
func minHeap(array []int, father, end int) {
	//左子节点索引
	l := 2*father + 1
	//如果左子节点索引大于堆的索引长度（父节点没有子节点了），终止
	if l > end {
		return
	}
	//n为最小值的索引
	min := l
	//右子节点索引
	r := 2*father + 2
	//如果右子节点比左子节点大，那么largest为右子节点索引
	if r <= end && array[r] < array[l] {
		min = r
	}
	//如果父节点值比largest索引的值小，退出
	if array[father] < array[min] {
		return
	}
	//否则i与n的位置调换
	array[min], array[father] = array[father], array[min]
	//递归对比min后续的值是否也是下面子节点最大的
	minHeap(array, min, end)
}
