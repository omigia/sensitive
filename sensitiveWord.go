package sensitive

import (
	"container/list"
)

type TrieNode struct {
	Data rune
	Fail *TrieNode
	Children map[rune]*TrieNode
	Reach bool
}

func Init(patterns []string) *TrieNode {
	root := new(TrieNode)
	root.Fail = root
	createTrie(root, patterns)
	createAc(root)
	return root
}

// 构建 Trie 树
func createTrie(root *TrieNode, patterns []string) {
	for _, pattern :=range patterns {
		p := root
		for _, c :=range pattern {
			if p.Children == nil {
				p.Children = make(map[rune]*TrieNode)
			}
			if _, ok := p.Children[c]; !ok {
				p.Children[c] = newTrieNode(c)
			}
			p = p.Children[c]
		}
		if p != root {
			p.Reach = true
		}
	}
}

func newTrieNode(c rune) *TrieNode {
	node := new(TrieNode)
	node.Data = c
	return node
}

// 构建自动机，层遍历：处理某结点的子结点的失败指针，并将子结点放入层遍历的队列
func createAc(root *TrieNode)  {
	root.Fail = nil
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		ele := l.Front()
		el := ele.Value.(*TrieNode)
		l.Remove(ele)
		pFailStatic := el.Fail
		for k, v := range el.Children {
			pFail := pFailStatic
			for pFail != nil {
				if _, ok := pFail.Children[k]; ok {
					v.Fail = pFail.Children[k]
					// 对于123456789、345，匹配12345678
					// 没有这句的话，会匹配不到
					if v.Fail.Reach {
						v.Reach = true
					}
					break
				} else {
					pFail = pFail.Fail
				}
			}
			if v.Fail == nil {
				v.Fail = root
			}
			l.PushBack(v)
		}
	}
	root.Fail = root
}

func Match(root *TrieNode, text string) bool {
	if len(text) == 0 {
		return false
	}
	p := root
	// 遍历 text
	for _, c :=range text {
		// 指针在不断在字典树上移动
		for {
			if _, ok := p.Children[c]; ok {
				p = p.Children[c]
				if p.Reach {
					return true
				}
				break
			} else {
				if p == root {
					break
				} else {
					p = p.Fail
				}
			}
		}
	}
	if p!=root && p.Fail.Reach {
		return true
	}
	return false
}