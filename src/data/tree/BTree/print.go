package BTree

import "fmt"

func (b *BTree) Print() {
	fmt.Print(b.Root.KeyNum)

	fmt.Println("---------------------------")
	var queue  []*BTNode
	queue = append(queue,b.Root)
	//fmt.Println(b.Root.Keys,b.Root.KeyNum)
	i := 0       //当前要出队列的下标s
	for  i < len(queue) {
		current := queue[i]
		i=i+1
		fmt.Print("[")
		for k:=1;k<=current.KeyNum;k++ {
			fmt.Printf(" %d ",current.Keys[k])
		}
		fmt.Print("]\n")
		for k:=0;k<=current.KeyNum;k++ {
			if current.Children[k] != nil {
				queue = append(queue,current.Children[k])
			}
		}
	}
}

