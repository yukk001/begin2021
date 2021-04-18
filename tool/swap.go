package tool

func Swap(i, j int, a []int) {
	a[j], a[i] = a[i], a[j]
}
