package main

import "fmt"

/**
	斐波那契查找：
	原理：斐波那契查找与折半查找很相似，他是根据斐波那契数列的特点对有序表进行分割的。他要求开始表中记录的个数比某个斐波那契数小1，即n=F(K)-1;当记录不满足时，后面的元素都赋值为最后一个值。(注：使得n=F(K)-1是因为：如果表中的数据为F(K)-1个，mid分割又用掉一个，剩下F(K)-2个，正好分给两个子序列，每个子序列的长度为F(K-1)-1和F(K-2)-1，格式和之前的统一，方便递归编程实现)

mid = low + F(k-1) - 1
将value值与第mid位置的记录进行比较,比较结果分为三种：

相等，则返回mid
大于，low = mid + 1low=mid+1,k = k - 2k=k−2
小于，high = mid - 1high=mid−1,k = k - 1k=k−1
（说明：low=mid+1说明待查找的元素在[mid+1,high]范围内，k-=2 说明范围[mid+1,high]内的元素个数为F(k-2)-1个）
 */
/**
	斐波那契序列
	n <= 2  f[n] = 1
	n > 2   f[n] = f[n - 1]+ f[n - 2]
 */
/**
	@param arr []int 查询数组
	@param n int 查询范围0-n-1 之间
	@param key 关键字
 */
func FibonacciSearch(arr []int,n int,key int)int {
	//声明返回值
	rtIndex := -1

	//判断参数
	if len(arr) == 0 || n <= 0 {
		return rtIndex
	}

	//判断n
	if n > len(arr) {
		n = len(arr)
	}

	//根据n来构造斐波那契
	F := map[int]int{}

	//斐波那契序列
	F[0] = 0
	F[1] = 1
	k := 1

	//获取序列的k
	for F[k] - 1 < n {
		k++
		F[k] = F[k - 2] + F[k - 1]
	}

	//将数组扩展到F[k]-1的长度
	arr = arr[:n]
	for n < F[k] - 1 {
		arr = append(arr,arr[n - 1])
		n++
	}
	//定义起始
	low,high := 0,n - 1

	//开始查找
	for low <= high {
		//mid
		mid := low + F[k - 1] - 1

		//判断
		if arr[mid] > key {//左半长度为 F[k - 1] - 1
			high = mid - 1
			k --
		} else if arr[mid] < key {//右半长度为F[k-2]-1  （mid占一位) F[k] - 1 = F[k -1] + F[k -2] - 1 得到
			low = mid + 1
			k -= 2
		} else {
			//判断
			if mid < n {
				rtIndex = mid
				break
			} else {
				//表示填充的
				rtIndex = n - 1
				break
			}

		}
	}

	//返回
	return rtIndex

}

/**
	test
 */

 func main() {
 	arr :=[]int{1,2,3,4,5,6,7,8,9}

 	fmt.Println(FibonacciSearch(arr,6,5))

 }