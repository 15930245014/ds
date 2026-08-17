package main

import "fmt"

/**
	折半查找----前提条件---有序表
	算法描述略
 */

 /**
 	@param arr []int 查询序列
 	@param n int 查询0-n-1之间元素
 	@param key int 关键字
  */
 func binarySearch(arr []int,n int ,key int) int {
 	//定义返回值
 	rtIndex := -1

 	//判断
 	if len(arr) == 0 || n <= 0 {
 		//返回
 		return rtIndex
	}

 	//参数
 	if n > len(arr) {
 		n = len(arr)
	}

 	//定义起始
 	begin,end := 0,n-1
 	mid := -1

 	//开始查找
 	for begin <= end {
 		//折半
 		mid = begin + (end - begin) >> 1
 		//判断
 		if arr[mid] < key {//在mid-end之间
 			begin = mid + 1

		} else if arr[mid] > key {//在begin-mid之间
			end = mid - 1
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
 	arr := []int{1,3,3,5,5,7}

 	fmt.Println(binarySearch(arr,len(arr),7))
 }