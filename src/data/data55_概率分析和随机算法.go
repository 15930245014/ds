package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	随机算法：一般地如果一个算法的行为不仅由输入决定，而且也有随机数生成器产生的数值决定的，则称这个算法是随机的（randomized）;在实践中绝大多数编程环境会提供一个伪随机数生成器它是一个确定性算法返回值在统计上看起来是随机的。
	随机构造二叉搜索树：每次从切片插入数据时，随机当前切片的元素
 */

/**
	随机排列数组实现---描述随机打乱一个数组（arr）：
	辅助一个数组p[i] = rand(1,i^3)
	1、利用随机数构造数组p
	2、根据p的值来对数组arr排序进而随机
	3、测试结果不满足均匀随机分布
 */
func randomArray(arr []int) []int {
	//声明p
	p := make([]int, len(arr))
	copy(p, arr)

	//构造p
	for i := 0; i < len(arr); i++ {
		p[i] = randInt(0, i*i*i)

	}

	//采用折半插入排序
	for i := 1; i < len(arr); i ++ {
		//记录key
		key := arr[i]
		pkey := p[i]

		//判断
		if p[i] < p[i-1] {
			//起止下标
			low, high := 0, i - 1

			//寻找插入位置
			for low <= high {
				//折半
				mid := (high + low) / 2

				//判断
				if p[mid] <= pkey {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

			//low即为位置右移
			for j := i; j > low; j-- {
				arr[j] = arr[j - 1]
				p[j] = p[j - 1]
			}

			//写入
			arr[low] = key
			p[low] = pkey

		}
	}
	//返回
	return arr
}
/**
	均匀随机分布

 */
 func randomizedArray(arr []int) []int{
 	//循环
 	for i := 0; i < len(arr);i++ {
 		//交换
 		random := randInt(i,len(arr) - 1)
 		arr[i],arr[random] = arr[random],arr[i]
	}

 	return arr
 }

 /**
 	需要马上给出最后
 	在线雇佣问题
	k 是一个概率学 当 k = n/e时 有e/1的概率能录取到好的员工
  */
  func onLineMaximum(score []int,k,n int) int {
	//找出k之前最好的并不录取
	bestScore := score[0]
	for i := 0; i <= k;i++ {
		if bestScore < score[i] {
			bestScore = score[i]
		}
	}

	//录取k之后的第个
	for i := k + 1; i <= n; i ++ {
		if score[i] >= bestScore {
			return i
		}
	}
	//录用最后一个
	return n

  }
/**
	随机数
 */
func randInt(min, max int) int {
	//设置
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func main() {
	//普通随机思想
	arr := []int{1, 2, 3, 4, 5, 6, 7, 7, 9}
	fmt.Println(randomArray(arr))

	//满足均匀随机分布
	fmt.Println(randomizedArray(arr))
}
