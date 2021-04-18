package mysort

import "mytest/tool"

func OneByOne(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tool.Swap(j,j+1,arr)
			}
		}
	}

	return arr
}
