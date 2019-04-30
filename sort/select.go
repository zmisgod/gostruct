package sort

// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，
// 然后放到已排序序列的末尾。
// 以此类推，直到所有元素均排序完毕。
func SelectSort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
