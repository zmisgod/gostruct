package sort

// 对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
func InsertSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	for i := 0; i < len(data); i++ {
		for j := i - 1; j >= 0; j-- {
			if data[j+1] < data[j] {
				data[j+1], data[j] = data[j], data[j+1]
			} else {
				break
			}
		}
	}
	return data
}
