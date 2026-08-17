package main

import "fmt"

/**
	线段树
	SegmentTree :
    1、线段树是一棵二叉搜索树，它储存的是一个区间的信息。

    2、每个节点以结构体的方式存储，结构体包含以下几个信息：

     区间左端点、右端点；（这两者必有）

     这个区间要维护的信息（事实际情况而定，数目不等）。

	3、线段树的基本思想：二分。

	线段树的性质--平衡二叉树而且为完全二叉树
    1、每个节点的左孩子区间范围为[l，mid]，右孩子为[mid+1,r] 且 mid = (l + r) / 2

    2、对于结点k，左孩子结点为2*k，右孩子为2*k+1，这符合完全二叉树的性质
 */
//线段树数据结构
type segmentTree struct {
	l,r int //左区间和右区间
	sum int //附加属性---表示区间和
	p *segmentTree //父
	left,right *segmentTree //左孩子和右孩子
	lazy int //懒标记
}

type Segment struct {//跟
	root *segmentTree
}

/**
	实例以切片为主 求任意子数组的和
 */

/**
	根据切片初始化segmentTree
 */

func initSegment(l,r int,arr []int,p *segmentTree)*segmentTree {
	//实例化结点
	newNode := new(segmentTree)

	newNode.l = l
	newNode.r = r
	newNode.p = p
	//判断
	if l == r {
		newNode.sum = arr[l]

	} else {
		//构造子
		mid := (l + r) / 2

		//构造左右
		newNode.left = initSegment(l,mid,arr,newNode)
		newNode.right = initSegment(mid + 1,r,arr,newNode)

		//统计和
		newNode.sum = newNode.left.sum + newNode.right.sum
	}

	//返回
	return newNode
}

/**
	单点查询---查询叶子结点的区间
 */
func findX(x int,tree Segment) int {
	//获取当前
	p := tree.root

	//当前不为空
	for p != nil  {

		//获取中点
		mid := (p.l + p.r) / 2
		if p.l == p.r && mid == x {//当前为叶子节点并且和x相等
			//p.sum += p.lazy
			p.lazy = 0 //叶子节点置为0
			break
		}

		//判断---惰性
		if p.lazy > 0 {
			p.left.lazy += p.lazy
			p.right.lazy += p.lazy

			p.left.sum += (p.left.r - p.left.l + 1) * p.lazy
			p.right.sum += (p.right.r - p.right.l + 1) * p.lazy
			p.lazy = 0

		}


		//判断--标记下发
		if x > mid {
			p = p.right

		} else {
			p = p.left
		}
	}

	//返回
	return p.sum
}

/**
	查询区间的值
 */
func findXY(x,y int,tree *segmentTree) int {
	//[x,y]区间包含[l,r]区间
	if tree.l >= x && tree.r <= y {
		return tree.sum
	}

	//分发lazy
	if tree.lazy != 0 {
		tree.left.lazy += tree.lazy
		tree.right.lazy += tree.lazy
		tree.left.sum += (tree.left.r - tree.left.l + 1) * tree.lazy
		tree.right.sum += (tree.right.r - tree.right.l + 1) * tree.lazy
		tree.lazy = 0
	}

	//获取当前的l,r
	mid := (tree.l + tree.r) >> 1

	//区间三分律
	if x > mid {
		//[]右边
		return findXY(x,y,tree.right)


	} else if y <= mid {
		//[]左边
		return findXY(x,y,tree.left)

	} else {//相交
		return findXY(x,mid,tree.left) + findXY(mid,y,tree.right)
	}

}
/**
	区间数修改---定点修改--增加i
 */
func modifyX(x int,i int,tree *segmentTree)bool {
	//找到那个点在回溯
	if tree.r == tree.l {
		tree.sum += i
		return true
	}

	var res bool

	//获取中点
	mid := (tree.l + tree.r) / 2

	//位于左边
	if x <= mid {
		 res = modifyX(x,i,tree.left)

	} else {
		res = modifyX(x,i,tree.right)
	}

	//修改值
	tree.sum = tree.right.sum + tree.left.sum

	//返回
	return res
}
/**
	区间修改---同时修改区间值
 */
func modifyXY(x,y,val int,tree *segmentTree) bool{
	//判断当前是否在区间内
	if tree.l >= x && tree.r <= y {
		//叶子节点没有lazy
		if tree.l != tree.r {
			//更新lazy
			tree.lazy += val
		}

		//更新sum
		tree.sum += (tree.r - tree.l + 1) * val


		return true
	}

	//分发lazy
	if tree.lazy != 0 {
		tree.left.lazy += tree.lazy
		tree.right.lazy += tree.lazy
		tree.left.sum += (tree.left.r - tree.left.l + 1) * tree.lazy
		tree.right.sum += (tree.right.r - tree.right.l + 1) * tree.lazy
		tree.lazy = 0
	}

	//声明结果
	var res bool

	//获取mid
	mid := (tree.l + tree.r) / 2

	//区间三分率
	if mid >= y {

		res = modifyXY(x,y,val,tree.left)
		//tree.right.sum  += (tree.r - mid) * val
	} else if x > mid {
		res = modifyXY(x,y,val,tree.right)
		//tree.left.sum  += (mid - tree.l + 1) * val
	} else {//相交
		res = modifyXY(x,mid,val,tree.left) && modifyXY(mid + 1,y,val,tree.right)
	}

	//更新
	tree.sum = tree.left.sum + tree.right.sum


	//返回
	return res


}



func main() {
	//初始化切片
	arr := []int{1,5,1,3,4,2,0,9,0,9}

	//初始化线段树
	segMent := Segment{nil}
	segMent.root = initSegment(0,len(arr) - 1,arr,nil)

	//判断定点是否存在区间
	fmt.Println(findX(2,segMent))
	fmt.Println(modifyX(9,10,segMent.root))
	fmt.Println(modifyX(3,11,segMent.root))

	//区间求和
	fmt.Println(findXY(2,9,segMent.root))

	//区间修改
	fmt.Println(modifyXY(2,4,2,segMent.root))
	fmt.Println(modifyXY(4,6,1,segMent.root))
	fmt.Println(modifyXY(1,4,3,segMent.root))

	//区间求和
	fmt.Println(findXY(2,9,segMent.root))
	fmt.Println(findX(4,segMent))
	fmt.Println(findXY(2,9,segMent.root))


}







