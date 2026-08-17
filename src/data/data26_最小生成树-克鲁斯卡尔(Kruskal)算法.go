package main

import (
	"fmt"
	"strconv"
)

/**
	Kruskal算法：一般采用边集集合解决
	基本思想：
      按照权值从小到大的顺序选择n - 1条边，并保证这n - 1条边不构成回路

	算法要点:
	 1、对图的所有边按照权值大小排序。
     2、将边添加到最小生成树中，避免生成回路。
		回路判断方法:
          通过记录每个顶点在最小生成树中的终点。终点即在最小生成树中与它连通的最大顶点。每次添加一条边到最小生成树中时，判断该边的两个顶点的终点是否重合，重合则构成回路。
 */

/**
	定义边集集合--边的结构体
 */

type edgeNode struct {
	begin, end int //边的起始顶点索引
	weight     int //权
}

/**
	边集集合
 */
type edgeNodeSet struct {
	vertex             []string   //图的顶点
	edge               []edgeNode //边的集合
	numVertex, numEdge int        //顶点和边的数目
}

/**
		Kruskal 算法实现思路：
		1、对图的边集按照升序的顺序排序
	 2、初始化一个parent[begin] = end 的数组 用来寻找最大终点,treeEdge数组用来保存边---此时即为最小生成树的边
	 3、循环边 且当 treeEdge的长度 为（n个顶点的最小生成树有n-1条边）n-1时结束

 */

func Kruskal(edgeChart edgeNodeSet) []edgeNode {
	//获取图的边集集合
	edgeArr := edgeChart.edge

	//对边按照权进行升序排序---插入排序
	for i := 1; i < len(edgeArr); i ++ {
		//记录当前
		tempNode := edgeArr[i]

		//前一位
		j := i - 1

		//判断
		for j >= 0 && edgeArr[j].weight > tempNode.weight {
			edgeArr[j+1] = edgeArr[j]
			j --
		}

		//当前位置写入
		edgeArr[j+1] = tempNode
	}

	//初始化---前驱数组
	var parent []int
	for i := 0; i < edgeChart.numVertex; i++ {
		parent = append(parent, 0)
	}

	//定义结果集
	var result []edgeNode

	//循环条件
	for i := 0; i < len(edgeArr) && len(result) < edgeChart.numVertex - 1; i ++ {
		//获取边集集合的起始
		begin := edgeArr[i].begin
		end := edgeArr[i].end

		//获取终点
		beginTop := findVertexTop(parent, begin)
		endTop := findVertexTop(parent, end)

		//如果不重合
		if beginTop != endTop {
			//当前边写入
			result = append(result, edgeArr[i])

			//写入parent
			parent[beginTop] = endTop

		}

	}

	//返回
	return result

}

/**
	获取终点
	@param parent 父数组
	@param begin 起点下标
	   parent数组不存在时 当前为终点
 */
func findVertexTop(parent []int, begin int) int {
	//判断
	for parent[begin] > 0 {
		begin = parent[begin]
	}

	//返回
	return begin
}
  /**
  测试最小生成树
   */
func main() {
	//初始化图
	var chart = edgeNodeSet{}

	//输入顶点
	chart.numVertex = 9
	for i := 0; i < chart.numVertex; i++ {
		chart.vertex = append(chart.vertex,"v" + strconv.Itoa(i))
	}

	//输入边
	chart.numEdge = 15
	chart.edge = append(chart.edge,edgeNode{4,5,26})
	chart.edge = append(chart.edge,edgeNode{3,6,24})
	chart.edge = append(chart.edge,edgeNode{2,3,22})
	chart.edge = append(chart.edge,edgeNode{3,8,21})
	chart.edge = append(chart.edge,edgeNode{3,4,20})
	chart.edge = append(chart.edge,edgeNode{6,7,19})
	chart.edge = append(chart.edge,edgeNode{1,2,18})
	chart.edge = append(chart.edge,edgeNode{5,6,17})
	chart.edge = append(chart.edge,edgeNode{1,6,16})
	chart.edge = append(chart.edge,edgeNode{3,7,16})
	chart.edge = append(chart.edge,edgeNode{1,8,12})
	chart.edge = append(chart.edge,edgeNode{0,5,11})
	chart.edge = append(chart.edge,edgeNode{0,1,10})
	chart.edge = append(chart.edge,edgeNode{2,8,8})
	chart.edge = append(chart.edge,edgeNode{4,7,7})

	//打印
	fmt.Println(Kruskal(chart))

}