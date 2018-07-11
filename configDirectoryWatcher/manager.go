package configDirectoryWatcher

import (
	"github.com/Blizzardx/GoConfigTool/common"
	"sync"
)

//需要变更的文件列表
var needLoadFileQueue = common.NewSyncQueue()

// 当前正在处理中的文件列表
var currentLoadingFileMap = &sync.Map{}

//文件处理器池子 主要用来限制消费者数量
var fileLoaderPool = common.NewSyncQueue()

//文件加载器同时工作的最大数量
const maxLoaderCount = 100

//当前的版本库
var currentVersionInfo = &sync.Map{}

//需要变更刷新版本库的文件列表
var needUpdateVersionFileQueue = common.NewSyncQueue()

//目标文件的所在目录
var workspaceDirectory string

//输出的版本索引文件名称
var versionConfigOutputName = "version.xml"

type simpleFileInfo struct {
	filePath string
	fileName string
	fileSign string
}

func init() {
	for i := 0; i < maxLoaderCount; i++ {
		offerFileLoader()
	}

	tickFileChangeQueue()
	tickNeedUpdateVersionFileQueue()
}
