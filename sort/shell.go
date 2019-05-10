package sort

import "fmt"

//希尔排序就是将排序按照步长
func ShellSort(array []int) {
	//如果长度小于2则结束
	n := len(array)
	if n < 2 {
		return
	}
	//每次排序的步长（每次为排序数组长度的一半），最终key为1，所有数据都会被排序
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				fmt.Printf("%v \n", array)
				fmt.Printf("j = 【%2d】 key = 【%2d】 j-key = 【%2d】\n", j, key, j-key)
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		fmt.Printf("\n")
		key = key / 2
	}
}