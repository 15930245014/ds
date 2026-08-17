package main

import "fmt"

/**
希尔排序（Shell Sort）：插入排序的一种又称“缩小增量排序”（Diminishing Increment Sort），是直接插入排序算法的一种更高效的改进版本。希尔排序是非稳定排序算法。该方法因D.L.Shell于1959年提出而得名。
希尔排序是把记录按下标的一定增量分组，对每组使用直接插入排序算法排序；随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组，算法便终止。
 */

/**
	希尔排序具体思路：
1.设置增量，依据增量把数组分成若干组
2.每次遍历将每个组依次进行插入排序
3.改变增量大小（重复执行2，3步骤）
 */
 func shellSort( arr []int) []int {
 	//判断长度
 	if len(arr) == 0 || len(arr) == 1 {
 		//返回
 		return arr
	}

 	//获取长度
 	length := len(arr)

 	//设置增量
 	increment := length/3 + 1

 	for increment >= 1 {
 		//开始遍历
 		for i := increment; i < length ; i++ {
 			//记录值
 			key := arr[i]

 			//判断需不需要交换
 			if arr[i] < arr[i - increment] {
 				//移动位置到最前
 				j := i - increment
 				for  j >= 0 && arr[j] > key {
					arr[j + increment]  = arr[j]
 					j -= increment
				}
 				arr[j + increment] = key
			}

		}
 		//增量
		increment = increment/3
	}

 	//返回
 	return arr
 }

/**
测试希尔排序
 */
 func main() {
	 //声明数组
	 arr := []int {1,44,56,2,4,6,8,-1,-1}

	 //插入排序
	 fmt.Println(shellSort(arr))
 }
