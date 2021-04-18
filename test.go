package main

import (
	"fmt"
	"math/rand"
	"mytest/mysort"
	"time"
)

const RANGE_AREA = 20

func testQuickSort(){
	initArr:=RandArray(rand.Intn(RANGE_AREA))
	result := mysort.QuickSort(initArr)
	fmt.Println(result)
}
func testOneByOne(){
	initArr:=RandArray(RANGE_AREA)
	result := mysort.OneByOne(initArr)
	fmt.Println(result)
}

func main(){
	//快排
	//testQuickSort()
	//冒泡
	testOneByOne()
}

func RandArray(num int) []int {
	arr := make([]int, num)
	//以当前时间为随机数种子
	rand.Seed(time.Now().Unix())

	for i := 0; i < num; i++ {
		arr[i] = rand.Intn(100)

	}
	return arr

}
