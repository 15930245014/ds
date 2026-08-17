package main

import "fmt"

/**
	直接插入排序（InsertSort）:将一个记录插入到已经排好序的有序表中，从而得到一个新的，记录数增1的有序表
 */

 /**
 	直接插入排序
  */
 func insertSort(arr []int) []int {
 	//判断长度
 	if len(arr) == 0 || len(arr) == 1 {
 		return arr
	}

 	//开始排序--下标从1开始
 	for i := 1; i < len(arr); i ++ {
 		//记录当时的值
 		key := arr[i]

 		//比较之前
 		j := i - 1

 		for j >= 0 && arr[j] > key { //数据右移
 			arr[j + 1] = arr[j]
 			j -- //继续左移
		}

 		//判断
 		if (j + 1) != i {
 			//交换
 			arr[j + 1] = key
		}
	}

 	//返回
 	return arr

 }
/**
测试
 */
func main() {
	//声明数组
	arr := []int {1,44,56,2,4,6,8,-1}

	//插入排序
	fmt.Println(insertSort(arr))
}