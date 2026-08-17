package main

import (
	"fmt"
)

/**
	哈夫曼树：给定n个权值作为n个叶子结点，构造一棵二叉树，
   若该树的带权路径长度达到最小，称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。
   哈夫曼树是带权路径长度最短的树，权值较大的结点离根较近。

  路径： 树中一个结点到另一个结点之间的分支构成这两个结点之间的路径
  路径长度：路径上的分枝数目称作路径长度。
  结点的带权路径长度：在一棵树中，如果其结点上附带有一个权值，通常把该结点的路径长度与该结点上的权值之积称为该结点的带权路径长度
 */
 /**
 	构造哈夫曼树的思想：
 　(1) 以权值分别为W1,W2．．．Ｗｎ的ｎ各结点，构成n棵二叉树T1,T2,．．．Tn并组成森林F={T1,T2,．．．Tn},其中每棵二叉树 Ti仅有一个权值为 Wi的根结点；

　　(2) 在F中选取两棵根结点权值最小的树作为左右子树构造一棵新二叉树，并且置新二叉树根结点权值为左右子树上根结点的权值之和（根结点的权值=左右孩子权值之和，叶结点的权值= Wi）；

　　(3) 从F中删除这两棵二叉树，同时将新二叉树加入到F中;

　　(4) 重复(2)、(3)直到F中只含一棵二叉树为止，这棵二叉树就是Huffman树。
  */

  /**
   哈夫曼树 数据结构
   */
   type HuffmanTree struct {
   		data int  //权值
   		lChild *HuffmanTree //左孩子
   		rChild *HuffmanTree //右孩子
   		fg byte  //字母位
   }


   /**
     构造huffman tree
     mp [27=>'A','8'=> 'B']
   	 arr [27,8]
   */
   func makeHuffman(mp map[int]byte,arr []int) *HuffmanTree{
   		//声明huffmanTree数组
   		var treeS []*HuffmanTree

   		//构造huffmanTree
   		for _,value := range arr {
   			tempTree := new (HuffmanTree)
   			tempTree.data = value
   			tempTree.fg = mp[value]
   			treeS = append(treeS,tempTree)

		}

   		//判断
   		if len(treeS) == 1 {
   			//返回
   			return treeS[0]
		}


		//判断
   		for len(treeS) > 1 {
   			//排序
   			sortHuffmanTree(&treeS)

   			//取出前2个,1个作为左，一个作为右，并值域相加
   			left := shift(&treeS)
   			right := shift(&treeS)

   			//构造新的
   			newTree := new(HuffmanTree)
   			newTree.data = left.data + right.data
   			newTree.lChild = left
   			newTree.rChild = right

   			//全职和插入数组
   			treeS = append(treeS,newTree)

		}

   		//返回最后的根结点
   		return treeS[0]
   }
   /**
   	弹出切片的第一个元素
    */
   func shift (arr *[]*HuffmanTree) *HuffmanTree{
   	 //记录元素的值
   	 rtTree := (*arr)[0]

   	 //弹出第一个元素
   	 *arr = (*arr)[1:len(*arr)]

   	 //返回结果
   	 return rtTree

   }


   /**
   排序huffmanTree
    */
   func sortHuffmanTree(arr *[]*HuffmanTree){

   	   //采用插入排序
	   for i := 1; i < len(*arr); i++ {
		   //记录当前
		   current := (*arr)[i]

		   //进行比较
		   j := i - 1
		   for j >= 0 && (*arr)[j].data > current.data {
			   (*arr)[j + 1] = (*arr)[j]
			   j --
		   }

		   //交换
		   (*arr)[j + 1] = current

	   }

   }

   /**
   	中序遍历huffmanTree
    */
    func blHuffmanTree(huffmanTree *HuffmanTree) {
    	//判断条件
    	if huffmanTree != nil {
    		//遍历左
    		blHuffmanTree(huffmanTree.lChild)

    		//遍历结点
    		fmt.Println(huffmanTree.data)

    		//遍历右
    		blHuffmanTree(huffmanTree.rChild)
		}
	}


   func main() {
   	arr := []int{1,2,3,4}
    mp := map[int]byte{}
    mp[1] = 'A'
    mp[2] = 'B'
    mp[3] = 'C'
    mp[4] = 'D'


   	//构造huffmanTree
   	huffmanTree := makeHuffman(mp,arr)

   	//中序遍历
   	blHuffmanTree(huffmanTree)

   }



