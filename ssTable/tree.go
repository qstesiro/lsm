package ssTable

import (
	"sync"
)

// TableTree 树
type TableTree struct {
	levels []*tableNode // 无头链表,每个索引位代表一层
	// 用于避免进行插入或压缩、删除 SSTable 时发生冲突
	lock *sync.RWMutex
}

// 链表，表示每一层的 SSTable
type tableNode struct {
	index int
	table *SSTable
	next  *tableNode
}

// 插入一个 SSTable 到指定层
func (tree *TableTree) insert(table *SSTable, level int) (index int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	// 每次插入的，都出现在最后面
	node := tree.levels[level]
	newNode := &tableNode{
		table: table,
		next:  nil,
		index: 0,
	}

	if node == nil {
		tree.levels[level] = newNode
	} else {
		for node != nil {
			if node.next == nil {
				newNode.index = node.index + 1
				node.next = newNode
				break
			} else {
				node = node.next
			}
		}
	}
	return newNode.index
}
