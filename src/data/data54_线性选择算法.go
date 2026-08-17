package main

import (
	"fmt"
	"math/rand"
	"time"
)
/**
	线性选择算法：就是利用分治的策略来实现返回数组的任意位置的记录，并且时间复杂度为线性O(n)
 */

 /**
 	@param arr []int 数组
 	@param i int  i为位置--其他转化一下即可
  */
 func magicSelect(arr []int,i int) int  {
 	//第i小
 	i = i - 1
 	//第i大
 	i = len(arr) - i

	return randomized_select(arr,0,len(arr)-1,i)
 }

 /**
 	核心方法
  */
 func randomized_select(arr []int,start,end,i int)int {
 	//判断
 	if start == end {
 		return arr[start]
	}

 	//随机分割
 	pivot := randPivot(&arr,start,end)

 	//判断
 	if i == pivot {
 		return arr[pivot]
	} else if i > pivot { //后半部分
		return randomized_select(arr,pivot + 1,end,i)
	} else {//前半部分
		return randomized_select(arr,start,pivot - 1,i)
	}

 }
/**
	随机快速排序
 */
 func randPivot(arr *[]int,s,l int) int{
 	//随机杻轴
 	random := RandInt(s,l)

 	//快速排序划分数值
 	key := (*arr)[random]

 	//交换
	 (*arr)[s],(*arr)[random] = (*arr)[random],(*arr)[s]


 	//划分区间
 	for s < l {
 		//找到一个比key小的
 		for s < l && (*arr)[l] >= key {
 			l --
		}
 		//赋值
		(*arr)[s] = (*arr)[l]

		//从前面找到比keydade
		for s < l && (*arr)[s] <= key {
			s++
		}

		//交换
		(*arr)[l] = (*arr)[s]

	}
	 (*arr)[s] = key

	 //返回
	 return s

 }
 /**
 	随机数
  */
func RandInt(min, max int) int {
	//设置
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
func main() {
	arr := []int{1,2,8,4,5,6,7,3}
	fmt.Println(magicSelect(arr,1))

}