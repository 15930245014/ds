package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	拓扑排序:
	AOV网:在一个表示工程的有向图中，用顶点表示活动，用弧表示活动之间的优先关系，这样的有向图为顶点表示活动的网，叫做AOV网。网中的弧表示活动之间存在的某种制约关系。

	拓扑序列:
	设 G = (V,E)是一个具有n个顶点的有向图，V中的顶点序列v1,v2,....,vn,满足若从顶点vi到vj有一条路径，则在顶点序列中顶点vi必在顶点vj之前。则我们称这样的顶点序列为一个拓扑序列

	拓扑排序：
	 拓扑排序，其实就是对一个有向图构造拓扑序列的过程，构造时会有2个结果，如果此网的全部顶点被输出，则它是不存在环（回路）的AOV网；如果输出顶点数少了，哪怕是少了一个也说明这个网存在环（回路）不是AOV网

	拓扑排序算法：
      对AOV网进行拓扑排序的基本思路是:从AOV网中选择一个入度为0的顶点输出，然后删去此顶点，并删除以此顶点为尾的弧，继续重复此步骤，直到输出全部顶点或者AOV网中不存在入度为0的顶点为止

 */

 /**
 	采用邻接表的表现形式---顶点多加一项入度数目
  */
  /**
  	邻接表边结构体
   */
   type edgeVertex struct {
   		index int //顶点索引
   		weight int //权
   		next *edgeVertex //下一个指针域和
   }

   /**
   	邻接表顶点结构体
    */
    type vertexBox struct {
    	in int //入度之和
    	data string //顶点标识
    	firstEdge *edgeVertex //指向的第一个边
	}

    /**
    	邻接表
     */
     type ljChart struct {
     	vertexBox []vertexBox
     	numVertex,numEdge int //边，顶点数目
	 }

/**
	拓扑排序
 */
func TopologicalSort (chart ljChart)[]string {
	//初始化
	var stack []int //定义栈保存入度为0的顶点的索引

	//获取入度为0的顶点
	for i := 0 ; i < chart.numVertex; i ++ {
		//判断入度数目
		if chart.vertexBox[i].in == 0 {
			//索引入栈
			stack = append(stack,i)
		}
	}

	//声明访问
	visited := make([]bool,chart.numVertex)

	//记录顶点
	var result []string

	//当栈不为空时
	for len(stack) > 0 {
		//标记当前访问
		visited[stack[len(stack) - 1]] = true

		//栈顶元素出栈
		node := chart.vertexBox[stack[len(stack) - 1]]
		stack = stack[:len(stack) - 1]

		//遍历
		e := node.firstEdge

		//写入结果
		result = append(result,node.data)

		for e != nil {
			//获取e的指向顶点
			index := e.index

			//指向顶点
			target := chart.vertexBox[index]
			target.in --

			//判断
			if target.in == 0 && visited[index] == false {//入栈
				stack = append(stack,index)
			}

			//指针下移
			e = e.next

		}

	}

	//返回
	return result
}

func getLjChart() ljChart {
	//初始化图
	chartList := ljChart{}

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
		fmt.Print("请输入第",i,"个顶点元素:名称,入度数目")
		//读取顶点
		inputVertex,_ := inputReader.ReadString('\n')

		//去除\n
		inputVertex = strings.Replace(inputVertex,"\n","",-1)

		inputVertexArr := strings.Split(inputVertex,",")

		//构建顶点
		Vertex := vertexBox{}

		Vertex.data = inputVertexArr[0]
		Vertex.firstEdge = nil

		in,_ := strconv.Atoi(inputVertexArr[1])
		Vertex.in = in

		//存入图数组
		chartList.vertexBox = append(chartList.vertexBox,Vertex)
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
		indexB := getChartVertexIndex2(chartList,EdgeArr[0])

		//获取边或弧的终止索引
		indexE := getChartVertexIndex2(chartList,EdgeArr[1])

		//实例化边
		EdgeList := new(edgeVertex)
		(*EdgeList).index = indexE
		(*EdgeList).weight,_ = strconv.Atoi(EdgeArr[2])

		//头插法
		EdgeList.next = chartList.vertexBox[indexB].firstEdge
		chartList.vertexBox[indexB].firstEdge = EdgeList
	}

	//返回
	return chartList

}

/**
	返回顶点在图的顶点数组的索引
 */
func getChartVertexIndex2(adjList ljChart, vertex string) int {
	//遍历
	rtIndex := -1
	for key,value := range adjList.vertexBox {
		if value.data == vertex {
			rtIndex = key
			break
		}

	}

	return rtIndex

}

func main() {//测试
	//生成邻接表
	ljChart := getLjChart()

	//获取拓扑排序
	result := TopologicalSort(ljChart)

	//输出
	fmt.Println(result)

}





