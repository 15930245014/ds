package B_Tree

/**
	查找:根据索引查找,返回查询的data
	@return 是否查到，位置，叶子结点
 */
func (b *BAddTree) search(val int) (bool,int,*BAddTreeNode,[]int) {
	//判断
	if b == nil {
		return false,-1,Nil,nil
	}

	//从根开始
	p := b.Root

	//记录
	var Node *BAddTreeNode

	//记录索引栈
	var stack []int

	//遍历
	for p != Nil {
		//获取位置
		i := 1
		for i <= p.keyNum {
			if p.keys[i] < val {
				i ++
			} else if p.keys[i] == val {
				if p.leaf == true {
					//返回
					return true,i,p,stack
				}
				//中断下一层
				break

			} else {
				break
			}
		}

		//记录上级路径索引
		if i >= p.keyNum {
			stack = append(stack, p.keyNum)
		} else {
			stack = append(stack, i)
		}

		//判断最后
		if p.children[i - 1] == Nil {
			Node = p
		}

		//继续
		p = p.children[i - 1]

	}

	//返回
	return false,-1,Node,stack

}
