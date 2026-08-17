package AndOrTree
/**
	盲目求解： 1、广度优先搜索（BFS） 2、深度优先(DFS)
 */
/*
	type AndOrTree struct {
		Root *AndOrTreeNode
		Answer int
	}

	type AndOrTreeNode struct {
		P *AndOrTreeNode
		Answer int
		formula string
		Children []*AndOrTreeNode
		Fg int //1、与树 2、或树
		rank int //秩 用于压缩树的高度 初始为0
	}
*/
/**
  dfs
 (1）把原始问题作为初始节点S0，并把它作为当前节点
（2）应用分解或等价变换操作对当前节点进行扩展
（3）为每个子节点设置指向父节点的指针 ---done
（4）选择合适的子节点作为当前节点，反复执行第（2）步和第（3）步，在此期间需要多次调用可解标记过程或不可解标记过程，直到初始节点被标记为可解节点或不可解节点为止
*/

func DfsSolve(tree *AndOrTreeNode,answer map[string]int) int {
	//判断---
	if tree == nil {
		return -1
	}

	//判断当前节点是否为叶子节点
	if len(tree.Children) == 0 {
		//返回当前
		return answer[tree.formula]
	}

	//当前为中间节点
	if tree.Fg == 1 {//与树
		for i := 0; i < len(tree.Children);i++ {
			if DfsSolve(tree.Children[i],answer) == -1 {
				return -1
			}
		}

		return 1
	} else { //或树
		for i := 0; i < len(tree.Children);i++ {
			if DfsSolve(tree.Children[i],answer) == 1 {
				return 1
			}
		}
		return -1
	}
}

/**广度优先（DFS）
（1）把初始节点S0放入Open表中
（2）把Open表的第一个节点取出放入Closed表，并记该节点为n
（3）如果节点n可扩展，则做下列工作：
扩展节点n，将其子节点放入Open表的尾部，并为每一个子节点设置指向父节点的指针
考察这些子节点中是否有终止节点。若有，则标记这些终止节点为可解节点，并用可解标记过程对其父节点及其先辈节点中的可解节点进行标记。如果初始解节点S0能够被标记为可解节点，就得到了解树，搜索成功，退出搜索过程；若果不能确定S0为可解节点，则从Open表中删去具有可解先辈的节点
转第（2）步
（4）如果节点n不可扩展，则做下列工作：
标记节点n为不可解节点
应用不可解标记过程对节点n的先辈中不可解的节点进行标记。如果初始解节点S0也被标记为不可解节点，则从Open表中删去具有不可解先辈的节点
转第（2）步
 */
func BFSSolve(root *AndOrTreeNode,answer map[string]int) int {
	//定义队列
	var open []*AndOrTreeNode //队列

	//当前节点放入open表
	open = append(open,root)

	//当队列不为空
	for len(open) > 0 {
		//当前出队
		cur := open[0]

		//判断是否标记过
		if cur.Answer != 0 {
			//出队
			open = open[1:]
			continue
		}

		//判断当前是否可扩展
		if len(cur.Children) > 0 {
			//子问题才能确定当前-先不出队
			for i := 0; i < len(cur.Children); i ++ {
				child := cur.Children[i]
				if child.Children == nil {
					//获取answer
					child.Answer = answer[child.formula]

					//标记
					p := flagAndOrTree(child)
					if p == nil {//根有结果
						break
					}

				} else { //入队
					open = append(open,child)
				}
			}
		} else { //叶子节点
			//获取answer
			cur.Answer = answer[cur.formula]

			//标记
			p := flagAndOrTree(cur)
			if p == nil {//根有结果
				break
			}

			//当前出队
			open = open[1:]
		}

		//判断当前是否成功
		if cur.Answer == 0 {
			cur.Answer = -1
			flagAndOrTree(cur)
		}
	}

	//返回
	return root.Answer
}
/**
	标记自底向上
 */
func flagAndOrTree(cur *AndOrTreeNode) *AndOrTreeNode{
	var p *AndOrTreeNode
	if cur.Answer == 1 {//当前成功 祖先如果是或树也成功
		p = cur.P
		for p != nil && p.Fg == 2 {
			p.Answer = 1
			p = p.P
		}
	} else {
		//与树 失败向上回溯
		p = cur.P
		for p != nil && p.Fg == 1 {
			p.Answer = -1
			p = p.P
		}
	}

	return p
}

/**
	层序遍历（自底向上）：略
 */


/**
	启发式：由于获取代价有io消耗（略）
	就是对当前节点按照公示维护代价 每次选择代价最小的 求解
 */
