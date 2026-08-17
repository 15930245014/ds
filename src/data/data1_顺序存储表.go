package main

import "fmt"

//构建结构体 以int类型为例
type daTA struct {
	 data []int   //存放数据集 切片
	 length int   //改顺序存储的长度
}
func main() {
	//测试构造结构体
	arr := []int{2,3,4,5}
	daTa := create(arr)
	fmt.Print(daTa)

	//测试删除
	delete(&daTa,2)
	fmt.Print(daTa)

	//测试查找
	fmt.Print(find(daTa,5))

	//测试修改
	fmt.Print(modify(daTa,5,7))


	//测试添加

	add(&daTa,2,0)

	fmt.Print(daTa.data)
}
//实现增、删、改、查 以及构建一个顺序存储结构体

//创建单链表
func create(arr []int) daTA{
	//构造返回指针
	rtData := new(daTA)

	//构造length
	(*rtData).length = len(arr)

	//存放数据集
	for _,value := range arr {
		(*rtData).data = append((*rtData).data,value)
	}

	//返回
	return *rtData
}

//删除结构体索引为i,返回删除的元素，i从1开始 i+1左移
func delete(d *daTA,i int) int {
	var rtData int

	rtData = (*d).data[i - 1]

	//遍历所以左移 索引递增 右移反之
	for s:= i - 1 ;s < (*d).length - 1; s++ {//
		(*d).data[s] = (*d).data[s + 1]
	}

	//删除最后一个元素,原来长度减1  *最末还存在一个元素未处理
	(*d).length -= 1
	L := (*d).length
	(*d).data = (*d).data[0:L]

	//返回
	return rtData

}

//查找值,寻找不到返回-1
func find(d daTA,key int)int {
	var rt  = -1
	i := 0

	for i < d.length {//利用哨兵减少判断条件
		if d.data[i] == key {
			rt = i
			break
		}
		i++
	}

	return rt

}

//修改操作target值修改为key
func modify(d daTA,target,key int) daTA {
	//查找索引
	index := find(d,target)

	//修改
	d.data[index] = key

	return d
}

//添加元素，i位置插入  长度 增加1 i之后元素右移 i从1开始
func add (d *daTA,i,key int) {
	//长度加1
	(*d).length ++

	//i-1本身及其以后元素右移 从最右边开始
	(*d).data = append((*d).data,key)  //先让长度+1
	for j := (*d).length - 1; j >= i;j-- {
		(*d).data[j] = (*d).data[j - 1]
	}

	//放入 key
	(*d).data[i - 1] = key

}



