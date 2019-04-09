package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{3,5,1,54,8,42,4,6,89,3}
	res := SelectSort(arr)
	fmt.Println(res)
}