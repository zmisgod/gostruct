package sort

// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，
// 然后放到已排序序列的末尾。
// 以此类推，直到所有元素均排序完毕。
func SelectSort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	sorted := make([]int, 0)
	min := arr[0]
	for len(arr) > 0 {
		min, arr = minF(arr)
		sorted = append(sorted, min)
	}
	return sorted
}

func minF(arr []int) (int, []int) {
	res := make([]int, 0)
	index := 0
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
			index = i
		}
	}
	for i:=0;i<len(arr) ;i++  {
		if i != index {
			res = append(res, arr[i])
		}
	}
	return min, res
}
