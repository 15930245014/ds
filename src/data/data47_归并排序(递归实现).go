package main

import (
	"fmt"
	"sync"
)

/**
	归并排序(merging sort)：假设初始序列还有n个记录，则可以可成是n个有序序列的子序列，每个子序列的长度的长度为1，然后两两合并，得到[n/2]长的为2或1的有序子序列；再两两合并，......，如此重复，直到得到一个长度为n的有序序列为止，这种排序方法成为2路归并排序
 */
/**
	归并排序 ----许校
 */
 func mergingSort(arr []int) []int {
	var result []int

 	MSort(arr,0,len(arr) - 1,&result)

	return result
 }
 /**
 	排序递归方法
  */
 func MSort(arr []int,s,l int,result *[]int){

 	//判断
 	if s == l {
 		//返回
 		*result = append(*result,arr[s])


	} else {
		//协程组
		wg := sync.WaitGroup{}
		wg.Add(2)
		//分割
		m := (s + l) / 2

		var left,right []int


		//分治
		go func(arr []int,s,l int,r *[]int) {
			MSort(arr,s,l,r)
			wg.Done()
		}(arr,s,m,&left)
		go func(arr []int,s,l int,r *[]int) {
			MSort(arr,s,l,r)
			wg.Done()
		}(arr,m + 1,l,&right)
		wg.Wait()


		//表示都执行完
		*result = merging(left,right)
	}


 }



/**
	归并排序合并算法----和合并2个还好序的数组是一个意思
	arr [s-m] 和 [m+1,l] 下标有序
 */
func merging(tmp1,tmp2 []int) []int {
	//声明合并后的
	var result []int

	//开始合并
	i, j := 0,0

	for i < len(tmp1) && j < len(tmp2) {
		//判断
		if tmp1[i] < tmp2[j] {
			result = append(result,tmp1[i])
			i++
		} else {
			result = append(result,tmp2[j])
			j++
		}
	}

	//判断结束
	for i < len(tmp1) {
		result = append(result,tmp1[i])
		i++
	}
	for j < len(tmp2) {
		result = append(result,tmp2[j])
		j++
	}

	//返回
	return result
}

/**
  测试子方法
 */
 func main() {
	 arr1 := []int{1, 2, 3, -1, 0, 6}
	 //	arr2 := []int{-1,0,6}

 	fmt.Println(mergingSort(arr1))



 }
