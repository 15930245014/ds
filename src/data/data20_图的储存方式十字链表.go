package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
	邻接表的储存方式 在处理有向图的时候 需要建立逆邻接表故处理双向有向图的时候存在不足
	十字链表--主要处理双向有向图的问题 将逆邻接表和正邻接表结合在一起的存储方式
 */

/**
	十字链表（Orthogonal List）--顶点结构体
	A、data 值域用于保存顶点信息
	B、firstOut 第一个指向出表的结点指针域
	C、firstIn 第一个指向入表的结点指针域
 */
type OrthogonalVertex struct {
	data     string
	firstOut *OrthogonalEdge
	firstIn  *OrthogonalEdge
}

/**
	十字链表（Orthogonal List）--弧边结构体
	headVex 弧起点在顶点表下标
	tailVex 弧终点在顶点表下标
	headLink 出边表指针域
	tailLink 入表表指针域
	weight 权
 */
type OrthogonalEdge struct {
	headVex, tailVex   int
	headLink, tailLink *OrthogonalEdge
	weight int //权值
	tailCap int //1表示有入边占用
}

/**
 十字链表结构体
 */
type OrthogonalList struct {
	vertex  [20]OrthogonalVertex
	numVertex, numEdgeOut,numEdgeIn int
}
/**
	构造十字链表过程
 */
func main() {
	//初始化
	orthogonalList := OrthogonalList{}

	//实例化阅读器
	inputReader := bufio.NewReader(os.Stdin)

	//输入顶点
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
	orthogonalList.numVertex = vertexNum

	//输入顶点
	for i:= 1; i<= vertexNum; i++ {
		fmt.Print("请输入第",i,"个顶点元素:")
		//读取顶点
		inputVertex,_ := inputReader.ReadString('\n')

		//去除\n
		inputVertex = strings.Replace(inputVertex,"\n","",-1)

		//构建顶点
		Vertex := OrthogonalVertex{}

		Vertex.data = inputVertex
		Vertex.firstIn = nil
		Vertex.firstOut = nil

		//存入图数组
		orthogonalList.vertex[i - 1] = Vertex
	}

	//先输入各个顶点的出边集
	fmt.Print("请输入图的出边集数目:")
	inputNumEdgeOut,err := inputReader.ReadString('\n')

	//判断
	if err != nil {
		fmt.Print("出边集数目输入错误!")
	}

	//去掉\n
	inputNumEdgeOut = strings.Replace(inputNumEdgeOut,"\n","",-1)
	numEdgeOut,_ := strconv.Atoi(inputNumEdgeOut)

	orthogonalList.numEdgeOut += numEdgeOut

	//录入出边表
	for i := 1; i <= numEdgeOut; i++ {
		fmt.Print("请输入第",i,"条出表弧(<A,B,权值>):")

		//读取
		inputEdge,_ := inputReader.ReadString('\n')

		//去除空格
		inputEdge = strings.Replace(inputEdge,"\n","",-1)

		//转化成数组
		EdgeArr := strings.Split(inputEdge,",")

		//获取起始顶点索引
		indexB := getOrthogonalVertexIndex(orthogonalList,EdgeArr[0])

		//获取终止顶点索引
		indexE := getOrthogonalVertexIndex(orthogonalList,EdgeArr[1])

		//构造边边表
		EdgeList := new(OrthogonalEdge)

		//起止索引
		(*EdgeList).headVex = indexB
		(*EdgeList).tailVex = indexE

		//权值
		(*EdgeList).weight,_ = strconv.Atoi(EdgeArr[2])

		//连上出表
		(*EdgeList).headLink = orthogonalList.vertex[indexB].firstOut
		orthogonalList.vertex[indexB].firstOut = EdgeList

	}

	//构造完成测试出表
	for i := 0; i <orthogonalList.numVertex; i++ {
		vertex := orthogonalList.vertex[i]
		fmt.Println("当前顶点为",vertex.data)
		fmt.Println("出度结点为:")

		//获取出度
		p := vertex.firstOut

		for p != nil {
			//获取出度索引
			index := p.tailVex
			fmt.Print(orthogonalList.vertex[index].data,"权为：",p.weight,";")
			p = p.headLink
		}
	}

	//入边表
	fmt.Print("\n")
	fmt.Print("请输入十字链表的入边数组数量:")

	inputNumEdgeIn,err := inputReader.ReadString('\n')
	//判断异常
	if err != nil {
		fmt.Print("输入错误!")
	}

	//去掉回车
	inputNumEdgeIn = strings.Replace(inputNumEdgeIn,"\n","",-1)
	orthogonalList.numEdgeIn,_ = strconv.Atoi(inputNumEdgeIn)

	//录入入边表
	for j := 1; j <= orthogonalList.numEdgeIn; j++ {
		fmt.Print("请输入第",j,"条入边<B,A,权>:")
		EdgeIn,_ := inputReader.ReadString('\n')

		//去除\n
		EdgeIn = strings.Replace(EdgeIn,"\n","",-1)

		//转化成数组
		EdgeInArr := strings.Split(EdgeIn,",")

		//获取起始
		indexB := getOrthogonalVertexIndex(orthogonalList,EdgeInArr[1])
		indexE := getOrthogonalVertexIndex(orthogonalList,EdgeInArr[0])

		//获取最后入边头
		indexBVertex := getFirstInHeadVertex(&orthogonalList.vertex[indexB])

		//获取最后入边尾
		indexEVertex := getFirstInTailVertex(&orthogonalList.vertex[indexE])

		//连接
		if indexBVertex == nil {
			orthogonalList.vertex[indexB].firstIn = indexEVertex
		} else {
			indexBVertex.tailLink = indexEVertex
		}

		//表示占有了
		indexEVertex.tailCap = 1

	}

	//构造完成测试入表
	for i := 0; i <orthogonalList.numVertex; i++ {
		vertex := orthogonalList.vertex[i]
		fmt.Println("当前顶点为",vertex.data)
		fmt.Println("入度结点为:")

		//获取出度
		p := vertex.firstIn

		for p != nil {
			//获取出度索引
			index := p.headVex
			fmt.Print(orthogonalList.vertex[index].data,"权为：",p.weight,";")
			p = p.tailLink
		}
	}



}
/**
根据顶点的名称来判断在链表的顶点数组的位置
 */
func getOrthogonalVertexIndex(chart OrthogonalList,vertex string) int {
	//定义返回
	rtIndex := -1

	//查找
	for i := 0; i < chart.numVertex; i++ {
		//判断是否相等
		if chart.vertex[i].data == vertex {
			rtIndex = i
			break
		}

	}

	//返回结果
	return rtIndex
}
/**
	获取入边集弧头的最后一个元素
 */
func getFirstInHeadVertex(start *OrthogonalVertex) *OrthogonalEdge {
	p := start.firstIn
	if p != nil {
		for (*p).tailLink != nil {
			p = (*p).tailLink
		}
	}

	//返回
	return p

}
/**
	获取入边集弧头的最后一个元素
 */
func getFirstInTailVertex(start *OrthogonalVertex) *OrthogonalEdge {
	p := start.firstOut

	for (*p).tailCap == 1 {
		p = (*p).headLink
	}

	//返回
	return p
}