package configDirectoryWatcher

import (
	"encoding/xml"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"hash/crc32"
	"log"
	"strconv"
	"time"
)

func tickNeedUpdateVersionFileQueue() {
	go common.SafeCall(func() {
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				beginCheckVersionChange()
			}
		}
	})
}
func beginCheckVersionChange() {
	isChange := false

	for needUpdateVersionFileQueue.Length() != 0 {
		fileInstance := needUpdateVersionFileQueue.Poll()
		if nil == fileInstance {
			break
		}
		fileElem := fileInstance.(*simpleFileInfo)

		if tryUpdateVersionFile(fileElem) {
			isChange = true
		}
	}
	if isChange {
		// save version file
		saveVersionFile()
	}
}
func tryUpdateVersionFile(newFileElem *simpleFileInfo) bool {
	// check is file in map
	fileElemInterface, _ := currentVersionInfo.Load(newFileElem.filePath + newFileElem.fileName)
	if nil != fileElemInterface {
		if fileElemInterface.(*simpleFileInfo).fileSign == newFileElem.fileSign {
			// do nothing
			return false
		}
	}

	// add to new version file
	currentVersionInfo.Store(newFileElem.filePath+newFileElem.fileName, newFileElem)
	return true
}
func saveVersionFile() {
	versionConfig := &common.VersionConfig{}
	currentVersionInfo.Range(func(key, value interface{}) bool {
		fileElem := value.(*simpleFileInfo)
		versionConfig.FileList = append(versionConfig.FileList, &common.VersionConfigElement{
			FilePath: fileElem.fileName,
			Sign:     fileElem.fileSign,
		})
		return true
	})

	data, err := xml.Marshal(versionConfig.FileList)
	if data == nil || err != nil {
		log.Println("error on encode filelist to get version file sign " + err.Error())
		return
	}

	versionConfig.Sign = strconv.Itoa(int(crc32.ChecksumIEEE(data)))

	// save file

	content, err := xml.MarshalIndent(versionConfig, "  ", "    ")
	if err != nil {
		fmt.Println("error on encode version config " + err.Error())
		return
	}

	err = common.WriteFileByName(versionConfigOutputName, content)
	if err != nil {
		fmt.Println("error on write version config " + err.Error())
		return
	}
}
