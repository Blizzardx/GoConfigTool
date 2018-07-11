package configManager

import (
	"errors"
	"github.com/Blizzardx/GoConfigTool/common"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func loadFile(filePath string) error {
	fileType := strings.Replace(filePath, "\\", "/", -1)
	tmpSlist := strings.Split(fileType, "/")
	if len(tmpSlist) < 1 {
		return errors.New("can't parser type by name " + fileType)
	}
	tmpSlist = strings.Split(tmpSlist[len(tmpSlist)-1], ".")
	if len(tmpSlist) < 1 {
		return errors.New("can't parser type by name " + fileType)
	}
	fileType = tmpSlist[0]

	// load file
	fileContent, err := common.LoadFileByName(filePath)
	if nil == fileContent || err != nil {
		return errors.New("can't load file by path type by name " + filePath)
	}

	tmpType := typeMaps[fileType]

	if nil == tmpType {
		return errors.New("can't find type by name " + fileType)
	}
	obj := reflect.New(tmpType).Interface()
	if nil == obj {
		return errors.New("instance obj fail by name " + fileType)
	}
	err = currentConfigDecoder.Decode(fileContent, obj)

	if nil != err {
		return errors.New("error on decode file ty struct by name " + filePath + " " + err.Error())
	}
	totalConfigPool.Store(fileType, obj)
	return nil
}
func loadAllConfig(currentConfig *common.VersionConfig, newConfig *common.VersionConfig) {
	if currentConfig != nil && currentConfig.Sign == newConfig.Sign {
		// do nothing
		log.Println("load file ,but nothing change")
		return
	}
	var needLoadConfigList []*common.VersionConfigElement

	if nil == currentConfig {
		// load all
		for _, fileElem := range newConfig.FileList {
			needLoadConfigList = append(needLoadConfigList, fileElem)
		}
		doLoadTargetConfig(needLoadConfigList, newConfig)
		return
	}

	// need compare ,which file is change
	for _, fileElem := range newConfig.FileList {
		isExist := false
		for _, oldFileElem := range currentConfig.FileList {
			if oldFileElem.FilePath == fileElem.FilePath {
				isExist = true
				if oldFileElem.Sign != fileElem.Sign || oldFileElem.Sign == "" {
					needLoadConfigList = append(needLoadConfigList, fileElem)
				}
				break
			}
		}
		if !isExist {
			needLoadConfigList = append(needLoadConfigList, fileElem)
		}
	}

	doLoadTargetConfig(needLoadConfigList, newConfig)
	return
}
func doLoadTargetConfig(needLoadConfigList []*common.VersionConfigElement, newVersionConfig *common.VersionConfig) {
	// do load file
	for _, fileElem := range needLoadConfigList {
		err := loadFile(fileElem.FilePath)
		if nil != err {
			// fixed as fail
			fileElem.Sign = ""
			log.Println(err)
		}
	}
	// fix version
	currentVersionConfigInfo = newVersionConfig
}
func onFileChange() {
	// load version config
	config, err := loadVersionConfig(targetConfigDirectory + "/" + targetVersionConfigName)
	if nil != err || nil == config {
		return
	}
	loadAllConfig(currentVersionConfigInfo, config)
}
func watchVersionFile() {

	go common.SafeCall(func() {

		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				beginCheckVersionFileChange()
			}
		}
	})
}
func beginCheckVersionFileChange() {
	fileInfo, err := os.Stat(targetConfigDirectory + "/" + targetVersionConfigName)
	if err != nil {
		log.Println(err)
	}
	if fileInfo != nil {
		modTime := fileInfo.ModTime()
		if !modTime.Equal(lastModifyVersionFileTime) {
			log.Println("version file is changed ,begin reload ", modTime, lastModifyVersionFileTime)

			onFileChange()
		}
	}
}
