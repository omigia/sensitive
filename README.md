根据 AC 自动机实现敏感词匹配。

## 使用举例
```
package main

import (
	"fmt"
	"github.com/omigia/sensitive"
)

func main() {
	patterns := []string{
		"123456789",
		"1234",
		"123",
	}
	root := sensitive.Init(patterns)
	if sensitive.Match(root, "12345") {
		fmt.Println("含敏感词")
	} else {
		fmt.Println("不含敏感词")
	}
	
	if sensitive.Match(root, "34") {
		fmt.Println("含敏感词")
	} else {
		fmt.Println("不含敏感词")
	}
}	
```

## 实现
1. 根据待匹配的敏感词生成字典树(Trie)；
2. 在字典树的基础上构建 AC 自动机(Aho-Corasick)。

字典树的生成不多赘述，这里只描述下 AC 自动机的构造过程以及查找。

构造自动机的目的在于，基于已匹配成功的部分，找出该部分的最长子串，
该子串是字典树中某敏感词的前缀，减少重复匹配。

具体的查找匹配是一个递归的过程，匹配到某字符不能继续，
字典树的位置跳转至最长子串继续进行；
再匹配不到，再跳转，直至字典树的根结点。

具体构造过程关键在于树的层遍历，
子节点同父结点相关联，初始的时候只有一个根结点作为父结点。
```
待匹配的串为 ...->a->b->q->w->e，
字典树中有 root->...->a->b->c，以及其他；
...->a->b 能匹配，c 不能匹配；
要么跳转另一个长度最长的 root->...->a->b 形式，
要么跳转 root->b->... 形式，

（不考虑 root->...->a->b 中间某一段匹配，
  如果有的话，在那时候就已经结束匹配）
```
