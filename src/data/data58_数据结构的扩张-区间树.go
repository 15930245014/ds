package main

import "fmt"

/**
	区间树：就是在树结点的基础上添加一个区间对象（可以用结构体）和 max最大右区间属性
	任何两个区间都满足区间三分律：
	a i与j重叠。 i.low <= j.high && j.low <= i.high
	b i在j的左边  i.high < j.low
	c i 在j 的右边 i.low > j.high

	//用途---课程设置

*/
/**
	区间结构体
 */
type section struct {
	low, high int //区间的两端
}

/**
	定义区间红黑树结点信息
 */
type SectionRBTree struct {
	src               section //区间
	color             int     //颜色0-red 1-black
	lChild, rChild, p *SectionRBTree
	max               int //最大右端点  x.max = max(x.src,high,x.lChild.max,x.rChild.max)
	val int //默认为src.low
}

/**
	定义红黑树
 */
type SectionRedBlackTree struct {
	root *SectionRBTree
}

/**
	辅助函数---3个int 比较最大
 */
func maxSectionHigh(a, b, c int) int {
	max := a
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	return max
}

/**
	左旋转
 */
func leftSectionRBT(T *SectionRedBlackTree, x *SectionRBTree) {
	//获取当前的右孩子
	y := x.rChild

	//y左孩子指向
	x.rChild = y.lChild

	//判断
	if y.lChild != nil {
		y.p = x
	}

	y.p = x.p
	if x.p == nil { //root
		T.root = y
	} else if x.p.lChild == x {
		x.p.lChild = y
	} else {
		x.p.rChild = y
	}

	//修改指向
	y.lChild = x
	x.p = y

	//修改max
	y.max = x.max

	//判断
	left := 0
	if x.lChild != nil {
		left = x.lChild.max
	}
	right := 0
	if x.rChild != nil {
		right = x.rChild.max
	}
	x.max = maxSectionHigh(x.src.high, left, right)

}
 /**
 	右旋转
  */
func rightSectionRBT(T *SectionRedBlackTree, x *SectionRBTree) {
	//获取当前的左孩子
	y := x.lChild

	//修改x.lChild
	x.lChild = y.rChild

	//判断
	if y.rChild != nil {
		y.rChild.p = x
	}

	//修改双亲
	y.p = x.p
	if x.p == nil {
		T.root = y
	} else if x.p.lChild == x {
		x.p.lChild = y
	} else {
		x.p.rChild = y
	}

	y.rChild = x
	x.p = y

	//修改max
	y.max = x.max

	//判断
	left := 0
	if x.lChild != nil {
		left = x.lChild.max
	}
	right := 0
	if x.rChild != nil {
		right = x.rChild.max
	}
	x.max = maxSectionHigh(x.src.high, left, right)
}
/**
	插入区间树
 */
 func insertSectionRBTree(T *SectionRedBlackTree,src section) *SectionRedBlackTree{
 	//实例化新的节点
	newNode := new(SectionRBTree)
	newNode.val = src.low
	newNode.max = src.high
	newNode.src = src
	//判断
	if T.root == nil {
		//设为黑色
		newNode.color = 1
		T.root = newNode
		return T
	}

	//寻找位置
	p := T.root
	var x *SectionRBTree
	for p != nil {
		if p.max <= newNode.max {
			p.max = newNode.max
		}
		if p.val > src.low {
			x = p
			p = p.lChild
			continue
		}
		x = p
		p = p.rChild

	}

	//新结点插入位置
	if newNode.val < x.val {
		x.lChild = newNode
	} else {
		x.rChild = newNode
	}

	//调节颜色
	newNode.color = 0
	newNode.p = x

	//调节颜色
	adjustSectionRBTreeColor(T,newNode)

	//返回
	return T
 }

 /*
 	调节颜色
 */
 func adjustSectionRBTreeColor(T *SectionRedBlackTree,x *SectionRBTree) {
 	//当前的父为红色
 	for x.p != nil && x.p.p != nil && x.p.color == 0 {
 		//当前的父为左孩子
 		if x.p == x.p.p.lChild {
 			//获取叔叔结点
 			y := x.p.p.rChild

 			//判断
 			if y != nil && y.color == 0 { //叔叔结点为红色
 				y.color = 1
 				x.p .color = 1
 				x.p.p.color = 0
 				x = x.p.p
			}  else { //叔叔结点为红色
				if x == x.p.rChild {
					//以当前的父为中心左旋转
					leftSectionRBT(T,x.p)
				}

				//第三种情况
				x.p.color = 1
				x.p.p.color = 0

				//右旋转
				rightSectionRBT(T,x.p.p)
			}

		} else {//当前的父为右孩子
			//获取叔叔结点
			y := x.p.p.lChild

			//判断叔叔结点颜色
			if y != nil && y.color == 0 {//叔叔结点为红色
				y.color = 1
				x.p.color = 1
				x.p.p.color = 0
				x = x.p.p
			} else {
				//右旋转
				if x == x.p.rChild {
					rightSectionRBT(T,x.p)
				}

				//以祖父结点为根左旋转
				x.p.color = 1
				x.p.p.color = 0
				leftSectionRBT(T,x.p.p)

			}
		}
	}
 }
 /**
 	区间覆盖
  */
 func overSlap(x,y section) bool {
 	//区间三分定律
 	if x.low >= y.high || y.low >= x.high {
 		return false
	}
 	return true
 }

 /**
 	获取覆盖src的区间
  */
 func intervalSection(T * SectionRedBlackTree,src section) section{
 	p := T.root

 	//判断
 	for p != nil && !overSlap(p.src,src) { //判断是不是当前
 		//判断
 		if p.lChild != nil && p.lChild.max >= src.low { //在右侧
 			p = p.lChild
		} else { //在右侧
			p = p.rChild
		}
	}

 	return p.src
 }

/**
	测试函数
 */
func main() {
	//实例化红黑树
	rbTree := new (SectionRedBlackTree)

	//测试
	rbTree = insertSectionRBTree(rbTree,section{3,5});
	rbTree = insertSectionRBTree(rbTree,section{1,2});
	rbTree = insertSectionRBTree(rbTree,section{-1,1});
	rbTree = insertSectionRBTree(rbTree,section{3,6});
	rbTree = insertSectionRBTree(rbTree,section{1,3});

	fmt.Println(intervalSection(rbTree,section{5,6}))


}
