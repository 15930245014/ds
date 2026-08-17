package main

import "fmt"

/**
	思路 ：利用折半查找来寻找记录插入位置
 */
 func binaryInsertSort(arr []int)[]int {
 	//判断条件
 	if len(arr) == 0 || len(arr) == 1 {
 		//返回
 		return arr
	}

 	//获取长度
 	length := len(arr)

 	//开始排序
 	for i := 1; i < length; i ++ {
 		//记录key
 		key := arr[i]

 		//查询位置
 		low,high := 0,i - 1

 		//二分查找---完事之后low即为要插入的位置
 		for low <= high {
 			mid := (low + high) / 2
 			if arr[mid] <= key {
 				low = mid + 1
			} else {
				high = mid - 1
			}

		}

 		//此时mid即为要插入的位置---右移元素 mid往后
 		for j := i; j > low; j -- {
 			arr[j] = arr[j - 1]
		}
 		arr[low] = key
	}

 	//返回
 	return arr

 }

 /**
 测试折半插入排序
  */
  func main() {
	  //声明数组
	  arr := []int {0 ,1 ,7 ,7, 18, 7, 56 ,15, 106}

	  //插入排序
	  fmt.Println(binaryInsertSort(arr))
  }
