package ssTable

import (
	"log"
	"os"
	"sync"
)

// SSTable 表，存储在磁盘文件中
type SSTable struct {
	// 文件句柄，要注意，操作系统的文件句柄是有限的
	f        *os.File
	filePath string
	// 元数据
	tableMetaInfo MetaInfo
	// 文件的稀疏索引列表,用于数据查询
	// 对于level比较大的情况数据会非常大,spareIndex全量存储在内存中不现实 ???
	sparseIndex map[string]Position
	// 排序后的key列表,用于规并排序
	sortIndex []string
	// SSTable 只能使排他锁
	lock sync.Locker
	/*
		sortIndex 是有序的，便于 CPU 缓存等，还可以使用布隆过滤器，有助于快速查找。
		sortIndex 找到后，使用 sparseIndex 快速定位
	*/
}

func (table *SSTable) Init(path string) {
	table.filePath = path
	table.lock = &sync.Mutex{}
	table.loadFileHandle()
}

// GetDbSize 获取 .db 数据文件大小
func (table *SSTable) GetDbSize() int64 {
	info, err := os.Stat(table.filePath)
	if err != nil {
		log.Fatal(err)
	}
	return info.Size()
}
