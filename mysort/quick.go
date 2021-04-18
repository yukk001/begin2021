package mysort

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 1{
		return arr
	}
	beginNode := arr[rand.Intn(len(arr))]

	var leftArr []int
	var midArr = []int{}
	var rightArr = []int{}
	for _,v :=range arr{
		if v < beginNode {
			leftArr = append(leftArr, v)
		}
		if v == beginNode {
			midArr = append(midArr, v)
		}
		if v > beginNode {
			rightArr = append(rightArr, v)
		}
	}

	rightArr = QuickSort(rightArr)
	leftArr = QuickSort(leftArr)

	result := append(leftArr, midArr...)
	result = append(result, rightArr...)

	return result
}
