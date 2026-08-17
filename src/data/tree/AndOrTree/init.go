package AndOrTree

import (
	"errors"
	"memory/utils"
	"strings"
)

/**
	人工智能-与或树：由搜索过程形成的与或树称为搜索树。
    解树：当搜索成功时,经可解标记过程标识的由初始节点极其下属的可解节点构成的子树称为解树。

	求解：
	将原问题经过一系列的分解和等价变换来一步步求得原解

	分解：（与树）
	将一个问题p 分解成若干个子问题p1,p2,p3....(p1,p2...问题也可以看作原问题通过等价)，只有所有问题的
	都有解时 问题才有解

	等价变换：（或树）
	讲一个问题p 等价变换成若干个子问题p1,p2..; p1、p2 只要又一个满足时问题 原问题有解

	叶子结点：默认为与树或或树

 */

/**
	定义与或树
 */
type AndOrTree struct {
	Root *AndOrTreeNode
	Answer int
}

/**
	与或树子节点
 */
type AndOrTreeNode struct {
	P *AndOrTreeNode
	Answer int
	formula string
	Children []*AndOrTreeNode
	Fg int //1、与树 2、或树
	rank int //秩 用于压缩树的高度 初始为0
	score int //权值 权值越大越靠后判断
}

/**
	@param formula a & b &(e & (d & f))
 */
func InitAndOrTree(formula string) (*AndOrTree,error)  {
	//构造与或树
	if len(formula) == 0 {
		return nil , errors.New("格式化错误：formula is not empty")
	}

	//去掉空格
	formula = strings.Replace(formula," ","",-1);

	//检查有效括号和操作符
	if !utils.CheckFormula(formula) {
		return nil,errors.New("格式化错误：formula is not valid ")
	}

	//获取后缀表达式
	suffix := utils.GetSuffixFormula(formula)

	//生成树
	funcHash := make(map[string]func(*AndOrTreeNode,*AndOrTreeNode)*AndOrTreeNode)
	funcHash["&"] = func (a,b *AndOrTreeNode) *AndOrTreeNode {//&
		//降低树的高度（按&合并）
		if a.Children == nil && b.Children == nil {
			//新节点
			newNode := new(AndOrTreeNode)
			newNode.Fg = 1
			newNode.rank = 1
			newNode.Children = append(newNode.Children,a)
			newNode.Children = append(newNode.Children,b)
			newNode.formula = "(" + a.formula + "&" + b.formula + ")"
			a.P = newNode
			b.P = newNode
			return newNode

		} else {
			//判断秩
			if a.rank == b.rank {
				a.Children = append(a.Children,b)
				b.P = a
				a.formula = "(" + a.formula + "&" + b.formula + ")"
				a.rank ++
				a.Fg = 1
				return a
			} else if a.rank > b.rank{
				a.Children = append(a.Children,b)
				a.formula = "(" + a.formula + "&" + b.formula + ")"
				b.P = a
				a.Fg = 1
				return a
			} else {
				b.Children = append(b.Children,a)
				b.formula = "(" + a.formula + "&" + b.formula + ")"
				a.P = b
				b.Fg = 1
				return b
			}
		}
	}
	funcHash["|"] = func(a,b *AndOrTreeNode) *AndOrTreeNode {//｜
		//降低树的高度（按|合并）
		if a.Children == nil && b.Children == nil {
			//新节点
			newNode := new(AndOrTreeNode)
			newNode.Fg = 2
			newNode.rank = 2
			newNode.Children = append(newNode.Children,a)
			newNode.Children = append(newNode.Children,b)
			newNode.formula = "("+ a.formula + "|" + b.formula + ")"
			a.P = newNode
			b.P = newNode
			return newNode

		} else {
			//判断秩
			if a.rank == b.rank {
				a.Children = append(a.Children,b)
				a.formula = "(" + a.formula + "|" + b.formula + ")"
				a.rank ++
				a.Fg = 2
				b.P = a
				return a
			} else if a.rank > b.rank{
				a.Children = append(a.Children,b)
				a.formula = "(" + a.formula + "|" + b.formula + ")"
				b.P = a
				a.Fg = 2
				return a
			} else {
				b.Children = append(b.Children,a)
				a.P = b
				b.formula = "(" + b.formula + "|" + a.formula + ")"
				b.Fg = 2
				return b
			}
		}
	}

	//结果
	var result []*AndOrTreeNode

	//遍历后缀表达式
	for i := 0 ; i < len(suffix);i++ {
		//判断
		if _,ok := funcHash[suffix[i]]; !ok {//节点

			//初始
			node := new(AndOrTreeNode)

			//记录变量
			node.formula = suffix[i]

			//叶子节点默认为&
			node.Fg = 1

			//秩
			node.rank = 0

			//写入
			result = append(result,node)

		} else { //操作符
			//取末尾
			node1 :=  result[len(result) - 1]
			node2 :=  result[len(result) - 2]


			//弹出最后2个
			result = result[:len(result) - 2]

			//合并
			result = append(result,funcHash[suffix[i]](node1,node2))
		}
	}

	//返回
	if len(result) >= 0 {
		aoTree := new(AndOrTree)
		aoTree.Root = result[0]

		//返回
		return aoTree,nil

	} else {
		return nil,errors.New("格式化错误:生成与或树错误!")
	}

}

