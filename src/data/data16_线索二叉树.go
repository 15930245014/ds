package main

import "fmt"

/**
	线索二叉树：
	n个结点的二叉链表中含有n+1(2n-(n-1)=n+1)个空指针域。
   利用二叉链表中的空指针域，
  存放指向结点在某种遍历次序下的前驱和后继结点的指针（这种附加的指针称为"线索"）。
	其中前驱和后继叫做线索
	原来二叉树标记前驱和后继的过程叫做线索化
 */

/**
   线索二叉树的结构体如下
*/
type BiThrTree struct {
	//结点结构
	data   byte         //值域
	lChild *BiThrTree //左孩子
	rChild *BiThrTree //右孩子
	lTag int   //左边的标记  1代表没有左孩子 可作前驱指向  1反之有
	rTag int   //右边的标记 1代表没有右孩子 可作后继指向   1反之有
}

var pre *BiThrTree

func main() {
	 arr := []byte{'A','B','E','M','D','#','#','C','#','#'}
	 index := 0
	 tree := makeBiThrTree(arr,&index)
	 //中序遍历
	 foreachMidBiThrTree(tree) //B->D->A->C

	 //线索化
	 midBiThrTree(tree)

	 maxLeft := getMaxLeft(tree)

	 blMidBiThrTree(maxLeft)

 }
/**
构造线索二叉树
 */
func makeBiThrTree(arr []byte,index *int) *BiThrTree{
	//判断是否越界
	if *index > len(arr) || arr[*index] == '#' {
		//返回nil
		return nil
	}

	//实例化线索二叉树
	biThrTree := new(BiThrTree)

	//存放值域
	biThrTree.data = arr[*index]
	biThrTree.lTag = 0
	biThrTree.rTag = 0

	//生成左孩子
	*index++
	if *index < len(arr) {
		biThrTree.lChild = makeBiThrTree(arr,index)
	}

	//生成右孩子
	*index++
	if *index < len(arr) {
		biThrTree.rChild = makeBiThrTree(arr,index)
	}

	//返回
	return biThrTree
}

/**
	中序遍历
 */

 func foreachMidBiThrTree(tree *BiThrTree) {
 	//判断根结点或叶子结点
 	if tree != nil {
 		//遍历左
 		foreachMidBiThrTree(tree.lChild)

 		//输出当下的值
 		fmt.Println(string(tree.data))

 		//遍历右
 		foreachMidBiThrTree(tree.rChild)

	}

 }


 /**
 	中序遍历构造线索二叉树----线索化
 	pre 前面的结点
  */
  func  midBiThrTree(tree *BiThrTree) {
	//判断当前结点
	if tree != nil {
		//遍历左
		midBiThrTree(tree.lChild)

		if pre != nil && tree.lChild == nil {
			tree.lChild = pre
			tree.lTag = 1
		}
		if pre != nil && pre.rChild == nil {
			pre.rChild = tree
			pre.rTag = 1
		}

		pre = tree
		//遍历右
		midBiThrTree(tree.rChild)

	}
  }

  /**
  中序遍历线索二叉树--后继
  tree为最左边的元素
   */
   func blMidBiThrTree(tree *BiThrTree) {

		p := tree

		for p != nil {
			fmt.Println(string(p.data))

			//判断标识
			p = p.rChild
		}
   }

   /*
   *
   	获取线索二叉树最左边--从根开始
    */
    func getMaxLeft(tree *BiThrTree) *BiThrTree {
    	if tree.lChild == nil {
    		return tree
 		} else {
 			return getMaxLeft(tree.lChild)
		}
	}






