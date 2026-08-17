package AndOrTree

import "fmt"

/**
先序遍历
*/

func FmtBoTree(Node *AndOrTreeNode) {
	if Node != nil {
		fmt.Println(Node.formula,Node.Fg)
		if Node.P != nil{
			fmt.Println(Node.P.formula)
		}
		if Node.Children != nil {
			for i := 0; i < len(Node.Children); i ++ {
				FmtBoTree(Node.Children[i])
			}

		}
	}

}
