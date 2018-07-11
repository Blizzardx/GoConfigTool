package configDirectoryWatcher

import (
	"github.com/Blizzardx/GoConfigTool/common"
	"hash/crc32"
	"strconv"
	"time"
)

func tickFileChangeQueue() {
	go common.SafeCall(func() {
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				beginCheckFileChange()
			}
		}
	})
}
func beginCheckFileChange() {

	for i := 0; i < 16; i++ {
		if fileLoaderPool.Length() == 0 {
			return
		}
		if needLoadFileQueue.Length() == 0 {
			return
		}

		// get one file loader
		pollFileLoader()

		fileInstance := needLoadFileQueue.Poll()
		if nil == fileInstance {
			offerFileLoader()
			return
		}

		fileElem := fileInstance.(*simpleFileInfo)
		// check is loading this file
		_, ok := currentLoadingFileMap.Load(fileElem.filePath + fileElem.fileName)
		if ok {
			offerFileLoader()
			continue
		}

		doLoadFile(fileElem)
	}

}
func doLoadFile(fileElem *simpleFileInfo) {
	// mark as loading
	currentLoadingFileMap.Store(fileElem.filePath+fileElem.fileName, fileElem)
	// do loading
	go common.SafeCall(func() {
		// load from file
		fileContent, _ := common.LoadFileByName(fileElem.filePath + "/" + fileElem.fileName)
		if nil != fileContent {
			// get crc32value
			fileElem.fileSign = strconv.Itoa(int(crc32.ChecksumIEEE(fileContent)))
			// trigger to update version file
			updateVersionFile(fileElem)
		}
		// mark as done
		currentLoadingFileMap.Delete(fileElem.filePath + fileElem.fileName)
		// offer file loader
		offerFileLoader()
	})
}
func updateVersionFile(newFileElem *simpleFileInfo) {
	// add to need update queue
	needUpdateVersionFileQueue.Offer(newFileElem)

}
func offerFileLoader() {
	fileLoaderPool.Offer("fl")
}
func pollFileLoader() {
	// get one file loader
	fileLoaderPool.Poll()
}
