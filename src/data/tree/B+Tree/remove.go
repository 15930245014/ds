package B_Tree

/**
	B+树删除:
	 旋转：
	 合并：
 */
/**
删除入口
*/
func (b *BAddTree) Remove(val int) bool{
	//判断
	if b.Root == nil {
		return false
	}
	//查询当前节点
	ok,i,p,stack := b.search(val)
	//fmt.Println(stack)

	//是否查询到
	if !ok {
		//删除节点不存在
		return false
	}

	//直接删删除
	return  b.delete(p,i,stack,val)

}

/**
删除当前叶子节点
*/
func (b *BAddTree) delete(p *BAddTreeNode,i int,stack []int,val int) bool  {
	//兜底一下错误
	if i > p.keyNum {
		return false
	}

	//删除当前子节点
	for j := i ; j < p.keyNum; j ++ {
		p.keys[j] = p.keys[j + 1]
		p.data[j] = p.keys[j + 1]
		//p.Children[j] = p.Children[j + 1]
	}

	//可以删也可以不删，这里删掉
	p.keys[p.keyNum] = NULL
	p.children[p.keyNum] = Nil
	p.keyNum --

	//判断是否调整
	if p.keyNum >= HALF {

		//判断是否删除的最右边
		if i == p.keyNum + 1 {
			//调整
			b.updateIndex(p.p,stack,p.keys[p.keyNum])
		}
		return true
	}

	//调整
	if p.p != Nil {
		//先调整
		ok,t,newVal := b.restore(p)
		if ok {
			//t调节
			b.updateIndex(t.p,stack,newVal)
			return true
		}


		//合并需要更新
		tt, newVal := b.merge(p)
		b.updateIndex(tt.p,stack,newVal)
	}

	return true

}

//调整B树,该节点与父亲节点和兄弟节点之间的调整
func (b *BAddTree) restore (t *BAddTreeNode) (bool,*BAddTreeNode,int){
	//获取当前的parent
	p := t.p
	//寻找在父结点位置
	j := 0 //children是从0开始的
	for j <= p.keyNum && p.children[j] != t {
		j++
	}

	//左兄弟节点存在
	if j > 0 {   //p有左邻兄弟节点
		//获取左兄弟
		left := p.children[j - 1]

		//存在多余
		if left.keyNum > HALF {//左兄弟有多余关键字最后一个
			//右旋转
			for k := t.keyNum; k >= 1; k -- {
				t.keys[k + 1] = t.keys[k]
				t.data[k + 1] = t.data[k + 1]
				//由于是叶子节点不需要调整children
			}

			//右旋转
			t.keys[1] = left.keys[left.keyNum]
			t.data[1] = left.data[left.keyNum]
			//p.keys[j] = left.keys[left.keyNum - 1]
			left.keys[left.keyNum] = NULL

			//由于left也是叶子节点不需要调整children
			left.keyNum --
			t.keyNum ++

			//返回
			return true,t,left.keys[left.keyNum]

		}
	}

	//存在右孩子节点
	if j > 0 && j < p.keyNum { //p有右邻兄弟节点
		right := p.children[j + 1]
		//fmt.Println(right)

		//判断是否可以
		if right.keyNum > HALF {
			//right去掉keys[1]
			min := right.keys[1]

			//左移一位
			for k := 1; k < right.keyNum; k ++ {
				right.keys[k] = right.keys[k + 1]
				right.data[k] = right.data[k + 1]
			}

			//修改长度
			right.keyNum --
			t.keys[t.keyNum + 1 ] = min
			t.data[t.keyNum + 1 ] = min
			t.keyNum++

			return  true,t,min
		}

	}
	return false,t,NULL //没有调整成功,需要合并
}

/**
合并兄弟节点
*/
func (b *BAddTree) merge (t *BAddTreeNode) (*BAddTreeNode,int) {
	//获取当前的p
	p := t.p

	//寻找当前位置
	j := 0
	for j <= p.keyNum && p.children[j] != t {
		j ++
	}

	//存在左---合到左孩子
	if j > 0 {
		//获取左
		left := p.children[j - 1]

		//t合并到left
		for k := 1; k <= t.keyNum; k ++ {
			left.keyNum ++
			left.keys[left.keyNum] = t.keys[k]
			left.data[left.keyNum] = t.data[k]
		}
		//更新链表
		left.next = t.next
		if t.next != Nil {
			t.next.pre = left
		}
		t.keyNum = 0

		//修改p  ----损失keys[j]和 children[j]
		for i := j ; i < p.keyNum; i ++ {
			p.keys[i] = p.keys[i + 1]
			p.children[i] = p.children[i + 1]
		}
		p.keyNum --

		//p为空了索引上移
		if p.keyNum == 0 {
			p.keys[1] = left.keys[left.keyNum]
			p.children[0] = left
			p.children[1] = Nil
			p.keyNum++

		}
		return t,left.keys[left.keyNum]

	} else {
		//合并到右子树
		//获取right
		right := p.children[j + 1]

		//不存在右子树
		if right == Nil {
			return t,t.keys[t.keyNum]
		}
		//fmt.Println(right)

		for k := 1; k <= right.keyNum ; k ++ {
			t.keyNum++
			t.keys[t.keyNum] = right.keys[k]
		}

		//更新链表
		right.pre = t.pre
		if t.pre != Nil {
			t.pre.next = right
		}

		//更新p
		for i := j ; i < p.keyNum ; i ++ {
			p.keys[i] = p.keys[i + 1]
			p.children[i] = p.children[i + 1]
		}
		p.keyNum --

		//p为空了
		if p.keyNum == 0 {
			//p为空了索引上移
			p.keys[1] = t.keys[t.keyNum]
			p.children[0] = t
			p.children[1] = Nil
			p.keyNum++
		}

		return t,t.keys[t.keyNum]
	}

}
/**
	回溯修改索引
 */
func (b *BAddTree) updateIndex(p *BAddTreeNode,stack []int,newVal int) {
	//先出栈
	i := stack[len(stack) - 1]
	stack = stack[:len(stack) - 1]
	oldVal := p.keys[i]
	//回溯
	for p != nil {
		if p.keys[i] == oldVal && oldVal > newVal{
			p.keys[i] = newVal
			p = p.p
			if len(stack) == 0 {
				break
			}
			i = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		} else {
			break
		}
	}

}