package B_Tree

import "fmt"

func (b *BAddTree) Print() {
	fmt.Println("---------------------------")
	var queue  []*BAddTreeNode
	queue = append(queue,b.Root)
	for len(queue) > 0 {
		tmp := []*BAddTreeNode{}
		for i := 0; i < len(queue); i++ {
			fmt.Print("[")
			p := queue[i]
			for j := 1; j <= p.keyNum; j++ {
				fmt.Print(p.keys[j])
				fmt.Print(" ")
			}
			fmt.Print("]")
			fmt.Print("   ")

			for j := 0; j <= p.keyNum; j++ {
				if p.children[j] != Nil {
					//fmt.Println(p.children[j])
					tmp = append(tmp, p.children[j])
				}
			}

		}
		queue = tmp
		fmt.Print("\n")
	}

}

