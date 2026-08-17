package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	DFS----边集数组
 */

/**
  定义顶点结构体
*/
type edgeSetVertex struct {
	data string
}

/**
	定义边结构体
 */
type edgeSetEdge struct {
	begin, end int //边的起止顶点在顶点数组中的索引
	weight     int //权
}

/**
	边集数组的图的结构体
 */
type edgeSetArray struct {
	vertex             []edgeSetVertex //顶点数组
	edge               []edgeSetEdge   //边集数组
	numVertex, numEdge int             //顶点，和边的数组
}

/**
	深度优先拓扑
 */
 func DFS_edgeSetArray(chart edgeSetArray) {
 	//声明访问数组标识
 	var visited []bool

 	//初始化
 	for i := 0;i < chart.numVertex; i++ {
 		visited = append(visited,false)
	}

 	//遍历
 	for j := 0; j < chart.numVertex;j ++ {
 		//判断当前顶点是否访问过
 		if visited[j] == false {
 			DFS_Edge(chart,j,&visited)
		}

	}

 }
 /**
 	DFS访问单个顶点
  */
 func DFS_Edge(chart edgeSetArray,index int,visited *[]bool) {
 	//输出信息
 	fmt.Println(chart.vertex[index].data)

 	//当前顶点标记为已访问
	 (*visited)[index] = true

	 //寻找当前节点的临界点
	 for k := 0; k < chart.numEdge; k++ {
	 	edge := chart.edge[k]

	 	//判断是否邻接
	 	if edge.begin == index && !(*visited)[edge.end]{//表示未访问过
	 		DFS_Edge(chart,edge.end,visited)
		}
	 }
 }
/**
	广度优先拓扑
 */
func BFS_edgeSetArray(chart edgeSetArray) {
	//声明访问数组标识
	var visited []bool

	//初始化
	for i := 0;i < chart.numVertex; i++ {
		visited = append(visited,false)
	}

	//遍历
	for j := 0; j < chart.numVertex;j ++ {
		//判断当前顶点是否访问过
		if visited[j] == false {
			//输出当前
			fmt.Println(chart.vertex[j].data)
			BFS_Edge(chart,j,&visited)
		}

	}

}
/**
	DFS访问单个顶点
 */
func BFS_Edge(chart edgeSetArray,index int,visited *[]bool) {

	//当前顶点标记为已访问
	(*visited)[index] = true

	//记录邻接下标
	var children []int

	//寻找当前节点的临界点
	for k := 0; k < chart.numEdge; k++ {
		edge := chart.edge[k]

		//判断是否邻接
		if edge.begin == index {
			//输出
			fmt.Println(chart.vertex[edge.end].data)

			//标记为访问
			(*visited)[edge.end] = true

			//记录
			children = append(children,edge.end)

		}
	}

	//临界点在一次深度
	if len(children) > 0 {
		for _,value := range children {
			//遍历
			BFS_Edge(chart,value,visited)
		}

	}
}


func main() {
	//构造图
	chart := makeEdgeSet()

	//DFS
	DFS_edgeSetArray(chart)

	//BFS
	BFS_edgeSetArray(chart)

}

/**
	构造边集数组
 */

func makeEdgeSet() edgeSetArray{
	//初始化图
	chart := edgeSetArray{}

	//实例化读取器
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入顶点数目:")

	//输入顶点个数
	inputVertexNum, err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Print("输入错误")
	}

	//去除\n
	inputVertexNum = strings.Replace(inputVertexNum, "\n", "", -1)

	//转化成int
	vertexNum, err := strconv.Atoi(inputVertexNum)
	chart.numVertex = vertexNum

	//输入顶点
	for i := 1; i <= chart.numVertex; i++ {
		fmt.Print("请输入第", i, "个顶点:")

		//读取顶点
		inputVertex, _ := inputReader.ReadString('\n')

		//构造顶点
		vertex := edgeSetVertex{}
		vertex.data = strings.Replace(inputVertex, "\n", "", -1)

		//保存
		chart.vertex = append(chart.vertex, vertex)

	}

	//输入边
	fmt.Print("请输入边的总数目:")

	//输入顶点个数
	inputEdgeNum, err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Print("输入错误")
	}

	//去除\n
	inputEdgeNum = strings.Replace(inputEdgeNum, "\n", "", -1)

	//转化成int
	edgeNum, err := strconv.Atoi(inputEdgeNum)
	chart.numEdge = edgeNum

	//输入边:
	for j := 1; j <= chart.numEdge; j++ {
		fmt.Print("请输入第",j,"条边（begin,end,weight）:")
		edge,_ := inputReader.ReadString('\n')
		edge = strings.Replace(edge,"\n","",-1)

		//分割成数组
		edgeArr := strings.Split(edge,",")

		//起始下标
		begin,_ := strconv.Atoi(edgeArr[0])

		//终止下标
		end,_ := strconv.Atoi(edgeArr[1])

		//权值
		weight,_ := strconv.Atoi(edgeArr[2])

		//构造结构体
		edgeStruct := edgeSetEdge{begin,end,weight}

		//保存图
		chart.edge = append(chart.edge,edgeStruct)
	}

	//返回
	return chart
}
