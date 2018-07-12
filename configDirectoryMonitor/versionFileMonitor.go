package configDirectoryMonitor

import (
	"encoding/xml"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"hash/crc32"
	"log"
	"strconv"
	"sync"
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
	tmpChangeVersionMap := map[string]string{}
	for needUpdateVersionFileQueue.Length() != 0 {
		fileInstance := needUpdateVersionFileQueue.Poll()
		if nil == fileInstance {
			break
		}
		fileElem := fileInstance.(*simpleFileInfo)

		if tryUpdateVersionFile(fileElem) {
			isChange = true
			tmpChangeVersionMap[fileElem.filePath] = ""
		}
	}
	if isChange {
		// save version file
		saveVersionFileByChangedMap(tmpChangeVersionMap)
	}
}
func tryUpdateVersionFile(newFileElem *simpleFileInfo) bool {
	currentVersionInfoByDirInterface, _ := currentVersionInfo.Load(newFileElem.filePath)
	if nil == currentVersionInfoByDirInterface {
		currentVersionInfoByDirInterface = &sync.Map{}
		currentVersionInfo.Store(newFileElem.filePath, currentVersionInfoByDirInterface)
	}
	currentVersionInfoByDir := currentVersionInfoByDirInterface.(*sync.Map)

	// check is file in map
	fileElemInterface, _ := currentVersionInfoByDir.Load(newFileElem.filePath + newFileElem.fileName)
	if nil != fileElemInterface {
		if fileElemInterface.(*simpleFileInfo).fileSign == newFileElem.fileSign {
			// do nothing
			return false
		}
	}

	// add to new version file
	currentVersionInfoByDir.Store(newFileElem.filePath+newFileElem.fileName, newFileElem)
	return true
}
func saveVersionFileByChangedMap(changedDirectoryMap map[string]string) {
	for filePath := range changedDirectoryMap {
		currentVersionInfoByDirInterface, _ := currentVersionInfo.Load(filePath)
		if nil == currentVersionInfoByDirInterface {
			continue
		}
		currentVersionInfoByDir := currentVersionInfoByDirInterface.(*sync.Map)
		saveVersionFile(filePath, currentVersionInfoByDir)
	}
}
func saveVersionFile(filePath string, currentVersion *sync.Map) {

	versionConfig := &common.VersionConfig{}
	currentVersion.Range(func(key, value interface{}) bool {
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

	err = common.WriteFileByName(filePath+"/"+versionConfigOutputName, content)
	if err != nil {
		fmt.Println("error on write version config " + err.Error())
		return
	}

	fmt.Println("refresh version file at path " + filePath + " at time " + time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
}
