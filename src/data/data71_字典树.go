package main

import "fmt"

/**
	Trie树：是统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。它的优点是：最大限度地减少无谓的字符串比较，查询效率比哈希表高。
    Trie的核心思想是空间换时间。利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的。
    Trie树也有它的缺点，Trie树的内存消耗非常大。当然，或许用左儿子右兄弟的方法建树的话，可能会好点。可见，优化的点存在于建树过程中。
 和二叉查找树不同，在trie树中，每个结点上并非存储一个元素。trie树把要查找的关键词看作一个字符序列，并根据构成关键词字符的先后顺序构造用于检索的树结构。在trie树上进行检索类似于查阅英语词典。

	字典树3个基本性质：
	1、根结点不包含字符，每条边代表一个字符
	2、从根结点到某一结点，路径上的字符连接起来为该结点对用的字符串
 	3、每个节点的所有子节点包含的字符都不相同。

	/**
		应用
		字典树的应用很广泛，下面是摘自百度百科的3个应用介绍：

    1、在串快速检错中的应用

    字典树的应用很广泛，下面是摘自百度百科的3个应用简介：

    1、字典树在串的快速检索中的应用。

    给出N个单词组成的熟词表，以及一篇全用小写英文书写的文章，请你按最早出现的顺序写出所有不在熟词表中的生词。在这道题中，我们可以用字典树，先把熟词建一棵树，然后读入文章进行比较，这种方法效率是比较高的。
	还有一个就是不好单词问题  ----将所有不好单词维护成字典树然后在将匹配的单词拆分 ----类似于strstr()

    2、字典树在“串”排序方面的应用

    给定N个互不相同的仅由一个单词构成的英文名，让你将他们按字典序从小到大输出用字典树进行排序，采用数组的方式创建字典树，这棵树的每个结点的所有儿子很显然地按照其字母大小排序。对这棵树进行先序遍历即可。

    3、字典树在最长公共前缀问题的应用

    对所有串建立字典树，对于两个串的最长公共前缀的长度即他们所在的结点的公共祖先个数，于是，问题就转化为最近公共祖先问题。
给定一些键值字符串 S = [S_1,S_2 \ldots S_n][S
1
​
 ,S
2
​
 …S
n
​
 ]，我们要找到字符串 q 与 S 的最长公共前缀。 这样的查询操作可能会非常频繁。

我们可以通过将所有的键值 S 存储到一颗字典树中来优化最长公共前缀查询操作。 如果你想学习更多关于字典树的内容，可以从 208. 实现 Trie (前缀树) 开始。在字典树中，从根向下的每一个节点都代表一些键值的公共前缀。 但是我们需要找到字符串q 和所有键值字符串的最长公共前缀。 这意味着我们需要从根找到一条最深的路径，满足以下条件：

这是所查询的字符串 q 的一个前缀

路径上的每一个节点都有且仅有一个孩子。 否则，找到的路径就不是所有字符串的公共前缀

路径不包含被标记成某一个键值字符串结尾的节点。 因为最长公共前缀不可能比某个字符串本身长

4、在海量数据处理方面的应用
寻找热门查询：搜索引擎会通过日志文件把用户每次检索使用的所有检索串都记录下来，每个查询串的长度为1-255字节。假设目前有一千万个记录，这些查询串的重复读比较高，虽然总数是1千万，但是如果去除重复和，不超过3百万个。一个查询串的重复度越高，说明查询它的用户越多，也就越热门。请你统计最热门的10个查询串，要求使用的内存不能超过1G。
(1) 请描述你解决这个问题的思路；
(2) 请给出主要的处理流程，算法，以及算法的复杂度。
复杂度分析
Trie树的平均高度h为len，所以Trie树的查询复杂度为O(h)=O(len)。
建立trie的复杂度为O(n*len)。
————————————————
版权声明：本文为CSDN博主「Durant_kevin」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/u010367506/java/article/details/24002203
	 */
 /**
 	数据结构
  */
type Trie struct {
 	count int //记录该字符出现的次数
 	ch byte //记录该字符
 	mark bool //表示是否为单词
 	children [26]*Trie //子结点
}
/** Initialize your data structure here. */
func Constructor() Trie {
	//返回
	trie := new(Trie)
	trie.count = 1
	return *trie
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	//判断字符串
	if len(word) == 0 {
		return
	}

	//遍历字符串
	for i := 0; i < len(word); i ++ {
		//获取当前字符
		ch := word[i]

		//获取子节点下标
		index := int(ch - 'a')

		//判断字符是否存在
		if this.children[index] != nil {
			this.children[index].count ++
		} else {//构造
			newTrie := new(Trie)
			newTrie.count = 1
			newTrie.ch = ch

			//写入
			this.children[index] = newTrie
		}

		//重置
		this = this.children[index]

	}

	//标记
	this.mark = true

}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	//查询
	if len(word) == 0 {
		return true
	}

	//声明查询结果

	for i := 0 ; i < len(word); i ++ {
		//获取字符
		ch := word[i]

		//获取对应孩子下标
		index := int(ch - 'a')

		//判断
		if this.children[index] == nil {
			return false
		}

		this = this.children[index]
	}

	//这里可以统计单词出现的次数---


	//返回
	return this.mark == true

}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	//查询前缀对应的结点
	if len(prefix) == 0 {
		return  true
	}

	for i := 0; i < len(prefix); i ++ {
		//获取字符
		ch := prefix[i]

		//获取索引
		index := int(ch - 'a')

		//判断
		if this.children[index] == nil {
			return false
		}

		this = this.children[index]
	}

	//前缀就存在
	if this.mark == true {
		return true
	}

	//判断子节点
	for i := 0;i < 26; i ++ {
		if this.children[i] != nil {
			return  true
		}
	}

	//返回假
	return false

}

func main() {



}