package main

import "fmt"

/**
	二叉排序树（Binary sort tree）BST：又称为二叉搜索树或二叉查找树，它或者是一颗空树，或者是具有下列性质的二叉树：
	1、若它的左子树不空，则左子树上所有的结点的值均小于它的根节点；
	2、若它的右子树不为空，则右子树上所有的结点的值均大于它的根节点；
	3、它的左、右子树也分别为二叉排序树；
	4、二叉排序树不能存在重复值

 */

/**
	二叉排序树的数据结构:
 */
type BinarySortTree struct {
	 data           int
	 lChild, rChild *BinarySortTree
	 p *BinarySortTree  //双亲---出自算法导论
}

/*
* 二叉排序树的查找操作
*@author 许校
 */
func searchBST(T *BinarySortTree, key int) bool {
	//判断
	if T == nil {
		return false
	} else if T.data == key {
		return true
	} else if T.data > key {
		return searchBST(T.lChild, key)
	} else {
		return searchBST(T.rChild, key)
	}
}

/*
*二叉排序树非递归实现
*@author 许校
 */
func searchBTSXuXiao(T *BinarySortTree, key int) bool {
	//临时变量
	current := T

	//定义返回值
	rt := false

	//循环并判断
	for current != nil {
		//判断是否相等
		if current.data == key {
			//查找成功
			rt = true

			//终止循环
			break
		}

		//判断根结点的值
		if current.data > key {
			//寻找左树
			current = current.lChild
			continue
		}

		//寻找右树
		current = current.rChild
	}

	//返回值
	return rt
}

/**
	二叉排序树的添加操作：
	1、先寻找双亲结点位置
	2、根据双亲结点的值来判断添加左孩子还是右孩子
	3、如果存在则返回添加失败
	@param T *BinarySortTree 当前结点
	@param key int 要添加的值
	@param parent *BinarySortTree 双亲结点
	@author 许校
 */
func insertBTS(T *BinarySortTree, key int, parent *BinarySortTree) bool {
	//当前为空时---表示找到要添加的位置
	if T == nil {
		//构造树结点
		newNode := new(BinarySortTree)
		newNode.data = key

		//插入时添加双亲
		newNode.p = parent


		//判断
		if parent == nil { //表示为第一个结点
			T = newNode

			//返回true
			return true
		}
		//当前双亲值与key判断
		if parent.data > key { //当前为双亲的左孩子
			parent.lChild = newNode
			//返回true
			return true
		} else {
			parent.rChild = newNode
			//返回true
			return true
		}

	} else if T.data == key {//存在返回失败
		return false

	} else if T.data > key { //key 要在左孩子域添加
		//记录双亲结点
		parent = T
		return insertBTS(T.lChild, key, parent)
	} else { //key 要在右孩子域添加
		//记录双亲结点
		parent = T
		return insertBTS(T.rChild, key, parent)
	}

}

//寻找最小值
func findMinNode(T *BinarySortTree) int{
	//判断
	if T.lChild != nil {
		return findMinNode(T.lChild)
	} else {
		return T.data
	}
}

//寻找最大值
func findMaxNode(T *BinarySortTree) int {
	//判断
	if T.rChild != nil {
		return findMaxNode(T.rChild)
	} else {
		return T.data
	}
}

//二叉排序树的删除操作

/**
	如果要删除的结点只有左子树或右子树，将它的左子树或右子树整个移动到删除的结点的位置即可（子承父业）
	如果既有左子树又有右子树，找到需要删除的结点的直接前驱（或直接后继）中序遍历临近的前驱后继结点
 */
 func deleteBTS(T *BinarySortTree,key int) bool {
 	//寻找要删除的结点
 	if T != nil {
 		//判断
 		if T.data > key {//寻找左子树
 			return deleteBTS(T.lChild,key)
		} else if T.data < key { //寻找右子树
			return deleteBTS(T.rChild,key)
		} else {
			//执行删除
			return delBTS(T)
		}

	} else {//没有找到
		return false
	}

 }

 func delBTS(node *BinarySortTree) bool{

 	//判断当前为叶子结点
 	/*if node.rChild == nil && node.lChild == nil {
 		node = nil
 		return true;
	}*/

 	//判断是否左子树为nil
 	if node.lChild == nil { //右子树继承---叶子结点可以合并
 		//更改右子树双亲
 		node.rChild.p = node.p
 		*node = *(node.rChild) //地址没变值改变了

	} else if node.rChild == nil  {//左子树继承
		//更改左子树双亲
		node.lChild.p = node.p
		*node = *(node.lChild)
	} else {
		//寻找左直接前驱
		firstLeft := node.lChild

		//判断右
		if firstLeft.rChild == nil {//当前就是
			node.data = firstLeft.data
			node.lChild = firstLeft.lChild

			//更改双亲
			firstLeft.lChild.p = node

		}  else {//寻找右子树---最后的右子树的前驱
			s := firstLeft.rChild
			pre := firstLeft

			//寻找
			for s .rChild != nil {
				pre = s
				s = s.rChild
			}

			//修改
			node.data = s.data
			pre.rChild = s.lChild

			//修改双亲
			s.lChild.p = pre
		}

	}

 	//返回
 	return true

 }
//寻找前驱和后继
func getBSTsuccessor(p *BinarySortTree) int{
	//判断
	if p.rChild != nil {
		return findMinNode(p)
	}

	//判断当前的父
	x := p
	for x.p != nil && x == x.p.rChild {
		x = x.p
	}

	//返回
	return x.p.data
}




/**
	测试二叉排序树
 */
func main() {
	//声明跟结点
	var parent *BinarySortTree

	//添加二叉排序树
	arr := []int{8,5,9,4,7,6,3,2}

	//创建---需要创建根结点---败笔
	root := new(BinarySortTree)
	root.data  = arr[0]

	isAllow := true

	for i := 1; i < len(arr) && isAllow; i ++ {
		isAllow = insertBTS(root,arr[i], parent)
	}

	//删除
	deleteBTS(root,8)

	//中序遍历
	LDRBTS(root)

	//打印最小值
	fmt.Println("二叉搜索树的最小值：",findMinNode(root))

	//打印最大值
	fmt.Println("二叉排序树的最大值:",findMaxNode(root))

}

/**
	中序遍历二叉搜索树
 */
 func LDRBTS(T *BinarySortTree) {
 	if T != nil {
 		//遍历左
 		LDRBTS(T.lChild)

 		//输出当前
 		fmt.Println(T.data)

 		//遍历右
		LDRBTS(T.rChild)
	}
 }