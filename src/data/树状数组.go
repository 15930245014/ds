package 算法

import "fmt"

/*
 树状数组 下标从1 开始
 C[i] 负责的区间长度为 lowBit 其实 就是 最右边1 后边0的个数k 区间长度为 2^k
 寻找前驱： i - lowBit
 寻找祖先： i + lowBit
 适用于 单点修改 前缀和 时间均为O(logN)
*/

type TreeArr struct {
	C   []int
	arr []int
	n   int
}

func NewTreeArr(arr []int) *TreeArr {
	trArr := &TreeArr{
		C:   make([]int, len(arr)+1),
		n:   len(arr) + 1,
		arr: arr,
	}
	//建树
	for i := 0; i < len(arr); i++ {
		trArr.Update(i+1, arr[i])
	}

	return trArr
}

func (tr *TreeArr) Update(i int, v int) {
	for i < tr.n {
		tr.C[i] += v
		i += tr.lowBit(i)
	}
}

/*
prefix[1;i]
*/
func (tr *TreeArr) Query(i int) int {
	sum := 0
	for i > 0 {
		sum += tr.C[i]
		i -= tr.lowBit(i)
	}
	return sum
}

/*
	prefix[i,j]
*/

func (tr *TreeArr) RangeQuery(i, j int) int {
	return tr.Query(j) - tr.Query(i-1)
}

/*
直接设置值
*/

func (tr *TreeArr) Set(i, v int) {
	diff := v - tr.arr[i-1]
	tr.arr[i-1] = v
	tr.Update(i, diff)
}

func (tr *TreeArr) lowBit(x int) int {
	return -x & x
}

func t() {
	nums := []int{1, 2, 3, 4, 5}
	bit := NewTreeArr(nums)

	fmt.Println(bit.Query(3))         // 6 (1+2+3)
	fmt.Println(bit.RangeQuery(2, 4)) // 9 (2+3+4)

	bit.Update(3, 2) // nums[3] += 2

	fmt.Println(bit.Query(3))         // 8
	fmt.Println(bit.RangeQuery(2, 4)) // 11
}
