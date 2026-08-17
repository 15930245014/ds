package main

import "fmt"

/**
	快排序（quick sort）：通过一趟排序将待排序记录分割成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，已达到整个序列有序的目的,这个关键字叫做杻轴（pivot）
 */
func  quickSort(arr []int)[]int {
	//调用子函数
	 QSort(&arr,0,len(arr) - 1)

	 //返回
	 return arr
}
/**
	递归排序
 */
func QSort(arr *[]int,s,l int) {
	//递归
	if s < l {
		//杻轴
		pivot := pivot(arr,s,l)

		//以杻轴分割
		QSort(arr,s,pivot - 1)
		QSort(arr,pivot + 1,l)

	}

}
/**
	以杻轴分割两部分
 */
func pivot(arr *[]int,low,high int) int {
	//记录杻轴
	key := (*arr)[low]

	//开始排列
	for low < high {
		//从后面找一个比key小的
		for low < high && (*arr)[high] >= key {
			high --
		}
		//交换
		if low < high {
			(*arr)[low],(*arr)[high] = (*arr)[high],(*arr)[low]
		}
		//从前面找一个比key大的
		for low < high && (*arr)[low] <= key {
			low ++
		}

		//交换
		if low < high {
			(*arr)[low],(*arr)[high] = (*arr)[high],(*arr)[low]
		}

	}
	//返回
	return low
}

/**
	测试
 */
 func main() {
	 //声明数组
	 arr := []int {1,44,56,2,4,6,8,-1,0,100}

	 //插入排序
	 fmt.Println(quickSort(arr))
 }