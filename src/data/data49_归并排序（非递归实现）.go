package main

import "fmt"

/**
	归并排序非递归实现 采用自底向上的形式--22合并直到结束
 */


/**
   归并排序合并算法----和合并2个还好序的数组是一个意思
   arr [s-m] 和 [m+1,l] 下标有序
*/
func mSort(tmp1,tmp2 []int) []int {
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
	归并排序--入口
 */
func MergingSort(arr []int)[]int {
	//初始化
	var result[][]int
	for i := 0; i < len(arr); i ++ {
		result = append(result,[]int{arr[i]})
	}

	//排序
	for  len(result) > 1 {
		//记录索引
		length := len(result)

		//开始两两合并
		for i := 0; i < length; i += 2 {
			//判断
			if i + 1 < length {
				//合并
				result = append(result,mSort(result[i],result[i + 1]))
				continue
			}
			//当前加入
			result = append(result,result[i])
		}

		//更新result
		result = result[length:]
	}

	//返回
	return result[0]
}

/**
测试
 */
 func main() {
	 arr1 := []int{1,2,3,-1,0,6,-8}
	 //	arr2 := []int{-1,0,6}

	 fmt.Println(MergingSort(arr1))

 }

