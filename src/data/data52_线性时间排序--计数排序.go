package main

import "fmt"

/**
	计数排序（counting sort）：计数排序是一个非基于比较的排序算法，该算法于1954年由 Harold H. Seward 提出。它的优势在于在对一定范围内的整数排序时，它的复杂度为Ο(n+k)（其中k是整数的范围），快于任何比较排序算法。 [1-2]  当然这是一种牺牲空间换取时间的做法，而且当O(k)>O(n*log(n))的时候其效率反而不如基于比较的排序（基于比较的排序的时间复杂度在理论上的下限是O(n*log(n)), 如归并排序，堆排序）
 */
 /**
 	计数排序---适用于整数排序非比较性型排序算法
 	1、获取数值k以来申请空间---k一般为最最大值
 	2、初始化
 	3、统计出现次数
 	4、dynamic Programming 来统计前面的元素个数
 	5、构造排序的序列
  */
 func countingSort(arr []int) []int {
	//声明返回数组
	 result := make([]int,len(arr))
	 copy(result,arr)

 	//获取数组的最大值
	max := arr[0]
	for i := 1; i < len(arr); i ++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	//初始化
	var countArr []int
	for i := 0; i <= max ; i ++ {
		countArr = append(countArr,0)
	}

	//统计出现次数
	for i := 0; i < len(arr); i ++ {
		countArr[arr[i]] += 1
	}


	//dp[i] = dp[i] + dp[i - 1](i >= 1)  来统计arr[i]前面的元素个数
	for i := 1; i < len(countArr); i ++ {
		countArr[i] += countArr[i - 1]
	}

	//构造排序序列
	for i := 0; i < len(arr); i ++ {
		countArr[arr[i]] -= 1 //countArr[arr[i]] - 1 下标从0 开始
		result[countArr[arr[i]]] = arr[i]

	}

	//返回排好序的数组
	return result
 }
 /**
 	测试
  */
 func main() {
 	arr := []int{2,1,5,1,6,3,2,1,8,9}

 	fmt.Println(countingSort(arr))
 }
