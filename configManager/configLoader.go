package configManager

import (
	"reflect"
	"log"
)

func loadFile(filePath string)error{
	// load file
	var fileContent []byte

	tmpType := typeMap[filePath]

	obj := reflect.New(tmpType).Interface()
	err := configCodeC.Decode(fileContent,nil)

	if nil != err{
		return err
	}
	configPool.Store(filePath,obj)
	return nil
}
func loadAllConfig(currentConfig *VersionConfig,newConfig *VersionConfig){
	if currentConfig != nil && currentConfig.Sign == newConfig.Sign{
		// do nothing
		log.Println("load file ,but nothing change")
		return
	}
	var needLoadConfigList []*VersionConfigElement

	if nil == currentConfig{
		// load all
		for _,fileElem := range newConfig.FileList{
			needLoadConfigList = append(needLoadConfigList,fileElem)
		}
		doLoadAllFile(needLoadConfigList,newConfig)
		return
	}

	// need compare ,which file is change
	for _,fileElem := range newConfig.FileList{
		isExist := false
		for _,oldFileElem := range currentConfig.FileList{
			if oldFileElem.FilePath == fileElem.FilePath{
				isExist = true
				if oldFileElem.Sign != fileElem.Sign || oldFileElem.Sign == ""{
					needLoadConfigList = append(needLoadConfigList,fileElem)
				}
				break
			}
		}
		if !isExist{
			needLoadConfigList = append(needLoadConfigList,fileElem)
		}
	}

	doLoadAllFile(needLoadConfigList,newConfig)
	return
}
func doLoadAllFile(needLoadConfigList []*VersionConfigElement,newVersionConfig *VersionConfig){
	// do load file
	for _,fileElem := range needLoadConfigList{
		err := loadFile(fileElem.FilePath)
		if nil != err{
			// fixed as fail
			fileElem.Sign = ""
		}
	}
	// fix version
	versionConfig = newVersionConfig
}

func onFileChange(){
	// load version config
	config,err := loadVersionConfig(configDirectory+"/"+versionConfigPath)
	if nil != err || nil == config{
		return
	}
	loadAllConfig(versionConfig,config)
}
func watchVersionFile(){

}