package main

import "fmt"

/**
	优先队列（priority queue）：概念略---基于堆排序 而衍生出来的
	利用数组来实现优先队列  ---大顶堆和小顶堆
	给定下标i (1 <i < length) 则其双亲下标  floor(i / 2) --- 下标从1开始
	floor ((i - 1) / 2) ---下标从0开始

	API:-----go暂时用函数来实现
	buildPriorityQueue(arr []int) 利用切片来构造一个优先队列
	insertPriorityQueue(arr []int,x int) 向优先队列中插入一个元素
	getTopItem(arr []int) 获取最大关键字元素并删除
	getPriorityQueueLength(arr []int) 获取优先队列长度
	modifyPriorityQueue(arr []int, i,key int) 修改优先队列下标为i的值
 */
/**
	采用大顶堆---优先级高的先出队
 */
func buildPriorityQueue(arr []int) []int {
	//调节
	length := len(arr) - 1

	for i := length / 2; i >= 0; i -- {
		adjustPriorityQueue(&arr, i, length)
	}

	//返回优先队列
	return arr
}

/**
	添加元素
 */
func insertPriorityQueue(arr []int, x int) []int {
	//添加入队列
	arr = append(arr, x)

	//获取索引
	index := len(arr) -1

	//获取父
	p := (index - 1) / 2

	//修改
	for p > 0 && arr[p] < x {
		//父的值先去子
		arr[index] = arr[p]

		index = p

		p = (index - 1) /2
	}

	arr[index] = x


	//重新构建
	return buildPriorityQueue(arr)
}

/**
	获取最大关键字并删除
 */
func getTopItem(arr *[]int) int {
	//获取关键字
	key := (*arr)[0]

	//赋值最后一个
	(*arr)[0] = (*arr)[len((*arr))-1]

	//释放
	*arr = (*arr)[0 : len((*arr))-1]

	//调节---元素个数2个以上
	if len(*arr) > 1 {
		adjustPriorityQueue(arr, 0, len(*arr)-1)
	}

	//返回return
	return key
}

/**
	获取队列长度
 */
func getPriorityQueueLength(arr []int) int {
	return len(arr)
}

/**
	修改i的的值假设 不小于原来的值 向上查找，如果不原来的小 则向下查找
 */
func modifyPriorityQueue(arr []int, i, key int) []int {
	//判断索引
	if i > len(arr)-1 {
		//错误
		fmt.Println("错误！")
		return arr
	}

	//调节
	p := (i - 1) / 2  //双亲下标

	for p > 0 && arr[p] < key {
		//双亲赋值给 i
		arr[i] = arr[p]

		//继续寻找双亲
		i = p
		p = (p - 1) / 2

	}

	//赋值
	arr[i] = key

	//返回
	return arr
}

/**
	调节函数
 */
func adjustPriorityQueue(arr *[]int, low, end int) {
	//记录当前的元素的值
	key := (*arr)[low]

	//arr下标从0开始
	for j := (low * 2) + 1; j <= end; j = j*2 + 1 {
		//比较low 的孩子
		if (j+1) < end && (*arr)[j+1] > (*arr)[j] { //<
			j ++
		}

		//判断
		if (*arr)[j] <= key { //<=
			break
		}

		//交换
		(*arr)[low] = (*arr)[j]
		low = j
	}

	//赋值
	(*arr)[low] = key
}

func main() {
	//切片
	arr := []int{1, 4, 3, 6, 0, 7, 9, 16, -1}

	//构造优先队列
	queue := buildPriorityQueue(arr)

	//添加元素
	queue = insertPriorityQueue(queue, 19)
	queue = insertPriorityQueue(queue, 21)

	//获取头部元素
	fmt.Println(getTopItem(&queue))
	fmt.Println(getTopItem(&queue))

	//修改
	fmt.Println(modifyPriorityQueue(queue,1,17))


}
