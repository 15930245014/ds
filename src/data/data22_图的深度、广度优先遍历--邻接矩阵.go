package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	从图中某一顶点出发遍历图中其余的顶点且使每个顶点仅被访问一次这个过程叫做图的遍历
	深度优先遍历（DFS）

 */

//定义图的最大顶点数目
const A = 40

//构造邻接矩阵
type AdjacencyMatrixs struct {
	vertex [A]string
	src [A][A]int
	numVertex,numEdges int //顶点数目和边的数目
}
/**
	邻接矩阵的深度优先拓扑
 */

 func DFS_AdjacencyMatrix(chart AdjacencyMatrixs) {
 	//声明已访问顶点数组标记
 	visitedArr := []bool{}

 	//初始化标记数组为false
 	for k := 0 ; k < chart.numVertex; k ++ {
 		visitedArr = append(visitedArr,false)
	}

 	//遍历
 	for i := 0; i < chart.numVertex; i ++ {
 		//判断顶点是否访问过
 		if visitedArr[i] == false {
 			DFS_AdvanceVertex(chart,i,&visitedArr)
		}

	}

 }
 /**
 	1图，要深度优先的顶点，访问数组
  */
 func DFS_AdvanceVertex(chart AdjacencyMatrixs,index int,visitedArr *[]bool) {
 	//输出当前顶点
 	fmt.Println(chart.vertex[index])

 	//当前顶点标记为访问
	 (*visitedArr)[index] = true

 	//以当前顶点寻找其临界点src[当前节点下标][j] 且值不为0
 	for j := 0; j < len(chart.src[index]);j ++ {
 		//判断值
 		if chart.src[index][j] != 0 {
 			//当前节点为顶点继续
 			DFS_AdvanceVertex(chart,j,visitedArr)
		}

	 }
 }
 /**
 	邻接矩阵的广度优先拓扑
  */

func BFS_AdjacencyMatrix(chart AdjacencyMatrixs) {
	//声明已访问顶点数组标记
	visitedArr := []bool{}

	//初始化标记数组为false
	for k := 0 ; k < chart.numVertex; k ++ {
		visitedArr = append(visitedArr,false)
	}

	//遍历
	for i := 0; i < chart.numVertex; i ++ {
		//判断顶点是否访问过
		if visitedArr[i] == false {
			//输出当前顶点
			fmt.Println(chart.vertex[i])
			BFS_AdvanceVertex(chart,i,&visitedArr)
		}

	}

}
/**
	1图，要深度优先的顶点，访问数组
 */
func BFS_AdvanceVertex(chart AdjacencyMatrixs,index int,visitedArr *[]bool) {
	//当前顶点标记为访问
	(*visitedArr)[index] = true

	//获取当前的邻接全部输出并记录下标
	var children []int

	for j := 0 ; j < chart.numVertex;j ++ {
		//邻接并且判断是否访问过
		if chart.src[index][j] != 0 && (*visitedArr)[j] != false {
			fmt.Println(chart.vertex[j])
			children = append(children,j)
			(*visitedArr)[j] = true
		}
	}

	//递归
	for k := 0; k < len(children);k++ {
		BFS_AdvanceVertex(chart,children[k],visitedArr)
	}

}


/**
	构造邻接矩阵
 */
func makeAdjacencyMatrix() AdjacencyMatrixs {
	//创建读取器
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Printf("请输入图顶点的数目: ")

	inputNumVertex, err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Println("输入错误!")
	}
	//去掉\n
	inputNumVertex = strings.Replace(inputNumVertex, "\n", "", -1)

	//获取图顶点数目
	numVertex,_ := strconv.Atoi(inputNumVertex)


	//初始化图
	adjacencyMatix := AdjacencyMatrixs{}

	//初始化邻接矩阵
	adjacencyMatix.numVertex = numVertex //顶点数目

	//初始化边
	for i:= 0; i< numVertex; i ++ {
		for j := 0; j < numVertex; j++ {
			adjacencyMatix.src[i][j] = 0
		}

	}

	//输入顶点
	for i := 1; i <= numVertex; i ++ {
		fmt.Printf("请输入第%d个顶点:",i)
		vertex, err := inputReader.ReadString('\n')
		if err != nil {
			i--
			fmt.Println("输入错误请重新输入")
			continue
		}

		//获取
		adjacencyMatix.vertex[i - 1] = strings.Replace(vertex,"\n","",-1)
	}

	//输入边
	fmt.Printf("请输入图边或弧的数目: ")

	inputNumEdgs, err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Println("输入错误!")
	}

	inputNumEdgs = strings.Replace(inputNumEdgs,"\n","",-1)

	//转化成整形
	numEdgs,_ := strconv.Atoi(inputNumEdgs)

	adjacencyMatix.numEdges = numEdgs

	//输入边
	for i := 1; i <= numEdgs; i ++ {
		fmt.Printf("请输入第%d个边或弧 i,j,值形式:",i)
		edge, err := inputReader.ReadString('\n')
		if err != nil {
			i--
			fmt.Println("输入错误请重新输入")
			continue
		}

		//获取
		edge = strings.Replace(edge,"\n","",-1)

		//分割成数组
		edgeArr := strings.Split(edge,",")

		//转化为整形
		indexI,_ := strconv.Atoi(edgeArr[0])
		indexj,_ := strconv.Atoi(edgeArr[1])
		indexV,_ := strconv.Atoi(edgeArr[2])

		//保存
		adjacencyMatix.src[indexI][indexj] = indexV

	}

	//返回邻接矩阵
	return adjacencyMatix

}

//测试
func main() {
	//创建图
	chart := makeAdjacencyMatrix()

	//深度优先拓扑
	DFS_AdjacencyMatrix(chart)

	//广度优先拓扑
	BFS_AdjacencyMatrix(chart)
}
