package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{3, 5, 1, 54, 8, 42, 4, 6, 89, 3}
	res := InsertSort(arr)
	fmt.Println(res)
}
