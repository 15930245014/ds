    package main

import "fmt"

/**
	快排序优化：
	1、优化杻轴（三数取中）即选取三个关键字先进性排序，将这三个数的中间数作为杻轴，一般是取左端、右端和中间三个数
	2、不要的交换
	3、优化小数组时的排序方案 ---如果数组中只有几个记录时 则采取插入排序效率更高---高炮打蚊子
	4、优化递归操作
 */

const   MIN_ARR_LENGTH  = 3 //定义数组如果小于20采用直接插入排序
/**
	直接插入排序---优化小数组
 */
func InsertSort(arr *[]int,low,high int){
	//判断长度
	if (high - low) == 0 {
		return
	}

	for i := low + 1 ; i <= high ; i ++ {
		//记录key
		key := (*arr)[i]

		//判断
		j := i - 1

		for j >= low && (*arr)[j] > key {//右移
			(*arr)[j + 1] = (*arr)[j]
			j --
		}
		(*arr)[j + 1] = key
	}

}

func QuickSort(arr []int) []int {
	//快排序
	qSort(&arr,0,len(arr) - 1)
	return arr
}

func qSort(arr *[]int,low,high int) {
	//判断
	if (high - low) >= MIN_ARR_LENGTH  {//优化小数组
		//优化递归
		for low < high {
			//杻轴
			pivot := Pivot(arr,low,high)

			//分割
			qSort(arr,low,pivot)

			//优化递归
			low = pivot + 1

		}
	} else {
		InsertSort(arr,low,high)
	}

}
/**
	杻轴排列函数
 */
func Pivot(arr *[]int,low,high int) int {
	//三数取中
	mid := low + (high - low) / 2

	//三数取中
	if (*arr)[low] > (*arr)[high] { //和最右端比较保证左边较小
		(*arr)[low] , (*arr)[high] = (*arr)[high] ,(*arr)[low]
	}
	if  (*arr)[mid] > (*arr)[high] {//和中间比较保证最右边中间较小
		(*arr)[high] , (*arr)[mid] = (*arr)[mid] ,(*arr)[high]
	}

	//中间和最左端比较保证mid较小
	if (*arr)[mid] > (*arr)[low] {
		(*arr)[low] , (*arr)[mid] = (*arr)[mid] ,(*arr)[low]
	}

	key := (*arr)[low]

	//开始排列
	for low < high {
		for low < high && (*arr)[high] >= key {
			high --
		}

		//替换---去掉不必要的交换
		(*arr)[low] = (*arr)[high]

		for low < high && (*arr)[low] <= key {
			low ++
		}

		//同上
		(*arr)[high] = (*arr)[low]

	}

	//写入杻轴
	(*arr)[low] = key

	return low

}
/**
	测试
 */
func main() {
	//声明数组
	arr := []int {1,44,56,2,4,6,8,-1,0,100}

	//插入排序
	fmt.Println(QuickSort(arr))
}