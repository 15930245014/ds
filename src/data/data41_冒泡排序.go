package main

import "fmt"

/**
	冒泡排序（Bubble Sort）：是一种交换排序它的基本思想是 -> 两两比较记录的值如果反序则交换直到没有反序为止
	实现方式3种：
	1、之前写的
	2、正宗的冒泡排序
	3、优化后的冒泡排序（重点掌握）
 */

/**
 	之前的
 */
func bubbleSort1(arr []int) []int {
	//判断arr
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	//获取长度
	length := len(arr)

	//开始排序
	for i := 0; i < length; i ++ {
		for j := i + 1; j < length; j ++ {
			//判断
			if arr[j] < arr[i] { //交换记录
				arr[j], arr[i] = arr[i], arr[j]
			}

		}
	}

	//返回
	return arr

}

/**
	经典冒泡排序
 */
func bubbleSort2(arr []int) []int {
  	//判断长度
  	if len(arr) == 0 || len(arr) == 1 {
  		return arr
	}

  	//获取长度
  	length := len(arr)

  	//开始排序
  	for i := 0; i < length; i++ {
  		for j := length - 1 ; j > i ; j -- {//从后向前比较
  			if arr[j] < arr[j - 1] {//交换
  				arr[j],arr[j - 1] = arr[j - 1],arr[j]
			}

		}

	}
	//返回
	return arr
}
/**
	根据经典来优化
	待排序为{2,1,3,4,5,6,7,8} 除了第一和第二关键字需要交换外别的都是正常的顺序，此时序列已经有序所以不需要在往后比较了
 */
func bubbleSort3(arr []int)  []int {
	//判断长度
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	//增加一个flag
	fg := true

	//开始排序
	for i := 0; i < len(arr) && fg; i ++ { //只有交换过才继续
		fg = false //表示没有交换过
		for j := len(arr) - 1; j > i ; j -- {
			//比较并交换
			if arr[j] < arr[j - 1] {
				arr[j],arr[j - 1] = arr[j - 1],arr[j]
				fg = true
			}

		}
	}

	//返回
	return arr

}

/**
测试
 */
func main() {
	//声明arr
	arr := []int{2,1,3,4,5}

	//冒泡排序1
	//fmt.Println(bubbleSort1(arr))

	//冒泡排序2
	//fmt.Println(bubbleSort2(arr))

	//冒泡排序3
	fmt.Println(bubbleSort3(arr))
}
