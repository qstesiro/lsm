package ssTable

import (
	"github.com/whuanle/lsm/config"
	"io/ioutil"
	"log"
	"path"
	"sync"
	"time"
)

var levelMaxSize []int

// Init 初始化 TableTree
func (tree *TableTree) Init(dir string) {
	log.Println("The SSTable list are being loaded")
	start := time.Now()
	defer func() {
		elapse := time.Since(start)
		log.Println("The SSTable list are being loaded,consumption of time : ", elapse)
	}()

	// 初始化每一层 SSTable 的文件总最大值
	con := config.GetConfig()
	levelMaxSize = make([]int, 10)
	levelMaxSize[0] = con.Level0Size       // 1M
	levelMaxSize[1] = levelMaxSize[0] * 10 // 10M
	levelMaxSize[2] = levelMaxSize[1] * 10 // 100M
	levelMaxSize[3] = levelMaxSize[2] * 10 // 1G
	levelMaxSize[4] = levelMaxSize[3] * 10 // 10G
	levelMaxSize[5] = levelMaxSize[4] * 10 // 100G
	levelMaxSize[6] = levelMaxSize[5] * 10 // 1T
	levelMaxSize[7] = levelMaxSize[6] * 10 // 10T
	levelMaxSize[8] = levelMaxSize[7] * 10 // 100T
	levelMaxSize[9] = levelMaxSize[8] * 10 // 1P

	tree.levels = make([]*tableNode, 10)
	tree.lock = &sync.RWMutex{}
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Failed to read the database file")
		panic(err)
	}
	for _, info := range infos {
		// 如果是 SSTable 文件
		if path.Ext(info.Name()) == ".db" {
			tree.loadDbFile(path.Join(dir, info.Name()))
		}
	}
}
