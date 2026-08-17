package main

import (
	"fmt"
	"math"
)

/**
	斐波那契堆：斐波那契堆是具有一系列的最小堆序（min-heap ordered）的有根树的集合。在斐波那契堆中所有的树根都用其left和right 指针成一个环形的双向链表。并且有一个min指针指向最小关键字把这个结点称为斐波那契堆的最小结点。如果斐波那契堆为空则min == nil,斐波那契堆的结点还包含记录其孩子数目的度degree。并还有一个bool值mark表示结点从上一次成为另一个结点的孩子后是否失去过孩子。新结点是未标记的，结点的孩子结点通过left和right连接成一个环形的双向链表，称为孩子链表

	用途：优先队列等
*/

/**
	构造斐波那契结点结构体
 */
type fibonacciNode struct {
	degree         int            //当前结点的孩子数目
	child          *fibonacciNode //指向某一孩子结点
	left, right, p *fibonacciNode //左右指针,双亲指针
	mark           bool           //是否标记
	key            int            //结点值
}

/**
	斐波那契堆
 */
type fibonacciHeap struct {
	n   int            //结点数目
	min *fibonacciNode //最小结点
}

/**
	初始化一个斐波那契堆
 */
func initFibonacciHeap() *fibonacciHeap {
	//声明指针
	H := new(fibonacciHeap)

	//初始化参数
	H.n = 0
	H.min = nil

	//返回
	return H
}

/**
	将key插入到斐波那契堆中
 */
func insetFibonacciHeap(H *fibonacciHeap, key int) *fibonacciHeap {
	//实例化新结点
	newNode := new(fibonacciNode)
	newNode.p = nil
	newNode.child = nil
	newNode.degree = 0
	newNode.key = key
	newNode.mark = false

	//判断斐波那契堆
	if H.min == nil {
		//构造环
		newNode.right = newNode
		newNode.left = newNode
		H.min = newNode

	} else { //H.min右侧添加
		newNode.right = H.min.right
		H.min.right.left = newNode
		newNode.left = H.min
		H.min.right = newNode

		//比较--指向较小的
		if H.min.key > key {
			H.min = H.min.right
		}

	}

	//结点数目+1
	H.n ++

	//返回
	return H

}

/**
	合并斐波那契堆
 */
func unionFibonacciHeap(h1 *fibonacciHeap, h2 *fibonacciHeap) *fibonacciHeap {
	//合并斐波那契
	if h1.min == nil {
		return h2
	}
	if h2.min == nil {
		return h1
	}

	//h1和h2 均不为空
	h1.n += h2.n

	//连接循环双向链表h1.min连接h2.min
	rNode := h1.min.right //获取h1的右侧
	lNode := h2.min.left  //获取h2的左侧

	//连接
	h1.min.right = h2.min
	h2.min.left = h1.min

	lNode.right = rNode
	rNode.left = lNode

	//比较
	if h1.min.key > h2.min.key {
		h1.min = h1.min.right
	}

	//返回合并后的
	return h1
}

/**
	抽取最小结点--并删除
 */
func getTopHeapItem(H *fibonacciHeap) *fibonacciNode {
	//获取最小元素
	z := H.min

	//判断
	if H.min != nil {
		//判断是否有孩子
		if H.min.degree > 0 && H.min.child != nil {
			//将孩子结点添加到根链表中
			p := H.min.child

			for i := 1; i <= H.min.degree && p != nil; i ++ {
				//修改双亲
				p.p = nil

				//记录向下结点
				col := p.right

				//添加结点
				p.right = H.min.right
				H.min.right.left = p
				H.min.right = p
				p.left = H.min

				//向下遍历
				p = col
			}
		}

		//判断
		if H.min.right == H.min {
			H.min = nil
		} else {
			//双向链表移除H.min
			H.min.left.right = H.min.right
			H.min.right.left = H.min.left

			H.min = H.min.right

			//按度数合并
			consolidate(H)
		}
	}

	//结点数目-1
	H.n -= 1

	//返回
	return z

}

/**
	调整斐波那契堆
 */
func consolidate(H *fibonacciHeap) *fibonacciHeap {

	//构造一个辅助数组lgn
	var dp [] *fibonacciNode

	//初始化
	for i := 0; i <= H.n; i ++ {
		dp = append(dp, nil)
	}

	//遍历根链表
	p := H.min
	dp[p.degree] = p

	//指向下一个
	p = p.right

	for p != H.min {
		k := p.right
		//获取度
		x := p

		//获取度
		degree := x.degree
		//合并
		for dp[degree] != nil {
			//获取元素
			y := dp[degree]

			//判断值域
			if x.key > y.key {
				//交换y,x---交换值域指针未交换
				tmp := *y
				*y = *x
				*x = tmp

				//修改指针确保x正确即可
				//x.left = y.left
				//x.right = y.right

			}

			//删除 y并且成为x的子
			//1、移除y
			/*y.left.right = y.right
			y.right.left = y.left*/

			//2、将y成为x的子
			if x.child == nil {
				x.child = y
				y.right = y
				y.left = y
				y.mark = false
			} else {
				y.right = x.child.right
				x.child.right.left = y

				y.left = x.child
				x.child.right = y
				y.p = x
			}

			//3、x的度+1
			x.degree ++
			y.mark = false

			dp[degree] = nil
			degree ++
		}
		//写入合并后的x
		dp[degree] = x

		//继续循环
		//p = x.right
		p = k
	}

	//重置
	H.min = nil

	for i := 0; i < len(dp); i ++ {
		//判断
		if dp[i] == nil {
			continue
		}

		//构造斐波那契
		if H.min == nil {
			H.min = dp[i]
			H.min.left = H.min
			H.min.right = H.min

		} else { //添加到斐波那契堆
			dp[i].right = H.min.right
			H.min.right.left = dp[i]
			H.min.right = dp[i]
			dp[i].left = H.min

			//更新最小
			if H.min.key > dp[i].key {
				H.min = dp[i]
			}

		}

	}

	return H
}

/**
	交换斐波那契结点
	*交换双向链表的任意2个结点
	1、x与y相邻（x在y左边，x在y的右边）
	2、x与y不相邻
 */
func swapFibonacciNode(x, y *fibonacciNode) {
	//判断x,y的
	if x == y.left { //x -> y
		//处理
		x.left.right = y
		y.left = x.left

		y.right.left = x
		x.right = y.right

		//x与y相连
		y.right = x
		x.left = y

	} else if x == y.right { // y -> x
		y.left.right = x
		x.left = y.left

		x.right.left = y
		y.right = x.right

		x.right = y
		y.left = x

	} else {

		//交换结点x和y的顺序
		tmpLeftNode := x.left
		tmpRightNode := x.right

		//将x换到y
		y.left.right = x
		x.left = y.left
		x.right = y.right
		y.right.left = x

		//将y换到x
		tmpLeftNode.right = y
		y.left = tmpLeftNode
		y.right = tmpRightNode
		tmpRightNode.left = y
	}
}


/**
	将斐波那契堆的结点减值
 */
func decreaseFibonacci(H *fibonacciHeap,x *fibonacciNode, key int) {
	//判断
	if key > x.key {
		fmt.Println("错误值不能小于当前值！")
		return
	}
	//赋值
	x.key = key


	//判断当前值是否为根root
	if x.p != nil && key < x.p.key { //破坏了最小堆原则
		y := x.p
		cutFibonacciNode(H,y,x)

		//向上递归
		/*z := y.p
		for z != nil && y.mark == true {
			cutFibonacciNode(H,z,y)
			y = z
			z = y.p
		}
		y.mark = false*/
		fixedUpFibonacciHeap(H,y) //---与上面等价
	}

	//修改最小值
	if H.min.key > x.key {
		H.min = x
	}

}
/**
	自底向上调节
 */
func fixedUpFibonacciHeap(H *fibonacciHeap,y *fibonacciNode) {
	z := y.p
	if z != nil {
		if y.mark == false {
			y.mark = true
		} else {
			cutFibonacciNode(H,z,y)
			fixedUpFibonacciHeap(H,z)
		}

	}
}

/**
	删除斐波那契堆结点x
 */
func deleteFibonacciHeap(H *fibonacciHeap,x *fibonacciNode) {
	//将结点x减值为-无穷
	decreaseFibonacci(H,x,math.MinInt64)

	//获取最小结点并移除
	getTopHeapItem(H)
}

/**
	将x从y移除并将x放入根列表中
 */
func cutFibonacciNode(H *fibonacciHeap,y,x *fibonacciNode) {
	//移除x
	if y.degree == 1 {
		y.child = nil

	} else {
		x.left.right = x.right
		x.right.left = x.left
		y.child = x.left //重定向 y.child = x 这种情况
	}
	y.degree --

	//将x 放入根列表
	 x.right = H.min.right
	 H.min.right.left = x
	 H.min.right = x
	 x.left = H.min
	 x.p = nil
	 x.mark = false
}



/**
	test fibonacciHeap
 */
func main() {
	//初始化斐波那契堆
	H := initFibonacciHeap()

	//插入
	H = insetFibonacciHeap(H, 2)
	H = insetFibonacciHeap(H, 1)
	H = insetFibonacciHeap(H, 3)
	H = insetFibonacciHeap(H, 4)
	H = insetFibonacciHeap(H, -2)
	H = insetFibonacciHeap(H, -9)
	H = insetFibonacciHeap(H, 19)

	//交换斐波那契堆结点
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)
	fmt.Println(getTopHeapItem(H).key)

}
