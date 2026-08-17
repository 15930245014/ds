package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/**
	图的存储方式---邻接表  -----解决邻接矩阵边数组和顶点较少时的空间浪费
	----一维数组来存放顶点，单链表来存储边，这种数组和链表相结合的存储方式称为邻接表（Adjacency List）
 */

/**
	顶点结构分为 data域 和 firstEdge域 firstEdge 指向出边表的第一个结点
 */

/**
	单链表结构（边表结点） adjVex 顶点的索引下标(0,1,2...) next指针域（指向下一个的边表结点）
	若存在权 则 添加weight域
 */

/**
	构造顶点的结构体
 */
type chartVertex struct {
	data      string     //顶点标识 "A","B"
	firstEdge *chartEdge //第一个出表
}

/**
	构造边表结构体
 */
type chartEdge struct {
	weight int    //权
	index  int        //顶点在图顶点集合的索引
	next   *chartEdge //下一个指针域
}

/**
	构造图结构体
 */
type AdjacencyList struct {
	vertex             [20]chartVertex //顶点数组
	numVertex, numEdge int             //顶点和边的数组
}
/**
	构造邻接表
 */
func main () {
	//初始化图
	chartList := AdjacencyList{}

	//读取控制台数据
	inputReader := bufio.NewReader(os.Stdin)

	//输入顶点的个数
	fmt.Print("请输入顶点的个数:")

	//读取数目
	inputVertexNum,err := inputReader.ReadString('\n')

	//判断是否读取失败
	if err != nil {
		fmt.Println("输入读取失败!")
	}

	//处理
	inputVertexNum = strings.Replace(inputVertexNum,"\n","",-1)
	vertexNum,_ := strconv.Atoi(inputVertexNum)

	//图的顶点总数
	chartList.numVertex = vertexNum

	//输入顶点
	for i:= 1; i<= vertexNum; i++ {
		fmt.Print("请输入第",i,"个顶点元素:")
		//读取顶点
		inputVertex,_ := inputReader.ReadString('\n')

		//去除\n
		inputVertex = strings.Replace(inputVertex,"\n","",-1)

		//构建顶点
		Vertex := chartVertex{}

		Vertex.data = inputVertex
		Vertex.firstEdge = nil

		//存入图数组
		chartList.vertex[i - 1] = Vertex
	}

	//输入边
	fmt.Print("请输入边的总数目:")

	//读取
	inputNumEdge,err:= inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Print("读取错误请重新输入:")
	}
	inputNumEdge = strings.Replace(inputNumEdge,"\n","",-1)

	//转换成整型
	NumEdge,_ := strconv.Atoi(inputNumEdge)

	//输入边
	for j := 1; j <= NumEdge; j++ {
		fmt.Print("请输入第",j,"条边(A,B,权值):")

		//读取
		inputEdge,_ := inputReader.ReadString('\n')
		inputEdge = strings.Replace(inputEdge,"\n","",-1)

		EdgeArr := strings.Split(inputEdge,",")

		//获取边或弧的起点索引
		indexB := getChartVertexIndex(chartList,EdgeArr[0])

		//获取边或弧的终止索引
		indexE := getChartVertexIndex(chartList,EdgeArr[1])

		//实例化边
		 EdgeList := new(chartEdge)
		(*EdgeList).index = indexE
		(*EdgeList).weight,_ = strconv.Atoi(EdgeArr[2])

		//头插法
		EdgeList.next = chartList.vertex[indexB].firstEdge
		chartList.vertex[indexB].firstEdge = EdgeList
	}

	//构造完成测试
	for i := 0; i <chartList.numVertex; i++ {
		vertex := chartList.vertex[i]
		fmt.Println("当前顶点为",vertex.data)
		fmt.Println("出度结点为:")

		//获取出度
		p := vertex.firstEdge

		for p != nil {
			//获取出度索引
			index := p.index
			fmt.Print(chartList.vertex[index].data,"权为：",p.weight,";")
			p = p.next
		}
	}

}
/**
	返回顶点在图的顶点数组的索引
 */
func getChartVertexIndex(adjList AdjacencyList, vertex string) int {
	//遍历
	rtIndex := -1
	for key,value := range adjList.vertex {
		if value.data == vertex {
			rtIndex = key
			break
		}

	}

	return rtIndex

}

