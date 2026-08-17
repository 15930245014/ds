package main

import "fmt"

/**
	简单选择排序(Simple Selection Sort)：通过n-i次关键字间的比较，从 n - i + 1个记录中选出关键字的最小的记录，并和第i个记录交换
 */

 func selectSort(arr []int)[]int {
 	//判断
 	if len(arr) == 0 || len(arr) == 1 {
 		return  arr
	}

 	//获取长度
 	length := len(arr)

 	//选择排序
 	for i := 0 ; i < length; i ++ {
 		min := i //记录下标
 		for j := i + 1; j < length; j ++ {
 			//比较
 			if arr[j] < arr[min] {
 				min = j
			}

		}

 		//交换
 		if i != min {//判断一下
 			arr[i],arr[min] = arr[min],arr[i]
		}

	}

 	//返回
 	return arr
 }
/**
测试
 */
func main (){
	arr := []int{22,3,4,67,89,4,-1}

	fmt.Println(selectSort(arr))
}