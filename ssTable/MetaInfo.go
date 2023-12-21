package ssTable

// MetaInfo 是 SSTable 的元数据，
// 元数据出现在磁盘文件的末尾
type MetaInfo struct {
	// 版本号
	version int64
	// 数据区起始索引
	dataStart int64
	// 数据区长度
	dataLen int64
	// 稀疏索引区起始索引
	indexStart int64
	// 稀疏索引区长度
	indexLen int64
}
