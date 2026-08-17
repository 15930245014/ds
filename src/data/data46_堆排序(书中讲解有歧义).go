package main

import "fmt"

/**
	堆排序（heap sort）：堆排序算法是floyd (罗伯特.弗洛伊德) 和 Williams在1964年共同发明的，同时他们发明了“堆”这样的数据结构
	堆是具有下列性质的完全二叉树：每个结点的值都大于或等于其左右孩子的结点的值成为大顶堆 或者每个结点的值都小于或等于其左右孩子的值成为小顶堆
	堆的特性：
		若果数组的序列从下表1开始 则 下标为i的结点的左孩子下标为2*i 右孩子的下标为2*i + 1
		若数组的下标从0开始，则下标为i的结点的左孩子下表为2*i + 1 右孩子为 2*(i+1)
	堆排序思路 ---升序利用大顶堆 降序利用小顶堆
		将待排序的序列构造一个大顶堆，此时整个序列的最大值就是堆顶的根结点。将它移走（其实就是将其与数组的末尾元素交换，此时末尾元素就是最大值），然后将剩余的n-1个序列重新构造一个堆，这样就会得到n个元素中的此最大值。如此反复便能得到一个有序序列了
 */

 /**
 	堆排序思路：heapSort
 	1、初始化---构建完全二叉堆（大顶堆）
 	2、循环交换
 	3、余下构建大顶堆
 4、-----当前案例假设数组下标从1开始
  */

  func heapSort(arr []int)[]int {
  	//判断长度
  	if len(arr) == 0 || len(arr) == 1 {
  		//返回
  		return arr
	}

  	//初始化构造大顶堆---下标从1开始
  	/*
  	for l := len(arr)-1 /2; l >= 1; l-- {
  		arr = heapAdjust(arr,l,len(arr)-1)
	}
  	*/

  	//初始化下标从0开始
  	for l := (len(arr)-1)/2; l >= 0; l-- {// len(arr)/2 同样适用 但len(arr) - 1更准确
  		arr = heapAdjust(arr,l,len(arr)-1)
	}


  	//进行排序 i从最后一个元素开始
  	for i :=  len(arr) - 1; i >= 1; i -- {//下标从0开始 到 i>=1 下标从1开始  到 i>1

  		//交换i 和 len(arr) - i
		arr[0], arr[i] = arr[i] ,arr[0] //下标从0开始 交换arr[0] 下标从一开始交换arr[1]

		//调整
		heapAdjust(arr,0,i - 1) //下标从0开始  i-1 因为i已经为最大元素 找次最大
		//heapAdjust(arr,1,i - 1) //下标从1开始
	}

  	return arr

  }
  /**
  	调节是指成为大顶堆
  	*param arr []int 原始数组
  	*param start int 调节开始元素下标---对顶元素下标
  	*param end int 调节结束元素小标
   */
  func heapAdjust(arr []int,start int,end int) []int {
  		//先记录当前堆顶元素
  		key := arr[start]

  		//循环调节
  		for j := start * 2 + 1; j <= end; j = j * 2+1 { //此时j为当前栈顶的左孩子下标
  			//获取当前的左右孩子较大的下标
  			if j < end && arr[j + 1] > arr[j] {
  				j ++
			}

  			//判断key
  			if key >= arr[j] {
  				break
			}

  			//继续下一次
  			//1、先让j元素占据start位置 ，j位置在判断其子与key 来决定j该用谁来填充
  			arr[start] = arr[j]

  			//此时start位置轮空
  			start = j

		}

  		//写入空的start
  		arr[start] = key

  		//返回
  		return arr
  }

/**
 测试堆排序
  */
func main() {
	//声明数组
	arr := []int {-7,1,44,56,2,4,6,8,-1,-9,-10,-90,900,488,39,-100}

	//插入排序
	fmt.Println(heapSort(arr))
}