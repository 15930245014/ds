package BTree

/**
	删除入口
 */
func (b *BTree) Remove(val int) bool{
	//判断
	if b.Root == nil {
		return false
	}
	//查询当前节点
	ok,i,p := b.Search(val)

	//是否查询到
	if !ok {
		//删除节点不存在
		return false
	}

	//判断是否为叶子结点
	if p.Children[0] != nil {
		//非叶子节点
		//查找后继
		nK,node := b.findAfterMinNode(p,i)

		//后继替换当前
		p.Keys[i] = nK

		//删除node
		return b.delete(node,1)

	} else { //叶子结点
		//直接删删除
		//fmt.Println(p.Keys,i)
		return  b.delete(p,i)
	}

}

/**
	删除当前叶子节点
 */
func (b *BTree) delete(p *BTNode,i int) bool  {
	//兜底一下错误
	if i > p.KeyNum {
		return false
	}

	//删除当前
	for j := i ; j < p.KeyNum; j ++ {
		p.Keys[j] = p.Keys[j + 1]
		//p.Children[j] = p.Children[j + 1]
	}

	//可以删也可以不删，这里删掉
	p.Keys[p.KeyNum] = NULL
	p.Children[p.KeyNum] = Nil
	p.KeyNum --

	//判断是否调整
	if p.KeyNum >= HALF {
		return true
	}

	//调整
	if p.P != Nil {
		//先调整
		ok, _ := b.Restore(p)
		if ok {
			return true
		}

		//合并
		b.merge(p)
	}

	return true

}


/**
	获取右子树最小节点
 */
func (b *BTree) findAfterMinNode(node *BTNode,i int) (int,*BTNode) {
	//判断
	if node == Nil || node.Children[i] == Nil {
		return -1,Nil
	}

	//获取最小值
	p := node.Children[i]
	for p.Children[0] != nil {
		p = p.Children[0]
	}

	//返回
	return p.Keys[1],p
}

//调整B树,该节点与父亲节点和兄弟节点之间的调整
func (b *BTree) Restore (t *BTNode) (bool,*BTNode){
	//获取当前的parent
	p := t.P

	//寻找在父结点位置
	j := 0 //children是从0开始的
	for j <= p.KeyNum && p.Children[j] != t {
		j++
	}

	//左兄弟节点存在
	if j > 0 {   //p有左邻兄弟节点
		//获取左兄弟
		left := p.Children[j - 1]

		//存在多余
		if left.KeyNum > HALF {//左兄弟有多余关键字最后一个
			//右旋转
			for k := t.KeyNum; k >= 1; k -- {
				t.Keys[k + 1] = t.Keys[k]
				//由于是叶子节点不需要调整children
			}

			//右旋转
			t.Keys[1] = p.Keys[j]
			p.Keys[j] = left.Keys[left.KeyNum]
			left.Keys[left.KeyNum] = NULL

			//由于left也是叶子节点不需要调整children
			left.KeyNum --
			t.KeyNum ++

			//返回
			return true,t

		}
	}

	//存在右孩子节点
	if j > 0 && j < p.KeyNum { //p有右邻兄弟节点
		right := p.Children[j + 1]

		//判断是否可以
		if right.KeyNum > HALF {
			//right去掉keys[1]
			min := right.Keys[1]

			//左移一位
			for k := 1; k < right.KeyNum; k ++ {
				right.Keys[k] = right.Keys[k + 1]
			}

			//修改长度
			right.KeyNum --
			t.Keys[t.KeyNum + 1 ] = p.Keys[j + 1]
			t.KeyNum++
			p.Keys[j + 1] = min

			return  true,t
		}

	}
	return false,t //没有调整成功,需要合并
}

/**
	合并兄弟节点
 */
func (b *BTree) merge (t *BTNode) *BTNode {
	//获取当前的p
	p := t.P

	//寻找当前位置
	j := 0
	for j <= p.KeyNum && p.Children[j] != t {
		j ++
	}

	//存在左---合到左孩子
	if j > 0 {
		//获取左
		left := p.Children[j - 1]

		//父亲key下移
		left.KeyNum ++
		left.Keys[left.KeyNum] = p.Keys[j]

		//t合并到left
		for k := 1; k <= t.KeyNum; k ++ {
			left.KeyNum ++
			left.Keys[left.KeyNum] = t.Keys[k]
		}

		//修改p  ----损失keys[j]和 children[j]
		for i := j ; i < p.KeyNum; i ++ {
			p.Keys[i] = p.Keys[i + 1]
			p.Children[i] = p.Children[i + 1]
		}
		p.KeyNum --

		//p为空了
		if p.KeyNum == 0 {
			left.P = p.P
			*p = *left

		}

	} else {
		//合并到右子树
		//获取right
		right := p.Children[j + 1]

		//不存在右子树
		if right == Nil {
			return p
		}

		//父亲下移
		t.KeyNum ++
		t.Keys[t.KeyNum] = p.Keys[j + 1]
		for k := 1; k <= right.KeyNum ; k ++ {
			t.KeyNum++
			t.Keys[t.KeyNum] = right.Keys[k]
		}

		//更新p
		for i := j ; i < p.KeyNum ; i ++ {
			p.Keys[i] = p.Keys[i + 1]
			p.Children[i] = p.Children[i + 1]
		}
		p.KeyNum --

		//p为空了
		if p.KeyNum == 0 {
			right.P = p.P
			*p = *right
		}
	}

	return p
}