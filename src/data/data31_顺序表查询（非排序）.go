package main

import "fmt"

/**
	顺序表查询---依次遍历
 */

 //传统
 /**
 	@param arr []int 顺序表 数据未排序
 	@param n int 查询终点（索引0-n之间）
 	@param key int 查询关键字
  */
 func myfindSearch(arr []int,n int,key int) int {
 	//声明返回值
 	rtIndex := -1

 	//判断
 	if len(arr) == 0 || n < 0{
 		return rtIndex
	}

 	//验证参数
 	if n >= len(arr) {
 		n = len(arr) - 1
	}

 	//查询
 	for i := 0; i <= n; i++ {
 		//判断
 		if arr[i] == key {//由于比较次数太多---这里可以设置哨兵
 			rtIndex = i
 			break
		}
	}

 	//返回
 	return rtIndex
 }

/**
	带哨兵的 顺序表查询
	@param arr []int 顺序表 数据未排序
	@param n int 查询数组的长度
	@param key int 查询关键字
 */
func findSearch(arr []int,n int,key int) int {
	//声明返回值
	rtIndex := -1

	//判断
	if len(arr) == 0 || n <= 0{
		return rtIndex
	}

	//验证参数
	if n > len(arr) {
		n = len(arr)
	}

	//比较
	if arr[n - 1] == key {
		rtIndex = n - 1
		return rtIndex
	} else {//n位置设置哨兵
		arr[n - 1] = key
	}

	//查询
	i := 0
	for  arr[i] != key {
		i++
	}

	//判断
	if i != n-1 {//i == 哨兵位置表示没有
		rtIndex = i
	}

	//返回
	return rtIndex
}


/**
	test
 */
func  main ()  {
	arr := []int{1,2,5,3,9,0}
	//不带哨兵
	fmt.Println(myfindSearch(arr,5,3))

	//带哨兵
	fmt.Println(findSearch(arr,5,3))
	fmt.Println(findSearch(arr,5,-1))
}

