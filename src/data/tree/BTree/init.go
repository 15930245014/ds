package BTree

/**
	B树：也称B-树，它是一颗多路平衡查找树。
	阶：阶数表示了一个结点最多有多少个孩子结点，假设阶为M一般把一棵B树称作M阶BTree.
   一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树。
	一颗m阶的B树定义如下：
	1、每个结点最多有m-1个关键字。
	2、根结点最少可以只有1个关键字。
	3、非根结点至少有Math.ceil(m/2)-1个关键字。
	4、每个结点中的关键字都按照从小到大的顺序排列，每个关键字的左子树中的所有关键字都小于它，而右子树中的所有关键字都大于它。
    5、所有叶子结点都位于同一层，或者说根结点到每个叶子结点的长度都相同。
原文链接：https://blog.csdn.net/qq_36183935/article/details/80382490
 */

const M = 6     //M阶
const HALF = 2  // ceil(m / 2) - 1

//空值定义
var  NULL = 0
var Nil *BTNode = nil

/**
	定义B树数据结构
 */
type BTNode struct {
	P *BTNode
	KeyNum int
	Keys [] int //值域从1开始 M+1
	Children []*BTNode //索引从0开始 到 M+1
}
type BTree struct {
	Root *BTNode
	Sum int
}

/**
	初始化BTree
 */
func InitBTree() *BTree {
	//初始化变量
	bTree := new(BTree)

	//设置root
	bTree.Root = initBNode()
	return bTree
}

/**
	初始化BNode
 */
func initBNode() *BTNode {
	bNode := new(BTNode)

	//初始化bNode
	bNode.P = nil
	bNode.KeyNum = 0
	bNode.Keys = make([]int, M + 1)
	bNode.Children = make([]*BTNode, M + 1)
	for i := 0; i < M + 1; i ++ {
		bNode.Keys[i] = NULL
		bNode.Children[i] = Nil
	}
	return bNode
}






