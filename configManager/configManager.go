package configManager

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

var currentConfigDecoder ConfigDecoder
var typeMaps = map[string]reflect.Type{}
var totalConfigPool = &sync.Map{}
var currentVersionConfigInfo *VersionConfig
var targetConfigDirectory string
var targetVersionConfigName string

func RegisterType(typeElem reflect.Type) {
	fmt.Println("register config type : " + typeElem.Name())

	if _, ok := typeMaps[typeElem.Name()]; ok {
		log.Fatal("already registed config type by name " + typeElem.Name())
		return
	}

	typeMaps[typeElem.Name()] = typeElem
}
func RegisterDecoder(tmpCode ConfigDecoder) {
	if nil == tmpCode {
		log.Fatal("error on set config decoder nil == tmpCode")
		return
	}

	currentConfigDecoder = tmpCode
}
func RegisterConfigPath(configDirectory string, versionConfigPath string) {
	targetVersionConfigName = versionConfigPath
	targetConfigDirectory = configDirectory
}
func Init(configDirectory string, versionConfigPath string, configDecoder ConfigDecoder) {
	RegisterConfigPath(configDirectory, versionConfigPath)
	RegisterDecoder(configDecoder)

	watchVersionFile()

	onFileChange()
}
func GetConfig(configName string) interface{} {
	v, _ := totalConfigPool.Load(configName)
	if nil == v {
		log.Fatal("error on get config instance by name " + configName)
	}
	return v
}
