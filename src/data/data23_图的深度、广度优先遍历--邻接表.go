package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	邻接矩阵的DFS
 */

/**
	构造顶点结构体
 */
type AdjacencyVertex struct {
	data      string //顶点信息
	firstEdge *AdjacencyEdge
}

/**
	构造边结构体
 */
type AdjacencyEdge struct {
	index  int //顶点索引
	weight int //权
	next   *AdjacencyEdge
}

/**
		声明邻接表图
 */
type AdjacencyChart struct {
	vertex             []AdjacencyVertex //顶点数组
	numVertex, numEdge int               //顶点数目和边的数目
}

/**
	邻接表的深度优先拓扑
 */

 func DFS_AdjacencyChart(chart AdjacencyChart) {
 	//声明访问切片
 	visited := []bool{}

 	//初始化
 	for i := 0 ; i < chart.numVertex; i ++ {
 		visited = append(visited,false)
	}

 	//遍历顶点
 	for j := 0; j < chart.numVertex; j ++ {
 		//判断是否访问过
 		if visited[j] == false {
 			DFSadjancyChart(chart,j,&visited)
		}

	}
 }
/**
	深度优先拓扑
 */
 func DFSadjancyChart(chart AdjacencyChart,index int, visited *[]bool) {
 	//打印当前顶点信息
 	fmt.Println(chart.vertex[index].data)

 	//当前节点标记为访问
	 (*visited)[index] = true

	 //获取当前的邻接点
	 p := chart.vertex[index].firstEdge

	 for p != nil {
	 	//遍历
	 	if !(*visited)[p.index] {//表示未访问过
			DFSadjancyChart(chart, p.index, visited)
		}

	 	//遍历下一个
	 	p = p.next
	 }

 }
/**
	邻接表的广度优先拓扑
 */

func BFS_AdjacencyChart(chart AdjacencyChart) {
	//声明访问切片
	visited := []bool{}

	//初始化
	for i := 0 ; i < chart.numVertex; i ++ {
		visited = append(visited,false)
	}

	//遍历顶点
	for j := 0; j < chart.numVertex; j ++ {
		//判断是否访问过
		if visited[j] == false {
			//打印当前顶点信息
			fmt.Println(chart.vertex[j].data)
			BFSadjancyChart(chart,j,&visited)
		}

	}
}
/**
	广度优先拓扑
 */
func BFSadjancyChart(chart AdjacencyChart,index int, visited *[]bool) {
	//当前节点标记为访问
	(*visited)[index] = true

	//获取当前的邻接点
	p := chart.vertex[index].firstEdge

	//定义当前的临界点数组
	var children []int

	for p != nil {
		//输出当前顶点信息
		fmt.Println(chart.vertex[p.index].data)

		//当前节点标记为访问
		(*visited)[p.index] = true

		//写入数组
		children = append(children,p.index)

		//遍历下一个
		p = p.next
	}

	//一次深度遍历子
	if len(children) > 0 {
		for k := 0; k < len(children); k++ {
			//遍历当前
			BFSadjancyChart(chart,children[k],visited)
		}

	}
}
 //测试
 func main() {
 	//构造邻接表
 	chart := makeAdjacencyChart()

 	fmt.Println("深度优先：")
 	//深度优先拓扑
 	DFS_AdjacencyChart(chart)

 	fmt.Println("广度优先:")
 	//广度优先拓扑
 	BFS_AdjacencyChart(chart)
 }

/**
	构造邻接表
 */
func makeAdjacencyChart() AdjacencyChart {
	//初始化图
	chartList := AdjacencyChart{}

	//读取控制台数据
	inputReader := bufio.NewReader(os.Stdin)

	//输入顶点的个数
	fmt.Print("请输入顶点的个数:")

	//读取数目
	inputVertexNum, err := inputReader.ReadString('\n')

	//判断是否读取失败
	if err != nil {
		fmt.Println("输入读取失败!")
	}

	//处理
	inputVertexNum = strings.Replace(inputVertexNum, "\n", "", -1)
	vertexNum, _ := strconv.Atoi(inputVertexNum)

	//图的顶点总数
	chartList.numVertex = vertexNum

	//输入顶点
	for i := 1; i <= vertexNum; i++ {
		fmt.Print("请输入第", i, "个顶点元素:")
		//读取顶点
		inputVertex, _ := inputReader.ReadString('\n')

		//去除\n
		inputVertex = strings.Replace(inputVertex, "\n", "", -1)

		//构建顶点
		Vertex := AdjacencyVertex{}

		Vertex.data = inputVertex
		Vertex.firstEdge = nil

		//存入图数组
		chartList.vertex = append(chartList.vertex, Vertex)
	}

	//输入边
	fmt.Print("请输入边的总数目:")

	//读取
	inputNumEdge, err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Print("读取错误请重新输入:")
	}
	inputNumEdge = strings.Replace(inputNumEdge, "\n", "", -1)

	//转换成整型
	NumEdge, _ := strconv.Atoi(inputNumEdge)

	//输入边
	for j := 1; j <= NumEdge; j++ {
		fmt.Print("请输入第", j, "条边(A,B,权值):")

		//读取
		inputEdge, _ := inputReader.ReadString('\n')
		inputEdge = strings.Replace(inputEdge, "\n", "", -1)

		EdgeArr := strings.Split(inputEdge, ",")

		//获取边或弧的起点索引
		indexB := getChartVertexIndex1(chartList, EdgeArr[0])

		//获取边或弧的终止索引
		indexE := getChartVertexIndex1(chartList, EdgeArr[1])

		//实例化边
		EdgeList := new(AdjacencyEdge)
		(*EdgeList).index = indexE
		(*EdgeList).weight, _ = strconv.Atoi(EdgeArr[2])

		//头插法
		EdgeList.next = chartList.vertex[indexB].firstEdge
		chartList.vertex[indexB].firstEdge = EdgeList
	}

	//返回图信息
	return chartList

}

/**
返回顶点在图的顶点数组的索引
*/
func getChartVertexIndex1(adjList AdjacencyChart, vertex string) int {
	//遍历
	rtIndex := -1
	for key, value := range adjList.vertex {
		if value.data == vertex {
			rtIndex = key
			break
		}

	}

	return rtIndex

}
