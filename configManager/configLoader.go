package configManager

import (
	"errors"
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"time"
)

func loadFile(filePath string) error {
	// load file
	var fileContent []byte

	tmpType := typeMaps[filePath]

	if nil == tmpType {
		return errors.New("can't find type by name " + filePath)
	}
	obj := reflect.New(tmpType).Interface()
	if nil == obj {
		return errors.New("instance obj fail by name " + filePath)
	}
	err := currentConfigDecoder.Decode(fileContent, nil)

	if nil != err {
		return err
	}
	totalConfigPool.Store(filePath, obj)
	return nil
}
func loadAllConfig(currentConfig *VersionConfig, newConfig *VersionConfig) {
	if currentConfig != nil && currentConfig.Sign == newConfig.Sign {
		// do nothing
		log.Println("load file ,but nothing change")
		return
	}
	var needLoadConfigList []*VersionConfigElement

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
func doLoadTargetConfig(needLoadConfigList []*VersionConfigElement, newVersionConfig *VersionConfig) {
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

	go safeCall(func() {

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
func safeCall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(debug.Stack()))
		}
	}()
	f()
}
