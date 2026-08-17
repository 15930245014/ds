package main

import (
	"fmt"
	"strconv"
)

/**
	桶排序 (Bucket sort)或所谓的箱排序，是一个排序算法，工作的原理是将数组分到有限数量的桶子里。每个桶子再个别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序）。桶排序是鸽巢排序的一种归纳结果。当要被排序的数组内的数值是均匀分配的时候，桶排序使用线性时间（Θ（n））。但桶排序并不是 比较排序，他不受到 O(n log n) 下限的影响。
 */

/**
	例：扑克牌 ♣<♠<♦<♥ 并且 2<3<4...A  如何将一个打乱的扑克牌排好序
	可以利用桶排序思想：划分4个桶 分别来获取的四种牌，然后各个桶排好序 最后连接起来就好了

	也可以 分为13个桶 每个同保存4种花色的牌 最后在按花色连接起来
 */

/**
	排序正整数
 */
 func bucketSort(arr []int) []int {
 	//声明二维数组来用作桶
 	var bucket [][]int

 	//初始化0-9
 	for i := 0; i <= 9; i ++ {
 		bucket = append(bucket,[]int{})
	}

 	//获取最大位数
 	max := arr[0]
 	for i := 1; i < len(arr); i ++ {
 		if max < arr[i] {
 			max = arr[i]
		}
	}

 	//获取最大位数
 	i := 10
 	digit := 1
 	for max >= i {
 		i *= 10
 		digit ++
	}

 	//桶排序---从个位开始
 	for i := 1; i <= digit ; i ++ {
 		//分配桶
 		for j := 0; j < len(arr); j ++ {
 			bucket[getDigit(arr[j],i)] = append(bucket[getDigit(arr[j],i)],arr[j])
		}

 		//重新排列
 		arr = []int{}
 		for k := 0; k < len(bucket); k ++ {
 			//判断
 			if len(bucket[k]) > 0 {
 				//排列
 				arr = append(arr,bucket[k]...)

 				//清空桶
				bucket[k] = []int{}
			}
		}

	}

 	//返回
 	return arr
 }
 /**
 	获取位数
 	1先转化成字符串
 	2、转换成数组
 	3、转换成整形
  */
 func getDigit(num int,digit int) int {
 	//转换成字符串
 	numStr := strconv.Itoa(num)

 	//判断长度
 	if len(numStr) < digit {
 		return 0
	}

 	//获取位数
 	rt := string(numStr[len(numStr) - digit])

 	//返回
 	intRT,_ := strconv.Atoi(rt)
 	return intRT

 }

func main() {
	arr := []int{999,23,45,9,98,83,3,55,44444,7}

	fmt.Println(bucketSort(arr))

}