package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
	最小生成树：n个顶点用n-1条弧连接起来使所有弧的权值之和最小，
    这种构造连通网的最小代价生成树叫做最小生成树。
	构造最小生成树构造算法：普里姆（Prim）算法和克鲁斯卡尔（Kruskal）
	Prim算法从顶点角度考虑适合稠密图 Kruskal算法从边角度来考虑 适合稀疏图
 */
/**
	Prim算法描述(一般采用邻接矩阵):
	声明lowCost 最小化花费，初始化为任意顶点的到其它顶点的距离
    1、从任意顶点开始，寻找该顶点到其他顶点的最小顶点，并将lowCost标记为0（表示该顶点已经纳入最小生成树），若为0代表该顶点已经为最小生成树顶点
		若顶点不连通用正无穷表示
	2、已上一步找到的顶点为基准在循环遍历与之连通的未纳入最小生成树的顶点，若与其他顶点到某一顶点距离与lowCost作比较，lowCost取最小
    3、重复1，2步骤 直到lowCost全部为0
 */

/**
声明邻接矩阵
 */
type chartMatrix struct {
	vertex    []string
	src       [][]int
	numVertex, numEdge int
}

/**
	定义树---采用双亲表示法
 */
/**
	孩子结点结构体
 */
type treeNode struct {
	data string //每个结点的值域
	parent int //双亲的索引
}
/**
	定义树的结构体
 */
type parentTree struct {
	nodes  []treeNode //结点的集合---树无非就是各个结点的集合
	r,n int //root结点的下标以及树中结点的数目
}

func main() {
	//构造邻接矩阵 ---构造过程简写
	maxInt := math.MaxInt64 //正无穷表示2个顶点不相关联
	chart := chartMatrix{}

	//构造顶点
	chart.numVertex = 9
	for i := 0; i < chart.numVertex; i++ {
		chart.vertex = append(chart.vertex,"v" + strconv.Itoa(i))
	}

	//构造邻接矩阵
	chart.src = append(chart.src,[]int{0,10,maxInt,maxInt,maxInt,11,maxInt,maxInt,maxInt})
	chart.src = append(chart.src,[]int{10,0,18,maxInt,maxInt,maxInt,16,maxInt,12})
	chart.src = append(chart.src,[]int{maxInt,maxInt,0,22,maxInt,maxInt,maxInt,maxInt,8})
	chart.src = append(chart.src,[]int{maxInt,maxInt,22,0,20,maxInt,maxInt,16,21})
	chart.src = append(chart.src,[]int{maxInt,maxInt,maxInt,20,0,26,maxInt,7,maxInt})
	chart.src = append(chart.src,[]int{11,maxInt,maxInt,maxInt,26,0,17,maxInt,maxInt})
	chart.src = append(chart.src,[]int{maxInt,16,maxInt,maxInt,maxInt,17,0,19,maxInt})
	chart.src = append(chart.src,[]int{maxInt,maxInt,maxInt,16,7,maxInt,19,0,maxInt})
	chart.src = append(chart.src,[]int{maxInt,12,8,21,maxInt,maxInt,maxInt,maxInt,0})

	tree := Prim(chart)

	//层序遍历树
	DFSTree(-1,tree,)

}
/**
层序遍历tree
 */
func DFSTree(parent int,tree parentTree) {
	//寻找根
	var children []int
	for i := 0; i < len(tree.nodes); i++ {
		//判断是否fu
		if tree.nodes[i].parent == parent {
			children = append(children,i)
		}

	}
	//一层一层的-----采用队列的形式
	for m := 0 ;m < len(children);m++ {
		//输出
		fmt.Println(tree.nodes[children[m]].data)

		//寻找子
		for i := 0; i < len(tree.nodes); i++ {
			//判断是否fu
			if tree.nodes[i].parent == children[m] {
				children = append(children,i)
			}
		}
	}



}

/**
	普里姆算法
 */
 func Prim(chart chartMatrix) parentTree {
	//初始化一棵树
	tree := parentTree{}

	//初始化parent和lowCost
	var parent,lowCost []int

	//初始化parent,第一个结点
	for k := 0;k < chart.numVertex;k++  {
		parent = append(parent, 0)
	}
	//第1个顶点的与之的权
	lowCost = chart.src[0] //此时lowCost[0] = 0 代表vo已经纳入生成树
	lowCost[0] = 0

	//生成最小生成树的根节点
	root := treeNode{}
	root.parent = -1
	root.data = chart.vertex[0]
	tree.nodes = append(tree.nodes,root)

	//循环条件--lowCost全部为0 每一次都找到一个最优顶点
	for i := 1; i < len(lowCost);i++ {
		//lowCost中寻找权值最小的
		min := math.MaxInt64
		nodeIndex := -1 //权值最小的索引，-1表示此图不为连通图图错误
		for j := 0; j < len(lowCost);j++ {
			//若果lowCost 为0表示已经纳入最小生成树了
			if lowCost[j] != 0 && lowCost[j] < min {
				min = lowCost[j]
				nodeIndex = j
			}

		}

		//寻找当前的父
		parentIndex := parent[nodeIndex]

		//构造当前的叶子节点
		tempTree := treeNode{}
		tempTree.parent = parentIndex
		tempTree.data = chart.vertex[nodeIndex]

		//存入tree
		tree.nodes = append(tree.nodes,tempTree)

		//并返回当前叶子结点的id ---最后一个
		parentID := len(tree.nodes) - 1

		//lowCost 更新为0---将当前的顶点纳入最小生成树
		lowCost[nodeIndex] = 0

		//已当前结点来寻找最小花费
		for k := 0; k < len(chart.src[nodeIndex]); k++ {
			//判断
			if lowCost[k] != 0 && lowCost[k] > chart.src[nodeIndex][k] {
				//更新lowCost
				lowCost[k] = chart.src[nodeIndex][k]

				//更新当前的父
				parent[k] = parentID

			}

		}

	}

	//返回tree
	return tree

 }
