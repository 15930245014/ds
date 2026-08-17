package main

import "fmt"

/**
  二叉树：是n个结点的有限结合，该集合或者为空集(空二叉树)，或者由一个根节点和2棵互不相交的左子树和右子树的二叉树组成
*/
/**
   二叉树类别：斜树（根节点只有一个子 l 左斜 r 右斜）、满二叉树（深度n-1以上中间结点全部有2个孩子），完全二叉树（排列和完全一样 但不见得是满）
*/

/**
	二叉树存储结构---顺序存储（一般只适用于完全二叉树 0 -1 -2 根的子下标 堆排序）、二叉链表
 */

/**
	二叉链表：二叉树每个结点最多有2个孩子，所以设计一个数据域和2个指针域 一个指向左孩子一个指向右孩子
 */

/**
	二叉链表的结构体如下
 */
type binaryNode struct {
	//结点结构
	data   byte         //值域
	lChild *binaryNode //左孩子
	rChild *binaryNode //右孩子
}

//下面是正常但是可以省略
type binaryTree struct {
	root *binaryNode //跟结点
	n    int         //结点个数
} //正常是需要建立一个binaryTree 但可以参考单链表 故rootNode可以作为树的标识 故binaryNode 就可以代表一个二叉树

func main() {
	arr := []byte{'A','B','#','D','#','#','C','#','#'}
	index := 0
	tree := createNewBinaryTree(arr,&index)

	//先序遍历输出
	foreachBTFront(tree) //A->B->D->C
	fmt.Println("---------------------------")
	//中序遍历
	foreachBTmedian(tree) //B->D->A->C

	fmt.Println("---------------------------")

	//后序遍历
	foreachBTlast(tree) //D->B->C->A

	fmt.Println("---------------------------")

	//层序遍历  //A->B->C->D
	start := []*binaryNode{tree}
	foreachBTheight(start)
}

/**
构建二叉树--利用切片创建（完全二叉树）----书中错误原因指针在函数体内重新赋值了已经不能够传引用了
 */
func createBinaryTree(tree *binaryNode, arr []byte, index *int)  {

	if *index < len(arr) { //-1表示结束的标识

		if arr[*index] == '#' {
			tree = nil

		} else {
			tree = new(binaryNode)
			(*tree).data = arr[*index]

			//左孩子
			*index ++
			createBinaryTree((*tree).lChild, arr, index)

			//右孩子
			*index ++
			createBinaryTree((*tree).rChild, arr, index)

		}
	}

}

/**
	许校实现
 */
func createNewBinaryTree(arr []byte,index *int) *binaryNode {
	//判断是否结束
	if *index> len(arr) || arr[*index] == '#' {
		return nil
	}
	//构造当前的树结点
	bNode := new(binaryNode)

	//存放值域
	(*bNode).data = arr[*index]

	//左孩子
	*index ++
	bNode.lChild = createNewBinaryTree(arr,index)

	//右孩子
	*index++
	bNode.rChild = createNewBinaryTree(arr,index)

	//返回二叉树
	return bNode
}

/**
	前序遍历---很好理解
 */
func foreachBTFront(bTree *binaryNode) {
	if bTree != nil {
		//输出当前的结点
		fmt.Println(string(bTree.data))

		//遍历左
		foreachBTFront(bTree.lChild)

		//遍历右
		foreachBTFront(bTree.rChild)
	}

}

/**
	中序遍历---很好理解
 */
func foreachBTmedian(bTree *binaryNode) {
	if bTree != nil {
		//遍历左
		foreachBTmedian(bTree.lChild)

		//输出当前的结点
		fmt.Println(string(bTree.data))

		//遍历右
		foreachBTmedian(bTree.rChild)
	}

}

/**
	后序遍历---很好理解
 */
func foreachBTlast(bTree *binaryNode) {
	if bTree != nil {
		//遍历左
		foreachBTlast(bTree.lChild)

		//遍历右
		foreachBTlast(bTree.rChild)

		//输出当前的结点
		fmt.Println(string(bTree.data))
	}

}

/**
层序遍历---类似广度优先---许校扩展
 */

 func foreachBTheight(trees [] *binaryNode) {
 	if len(trees) > 0 {
 		//输出当前构造下一次的层
 		next := []*binaryNode{}

 		for _,tree := range trees {
 			fmt.Println(string(tree.data))

 			//判断是否有左孩子
 			if tree.lChild != nil {
 				next = append(next,tree.lChild)
			}

 			//判断是否存在右孩子
 			if tree.rChild != nil {
 				next = append(next,tree.rChild)
			}

		}

 		//判断是否存在下一层
 		if len(next) > 0 {
 			foreachBTheight(next)
		}

	}


 }
