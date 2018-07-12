package configDirectoryMonitor

import (
	"github.com/Blizzardx/GoConfigTool/common"
	"log"
	"sync"
	"time"
)

//需要变更的文件列表
var needLoadFileQueue = common.NewSyncQueue()

// 当前正在处理中的文件列表
var currentLoadingFileMap = &sync.Map{}

//文件处理器池子 主要用来限制消费者数量
var fileLoaderPool = common.NewSyncQueue()

//当前的版本库
var currentVersionInfo = &sync.Map{}

//需要变更刷新版本库的文件列表
var needUpdateVersionFileQueue = common.NewSyncQueue()

var currentFileState = map[string]*fileState{}

//input 文件加载器同时工作的最大数量
var maxLoaderCount = 100

//input 目标文件的所在目录
var workspaceDirectory string

//input 输出的版本索引文件名称
var versionConfigOutputName = "version.xml"

type simpleFileInfo struct {
	filePath string
	fileName string
	fileSign string
}
type fileState struct {
	filePath string
	modTime  time.Time
}

func Init(workspace string, versionConfigName string, maxFileLoaderCount int) {
	if maxFileLoaderCount > 0 && maxFileLoaderCount < 100 {
		maxLoaderCount = maxFileLoaderCount
	}
	if versionConfigName != "" {
		versionConfigOutputName = versionConfigName
	}
	if workspace == "" {
		log.Fatal("error workspace,workspace cannot be empty ")
		return
	}
	workspaceDirectory = workspace

	for i := 0; i < maxLoaderCount; i++ {
		offerFileLoader()
	}

	tickFileChangeQueue()
	tickNeedUpdateVersionFileQueue()
	checkDirectoryChange()
}
