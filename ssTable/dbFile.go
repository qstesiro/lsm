package ssTable

import (
	"encoding/binary"
	"log"
	"os"
)

/*
管理 SSTable 的磁盘文件
*/

/*

索引是从数据区开始！
0 ─────────────────────────────────────────────────────────►
◄───────────────────────────
          dataLen          ◄──────────────────
                                indexLen     ◄──────────────┐
┌──────────────────────────┬─────────────────┬──────────────┤
│                          │                 │              │
│          数据区          │   稀疏索引区    │    元数据    │
│                          │                 │              │
└──────────────────────────┴─────────────────┴──────────────┘
*/

// 将数据写入文件
func writeDataToFile(filePath string, dataArea []byte, indexArea []byte, meta MetaInfo) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(" error create file,", err)
	}
	_, err = f.Write(dataArea)
	if err != nil {
		log.Fatal(" error write file,", err)
	}
	_, err = f.Write(indexArea)
	if err != nil {
		log.Fatal(" error write file,", err)
	}
	// 写入元数据到文件末尾
	// 注意，右侧必须能够识别字节长度的类型，不能使用 int 这种类型，只能使用 int32、int64 等
	_ = binary.Write(f, binary.LittleEndian, &meta.version)
	_ = binary.Write(f, binary.LittleEndian, &meta.dataStart)
	_ = binary.Write(f, binary.LittleEndian, &meta.dataLen)
	_ = binary.Write(f, binary.LittleEndian, &meta.indexStart)
	_ = binary.Write(f, binary.LittleEndian, &meta.indexLen)
	err = f.Sync()
	if err != nil {
		log.Fatal(" error write file,", err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(" error close file,", err)
	}
}
