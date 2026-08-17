package main

/**
	算法描述：
	为了解决--二叉排序树不平衡的现象
	平衡因子: 结点的左子树高度和右子树的高度差值 叫做平衡因子（BF）,平衡二叉树的平衡因子值仅是 -1，0，1
	平衡二叉树本质上也是一颗二叉排序树

	最小不平衡子树：距离插入结点最近的且平衡因子的绝对值大于1的结点为根
 */

/**
	二叉平衡树数据结构
 */
type birTree struct {
	data           int //值域
	bf             int //平衡因子
	lChild, rChild *birTree
}

/**
   二叉树排序树的旋转
 */

/**
   	左旋转
 */
func leftRotate(T *birTree) { //以T为中心进行旋转
	//记录临时
	p := T

	//获取当前节点的右孩子
	rChild := p.rChild

	//T右孩子指针转向
	p.rChild = rChild.lChild

	//rChild左孩子指向
	*(rChild.lChild) = *p

	//赋值
	*T = *rChild

}
/**
	右旋转
 */
 func rightRotate(T *birTree) {
 	//记录当前
 	p := T

 	//获取当前的右孩子
 	lChild := p.lChild

 	//交换
 	p.lChild = lChild.rChild
 	*(lChild.rChild) = *p

 	//置换结点
 	*T = *lChild
 }

/**
	左平衡旋转RL和R型处理
 */
 func leftBalance(T *birTree) {
 	//获取当前的左孩子
 	lChild := T.lChild

 	//判断当前的平衡因子
 	switch lChild.bf {
	 case 1:
	 	//修改平衡因子
	 	T.bf,lChild.bf = 0,0
	 	//右旋转
	 	rightRotate(T)
	 	break
	 case -1:
	 	rChild := lChild.rChild

	 	//判断右孩子的平衡因子
		 switch rChild.bf {
		    case 1:
		  	    T.bf = -1
		  	    lChild.bf = 0
		  	    break
		 	case 0:
		 	  	T.bf,lChild.bf = 0,0
		 	 	break
		 	case -1:
		 		T.bf = 0
		 		lChild.bf = 1
				break
		 }
	 	rChild.bf = 0
	 	//先左旋转
	 	leftRotate(lChild)

	 	//在右旋转
	 	rightRotate(T)

	 	break
	}

 }
/**
   右平衡旋转L和RL型处理
*/
func rightBalance(T *birTree) {
	//获取当前的右孩子
	rChild := T.rChild

	//判断当前的平衡因子
	switch rChild.bf {
	case 1: //左子树高 先右旋转在左旋转
		lChild := rChild.lChild

		//判断左孩子孩子的平衡因子
		switch rChild.bf {
		case 1:
			T.bf = 0
			rChild.bf = -1
			break
		case 0:
			T.bf,rChild.bf = 0,0
			break
		case -1:
			T.bf = 1
			rChild.bf = 0
			break
		}
		lChild.bf = 0
		//先右旋转
		rightRotate(rChild)

		//在左旋转
		leftRotate(T)

		break
	case -1://右边高
		T.bf,rChild.bf = 0,0
		//左旋转
		leftRotate(T)
		break

	}

}
 /**
 	插入操作
 	1、值相同的元素只插入一次
 	2、新元素插入返回true或false，布尔变量taller反映T的长高与否
  */
func  InsertAVL(T *birTree, e int,taller *bool) bool {
	if T == nil {
		//插入新结点，树长高 taller为true
		*T = birTree{e,0,nil,nil}
		*taller = true

	} else {
		//判断是否相同
		if e == T.data {
			*taller = false
			return false
		}

		if e < T.data {//应在T的左子树进行搜索
			if !InsertAVL(T.lChild,e,taller) {//未插入
				return false
			}

			//表示插入T的左子树上
			if *taller {
				switch T.bf {//检查T的平衡度
					case 1: //原本左子树比右子树高
						leftBalance(T)
						*taller = false
						break

					case 0: //等高
						T.bf = 1
						*taller = true
						break

					case -1://右子树高，现在平衡了
						T.bf = 0
						*taller = false
						break
				}
			}
		} else {//在右子树添加
			if !InsertAVL(T.rChild,e,taller) {//未插入
				return false
			}

			//表示插入T的右子树上
			if *taller {
				switch T.bf {//检查T的平衡度
				case 1: //原本左子树比右子树高
					T.bf = 0
					*taller = false
					break

				case 0: //等高
					T.bf = -1
					*taller = true
					break

				case -1://右子树高,进行右平衡
					rightBalance(T)
					*taller = false
					break
				}
			}

		}
	}
	//返回
	return true
}

