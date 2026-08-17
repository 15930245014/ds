package main

import "fmt"

/**
	顺序统计树（order-statistic tree）：用于快速的找到一个集合中第i小的数（之前的线性选择算法（时间复杂度O(n)）），采用红黑树来实现,时间复杂度为O(log2n)
	概念：
	元素的秩：即一个元素在集合线性中的位置。

	在红黑树的color ,val,lChild,rChild,p 的基础上添加一个size属性  哨兵nil.size = 0
	结点x的size = x.left.size + x.right.size + 1(本身)
 */

/**
	红黑树的结点
 */
type orderBTSTree struct {
	val               int           //关键字
	color             int           //颜色 0-red 1-black
	lChild, rChild, p *orderBTSTree //孩子结点和双亲结点
	size              int           //结点数
}

/**
	顺序统计树
 */
type orderRedBlackTree struct {
	root *orderBTSTree
}

/**
	左旋转,以x为中心左旋转
 */
func leftRotateOrderBTS(T *orderRedBlackTree, x *orderBTSTree) {
	//获取当前的右孩子
	y := x.rChild

	//修改指向
	x.rChild = y.lChild

	//修改y.lChild双亲
	if y.lChild != nil {
		y.lChild.p = x
	}

	y.p = x.p
	//修改y双亲指向
	if x.p == nil {
		T.root = y
	} else if x.p.lChild == x {
		x.p.lChild = y
	} else {
		x.p.rChild = y
	}

	//修改x指向
	y.lChild = x
	x.p = y

	//修改size
	y.size = x.size

	leftSize := 0
	if x.lChild != nil {
		leftSize = x.lChild.size
	}
	rightSize := 0
	if x.rChild != nil {
		rightSize = x.rChild.size
	}
	x.size = leftSize + rightSize + 1
}

/**
	右旋转---以x为中心
 */
func rightRotateOrderBTS(T *orderRedBlackTree, x *orderBTSTree) {
	//获取当前的左孩子
	y := x.lChild

	x.lChild = y.rChild
	//修改指向
	if y.rChild != nil {
		y.rChild.p = x
	}

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

	//修改size
	y.size = x.size

	leftSize := 0
	if x.lChild != nil {
		leftSize = x.lChild.size
	}
	rightSize := 0
	if x.rChild != nil {
		rightSize = x.rChild.size
	}
	x.size = leftSize + rightSize + 1
}

/**
	拆入数据默认新结点的颜色是红色的
 */
func insertOrderBTSTree(T *orderRedBlackTree, key int) *orderRedBlackTree {
	//获取根
	p := T.root

	//生成新结点
	newNode := new(orderBTSTree)

	//新结点的size 默认为1 nil(0) + nil(0) + 1
	newNode.size = 1
	newNode.val = key

	//判断
	if p == nil { //插入结点为根
		//颜色设为黑色
		newNode.color = 1

		T.root = newNode

		//返回
		return T
	}

	//查找双亲父位置
	var x *orderBTSTree
	for p != nil {
		//p.size += 1---维护size
		if p.val == key {
			break
		}
		if p.val > key { //在左子树查找
			x = p
			p = p.lChild
			continue
		}
		x = p
		p = p.rChild //在右子树查找
	}

	//判断
	if x == nil { //表示存在了
		return T
	}

	//拆入新结点
	newNode.color = 0 //默认红色
	newNode.p = x
	if x.val > key { //新结点在双亲的左子树位置上
		x.lChild = newNode
	} else {
		x.rChild = newNode
	}

	//自底向上zize + 1
	z := x
	for z != nil {
		z.size += 1
		z = z.p
	}


	//调节颜色
	adjustOrderBTSTreeColor(T,newNode)

	//返回
	return T
}

/**
	调节颜色
 */
func adjustOrderBTSTreeColor(T *orderRedBlackTree, x *orderBTSTree) {
	//判断当前父的颜色为红色
	for x.p != nil && x.p.color == 0 && x.p.p != nil{
		//当前的父为祖父的左孩子
		if x.p == x.p.p.lChild {
			//获取叔叔结点
			y := x.p.p.rChild

			//判断颜色
			if y != nil && y.color == 0 { //叔叔结点为红色,叔叔和双亲设为黑色祖父设为红色 继续向上
				y.color = 1
				x.p.p.color = 0
				x.p.color = 1
				x = x.p.p

			} else {
				if y != nil && x == x.p.rChild { //叔叔为黑色，x在双亲的右子树
					//左旋转
					leftRotateOrderBTS(T, x.p)

				}

				//将祖父设置为红色，双亲设置为黑色
				x.p.color = 1
				x.p.p.color = 0
				rightRotateOrderBTS(T, x.p.p)
			}

		} else {
			//获取叔叔结点
			y := x.p.p.lChild

			//判断
			if y != nil && y.color == 0 { //叔叔是红色的
				y.color = 1
				x.p.color = 1
				x.p.p.color = 0
				x = x.p.p

			} else  {
				if x == x.p.lChild {
					//右旋转
					rightRotateOrderBTS(T,x.p)
				}
				//叔叔是黑色
				x.p .color = 1
				x.p.p .color = 0
				leftRotateOrderBTS(T,x.p.p)
			}
		}

	}
	//根设置为黑色
	T.root.color = 1
}

/**
	获取指定秩的元素的值
 */
 func getOrderBTSTree(T *orderBTSTree, i int) *orderBTSTree{
 	if T != nil {
 		//获取秩
 		left := 0
 		if T.lChild != nil {
 			left = T.lChild.size
		}
 		r := left + 1

 		//判断
 		if r == i {
 			return T
		}

 		if r > i {
 			return getOrderBTSTree(T.lChild,i)
		} else {
			return getOrderBTSTree(T.rChild,i - r)
		}

	} else {
		return nil
	}

 }
 /**
 	根据结点来获取指定的秩
  */
func getOrderTntBTS(x *orderBTSTree)int{
	//循环
	left := 0
	if x.lChild != nil {
		left = x.lChild.size
	}
	r := left + 1

	for x != nil {
		if x.p != nil && x == x.p.rChild {
			left = 0
			if x.p.lChild != nil {
				left = x.p.lChild.size
			}
			r += left + 1
		}
		x = x.p
	}

	return r
}

func blOrderBTSTree(t *orderBTSTree) {
	if t != nil {
		blOrderBTSTree(t.lChild)
		fmt.Println(t.val)
		blOrderBTSTree(t.rChild)
	}
}

type MedianFinder struct {
	tree *orderRedBlackTree //顺序统计树
	n int //当前数量
}


/** initialize your data structure here. */
func Constructor5() MedianFinder {
	this := new(MedianFinder)
	this.tree.root = nil
	this.n = 0

	return *this
}


func (this *MedianFinder) AddNum(num int)  {
	this.tree = insertOrderBTSTree(this.tree,num)
	this.n ++
}


func (this *MedianFinder) FindMedian() float64 {
	//获取值
	if this.n % 2 == 0 {//偶数
		mid1 := this.n / 2
		mid2 := mid1 + 1

		x := float64(getOrderBTSTree(this.tree.root,mid1).val)
		y := float64(getOrderBTSTree(this.tree.root,mid2).val)

		return  (x + y)/ 2

	} else {
		mid := (this.n + 1) / 2

		return float64(getOrderBTSTree(this.tree.root,mid).val)
	}
}


func main() {
	var RBTree = new(orderRedBlackTree)
	RBTree = insertOrderBTSTree(RBTree, 1)
	RBTree = insertOrderBTSTree(RBTree, 2)
	RBTree = insertOrderBTSTree(RBTree, 3)
	RBTree = insertOrderBTSTree(RBTree, 10)
	RBTree = insertOrderBTSTree(RBTree, 7)
	RBTree = insertOrderBTSTree(RBTree, 5)
	RBTree = insertOrderBTSTree(RBTree, 6)
	RBTree = insertOrderBTSTree(RBTree, 8)
	RBTree = insertOrderBTSTree(RBTree, 9)
	RBTree = insertOrderBTSTree(RBTree, 23)
	RBTree = insertOrderBTSTree(RBTree, 13)
	RBTree = insertOrderBTSTree(RBTree, 12)
	RBTree = insertOrderBTSTree(RBTree, 19)
	RBTree = insertOrderBTSTree(RBTree, 27)
	RBTree = insertOrderBTSTree(RBTree, 89)
	RBTree = insertOrderBTSTree(RBTree, 45)
	RBTree = insertOrderBTSTree(RBTree, 54)
	RBTree = insertOrderBTSTree(RBTree, 26)
	RBTree = insertOrderBTSTree(RBTree, 29)
	RBTree = insertOrderBTSTree(RBTree, 34)

	//中序遍历
	x := getOrderBTSTree(RBTree.root,19)

	fmt.Println(getOrderTntBTS(x))
}
