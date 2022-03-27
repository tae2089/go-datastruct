package main

import (
	"DataStruct/sort"
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	arr2 := sort.MergeSort(arr)
	arr3 := sort.QuickSort(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
