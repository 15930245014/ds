package main

import "fmt"

/**
	插值查找---排序表
	基本思想：基于二分查找，将查找点的选择改进为自适应选择，可以提高查找效率。 将二分查找的插值计算公式改为
	mid = low +  (key - arr[low])/(arr[high] - arr[low]) * (high - low)

	###对于表长较大，而关键字分布又比较均匀的查找表来说，插值查找的平均性能比折半查找要好。
 */

 func insertValueSearch(arr []int,n int,key int) int {
 	//声明返回
 	rtIndex := -1

 	//判断
 	if len(arr) == 0 || n <= 0{
 		return rtIndex
	}

 	//判断参数
 	if n > len(arr) {
 		n = len(arr)
	}

 	//声明起止
 	low,high := 0,n - 1

 	//查找
 	for low <= high {
 		//插值公式---与二分查找不同
 		mid := low + (key - arr[low])/(arr[high] - arr[low])*(high - low)

 		//判断
 		if arr[mid] > key {
 			high = mid - 1
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			rtIndex = mid
			break
		}

	}

 	//返回
 	return rtIndex

 }
/**
	test
 */
 func main() {
 	arr := []int{1,2,3,4,5,6,7,8}

 	fmt.Println(insertValueSearch(arr,2,4))
 }