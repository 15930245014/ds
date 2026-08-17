package B_Tree
/**
	B+树和B树区别：
    	非叶子节点只存储键值信息。
		所有叶子节点之间都有一个链指针。
		数据记录都存放在叶子节点中。
 */

/**
	定义B+树节点结构
 */
type BAddTreeNode struct {
	 keyNum int  			 //节点数目
	 leaf bool   			 //是否是叶子节点
	 p *BAddTreeNode 	 	 //父节点
	 keys []int  			 //索引 从1开始
	 children []*BAddTreeNode //子节点
	 data []int	 		 	 //值域数组 可以用链表表示
	 pre,next *BAddTreeNode  //前驱后继块
}

/**
	定义b+树结构
 */
type BAddTree struct {
	Root *BAddTreeNode //根节点
	Num int			   //数目
	Min,Max int		   //索引最大、最小。。。。
}


const M = 5 //M阶
const HALF =  3 // ceil(m / 2) - 1

//空值定义
var  NULL = 0
var Nil *BAddTreeNode = nil

/**
	初始化B+树
 */
func initBAddTree() *BAddTree {
	baTree := new(BAddTree)
	//默认当前就是叶子结点
	baTree.Root = initBAddTreeNode(true)

	baTree.Min = 0
	baTree.Max = 0
	baTree.Num = 0
	return baTree
}

/**
	初始化B+树节点
 */
func initBAddTreeNode(isLeaf bool) *BAddTreeNode {
	//初始化节点
	node := new(BAddTreeNode)
	node.keyNum = 0
	node.leaf = isLeaf
	node.p = Nil
	node.pre = Nil
	node.next = Nil

	//初始化索引
	node.keys = make([]int,M + 1)

	//初始化children
	node.children = make([]*BAddTreeNode,M + 1)

	//初始化data ---叶子节点需要
	if isLeaf {
		node.data = make([]int, M+1)
	}

	//返回
	return node
}








