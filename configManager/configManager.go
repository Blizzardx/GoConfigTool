package configManager

import (
	"github.com/Blizzardx/GoConfigTool/common"
	"log"
	"reflect"
	"sync"
	"time"
)

var currentConfigDecoder common.ConfigDecoder
var typeMaps = map[string]reflect.Type{}
var totalConfigPool = &sync.Map{}
var currentVersionConfigInfo *common.VersionConfig
var targetConfigDirectory string
var targetVersionConfigName string
var lastModifyVersionFileTime time.Time

func RegisterType(typeElem reflect.Type) {
	log.Println("register config type : " + typeElem.Name())

	if _, ok := typeMaps[typeElem.Name()]; ok {
		log.Println("already registed config type by name " + typeElem.Name())
		return
	}

	typeMaps[typeElem.Name()] = typeElem
}
func RegisterDecoder(tmpCode common.ConfigDecoder) {
	if nil == tmpCode {
		log.Println("error on set config decoder nil == tmpCode")
		return
	}

	currentConfigDecoder = tmpCode
}
func RegisterConfigPath(configDirectory string, versionConfigPath string) {
	targetVersionConfigName = versionConfigPath
	targetConfigDirectory = configDirectory
}
func Init(configDirectory string, versionConfigPath string, configDecoder common.ConfigDecoder) {
	log.Println("begin init config manager")
	RegisterConfigPath(configDirectory, versionConfigPath)
	RegisterDecoder(configDecoder)

	watchVersionFile()

	onFileChange()
}
func GetConfig(configName string) interface{} {
	v, _ := totalConfigPool.Load(configName)
	if nil == v {
		log.Println("error on get config instance by name " + configName)
	}
	return v
}
