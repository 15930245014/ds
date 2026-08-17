package main

import (
	"fmt"
	"go/types"
)

/**
	红黑树（red_black_tree）：是一种和AVL树作用相同 但并不向平衡树二叉树那样对高度那么苛刻，红黑树确保没有一条路径会比其他路径长出2倍，是近似于平衡的。红黑树也是解决二叉搜索树不平衡问题的
	红黑树满足的条件：
	1、每一个结点要么是黑色要么是红色的
	2、根结点是黑色
	3、叶节点（nil）是黑色
	4、如果双亲结点为红色，孩子结点必定为黑色
	5、从当前结点到其子结点的简单路径的黑色结点数目都相同
 */

/*
	红黑树的插入规则：
	1、新插入的结点默认为红色 即 color = 0
	2、如果父结点为黑色满足红黑树特性不做任何处理
	3、根节点默认为黑色
	4、-------父结点颜色为红色分为2种大情况
		A、父结点为祖父节点的左孩子：
		  1、叔叔结点为红色，叔叔和双亲结点颜色设置为黑色，以祖父结点自底向上验证
		  2、叔叔结点为黑色，当前结点为双亲结点的左孩子，将双亲结点设置为黑色，祖父设置为红色，然后以祖父结点为中心右旋转
		  3、叔叔结点为黑色,当前结点为双亲结点的右孩子，先以双亲结点左旋转，然后执行步骤2
		B、父结点为祖父结点右孩子：
		 1、叔叔结点为红色，叔叔和双亲结点颜色设置为黑色，以祖父结点自底向上验证
		  2、叔叔结点为黑色，当前结点为双亲结点的右孩子，将双亲结点设置为黑色，祖父设置为红色，然后以祖父结点为中心左旋转
		  3、叔叔结点为黑色,当前结点为双亲结点的左孩子，先以双亲结点右旋转，然后执行步骤2

 */

/**
	红黑树结点结构体
 */
type RBTree struct {
	color             int     //0,代表红色，1代表黑色
	lChild, rChild, p *RBTree //左孩子、右孩子，双亲结点
	val               int     //值
}

/**
	红黑树
 */
type RedBlackTree struct {
	root *RBTree //根结点
	nil  types.Nil
}

/**
	辅助函数--左旋转
	以当前结点进行左旋转
 */
func lRotate(t *RBTree, T *RedBlackTree) {
	//获取右
	rChild := t.rChild

	//T rChild指向
	t.rChild = rChild.lChild

	//判断
	if rChild.lChild != nil {
		//修改双亲
		rChild.lChild.p = t
	}

	//修改父指向
	rChild.lChild = t
	if t.p == nil {
		T.root = rChild
	} else {
		if t.p.lChild == t {
			t.p.lChild = rChild
		} else {
			t.p.rChild = rChild
		}
	}
	//修改rChild双亲
	rChild.p = t.p
	t.p = rChild

}

/**
 右旋转
 */
func rRotate(t *RBTree, root *RedBlackTree) {
	//获取t的左孩子
	y := t.lChild

	//修改指向
	t.lChild = y.rChild

	//判断
	if y.rChild != nil {
		//修改双亲
		y.rChild.p = t
	}

	//修改y指向
	y.rChild = t

	//判断
	if t.p == nil {
		root.root = y
	} else {
		if t.p.lChild == t {
			t.p.lChild = y
		} else {
			t.p.rChild = y
		}
	}

	//修改y的双亲
	y.p = t.p
	t.p = y

}

/**
	调节红黑树
 */
func adjustRBTree(z *RBTree, root *RedBlackTree) {
	//当前结点的父为红色时调节，如果为黑色满足特性
	for z.p != nil && z.p.color == 0 && z.p.p != nil { //实际上添加z.p.p 和z.p != nil
		//当前的父为祖父的左孩子
		if z.p == z.p.p.lChild {
			//获取叔叔结点
			y := z.p.p.rChild

			//判断叔父的颜色
			if y != nil && y.color == 0 { //为红色时 --将祖父设为红色，叔父同时设为黑色 条件 y!= nil
				y.color = 1
				z.p.color = 1
				z.p.p.color = 0

				//递归向上调节
				z = z.p.p

			} else { //叔父的颜色为黑色
				//判断z在双亲的位置
				if z == z.p.rChild {
					z = z.p
					//以z.p左旋转
					lRotate(z, root)
				}
				//将祖父设置为红色，双亲设置为黑色
				z.p.color = 1
				z.p.p.color = 0

				//以z.p.p右旋转
				rRotate(z.p.p, root)

			}
		} else { //当前的父为祖父的右孩子
			//获取叔叔结点
			y := z.p.p.lChild

			//判断叔叔结点颜色
			if y != nil && y.color == 0 { //叔叔结点为红色----将双亲和叔叔同时设置为黑色祖父设置为红色并向上递归
				y.color = 1
				z.p.color = 1
				z.p.p.color = 0
				z = z.p.p
			} else { //叔叔结点的颜色为黑色
				//判断z在双亲结点的位置
				if z == z.p.lChild {
					z = z.p
					//右旋转
					rRotate(z, root)
				}

				//双亲和祖父颜色
				z.p.color = 1
				z.p.p.color = 0

				//左旋转
				lRotate(z.p.p, root)
			}
		}
	}
	//根设置为黑色
	root.root.color = 1

}

/**
	插入红黑树
	1、寻找插入位置
	2、将新节点设置为红色
	3、调节
 */
func insertRBTree(T *RedBlackTree, key int) *RedBlackTree {
	//构造新结点
	newNode := new(RBTree)
	newNode.val = key
	newNode.color = 0 //默认为红色

	//判断
	if T.root == nil {
		newNode.color = 1 //根结点为黑色
		T.root = newNode
		return T
	}

	//利用循环来寻找位置
	p := T.root
	var pre *RBTree

	for p != nil {
		//判断是否相等
		if p.val == key {
			pre = p
			break
		}

		//当前结点比key大
		if p.val > key {
			pre = p
			p = p.lChild //向左边寻找
			continue
		}

		//当前结点比key小
		pre = p
		p = p.rChild
	}
	//判断pre
	if pre.val == key {
		//不做任何处理
		return T
	}

	//判断当前结点在双亲的位置
	if pre.val > key {
		pre.lChild = newNode
	} else {
		pre.rChild = newNode
	}

	//设置父
	newNode.p = pre

	//调节
	adjustRBTree(newNode, T)

	//返回
	return T
}

/**
	中序遍历红黑树
 */
func blRBTree(T *RBTree) {
	if T != nil {
		//遍历左
		blRBTree(T.lChild)

		//输出当前
		fmt.Println(T.val)

		//遍历右
		blRBTree(T.rChild)

	}
}

/**
	红黑树的删除----思路同于二叉搜索树
	1、根据关键字先找到要删除的结点
		1、如果要删除的结点只有一个孩子，子承父业
		2、如果没有孩子 则当前的父的左或右孩子设为nil 可以合并和1
		3、左右孩子均有则从右子树找到最小的篡位
	2、调节颜色
 */

   func deleteRBTree(T *RedBlackTree,key int) {
   	//根据关键字查找要删除的结点
   	z := T.root
	var x *RBTree //目标结点
   	for z != nil {
   		//查找结点
   		if z.val == key {
   			x = z
   			break
		}

   		if z.val > key {
   			z = z.lChild
   			continue
		}

   		z = z.rChild
	}

   	//判断关键字
   	if x == nil {
   		//不做任何处理
   		fmt.Println("结点不存在")
   		return
	}

   	//删除关键字


   }
/**
	交换结点
 */
func exChangeRBTree(T *RedBlackTree,x *RBTree,z *RBTree) {
	//修改父
	z.p = x.p

	//判断
	if x.p == nil {//表示根结点
		T.root = z
	} else if x.p.lChild == x {
		x.p.lChild = z
	} else {
		x.p.rChild = z
	}

}
/**
	删除结点

 */
func deleteNode(T *RedBlackTree,x *RBTree) {
	//记录要删除的结点颜色
	xColor := x.color

	var k *RBTree

	//判断孩子颜色情况
	if x.rChild == nil { //只有左孩子
		k = x.lChild
		//交换结点
		exChangeRBTree(T,x,x.lChild)

	} else if x.lChild == nil { //只有右孩子
		k = x.rChild
		//交换结点
		exChangeRBTree(T,x,x.rChild)

	} else {//左右孩子具有--从右子树寻找最小的当作z
		y := x.rChild

		var z *RBTree

		//判断
		if y.lChild == nil {//当前即为最小
			z = y
		} else {
			//从左子树开始寻找最小
			for y.lChild != nil {
				z = y
				y = y.lChild
			}
		}

		//记录颜色
		xColor = z.color
		k = z.rChild

		//修改孩子结点
		if k != nil {
			//表示向下寻找的
			if z.p != x {
				//调节
				exChangeRBTree(T,z,z.rChild)

				//x和z调节右孩子
				z.rChild = x.rChild
				z.rChild.p = z
			}

		}

		//交换位置
		exChangeRBTree(T,x,z)

		//调节左孩子
		z.lChild = x.lChild
		z.lChild.p = z

		z.color = x.color
	}

	//调节颜色
	if xColor == 1 {//当为黑色时调节
		RB_DELETE_FIXUP(T,k)

	}


}
/**
	x为x.p的左孩子
	1、x兄弟w结点是红色的---兄弟结点设为黑色，父设为红色，以双亲结点为根左旋转  w= x.p.right
	2、在1的基础上判断，w的左右孩子均为黑色 w.color = red x= x.p
	3在1的基础上 w的右孩子为黑色  w.left.color = black w.color = red w为中心右旋转 w= w.p.right
	4在3的基础上 w.color = x.p.color x.p.color = black w.right.color = black 左旋转以x.p x = T.root
    x为x.p的右孩子时反过来

	x.color = black
 */
func RB_DELETE_FIXUP(T *RedBlackTree, x *RBTree) {
	//当前不为root并且x 为黑色
	for x != T.root && x.color == 1 {
		if x == x.p.lChild {//当前为左子树
			w := x.p.rChild //寻找兄弟
			if w.color == 0 {//兄弟为红色
				w.color = 1
				x.p.color = 0
				lRotate(x.p,T)
				w = x.p.rChild
			}

			if w.lChild .color == 1 && w.rChild.color == 1 {
				w.color = 0
				x = x.p
			} else {
				if w.rChild.color == 1 {//左孩子为红色有孩子
					w.lChild.color = 1
					w.color = 0
					rRotate(w, T)
					w = x.p.rChild
				}
				w.color = x.p.color
				x.p.color = 1
				w.rChild.color = 1
				lRotate(x.p,T)
				x = T.root
			}
		} else {//反过来


		}
	}
	x.color = 1

}




func main() {
	var RBTree = new(RedBlackTree)
	RBTree = insertRBTree(RBTree, 1)
	RBTree = insertRBTree(RBTree, 2)
	RBTree = insertRBTree(RBTree, 0)
	RBTree = insertRBTree(RBTree, 3)
	RBTree = insertRBTree(RBTree, -1)
	RBTree = insertRBTree(RBTree, 5)
	RBTree = insertRBTree(RBTree, 6)
	RBTree = insertRBTree(RBTree, 7)
	RBTree = insertRBTree(RBTree, 8)
	RBTree = insertRBTree(RBTree, 9)
	RBTree = insertRBTree(RBTree, 10)
	RBTree = insertRBTree(RBTree, -2)
	RBTree = insertRBTree(RBTree, -4)
	RBTree = insertRBTree(RBTree, -6)
	RBTree = insertRBTree(RBTree, -9)
	RBTree = insertRBTree(RBTree, 11)
	blRBTree(RBTree.root)
}
