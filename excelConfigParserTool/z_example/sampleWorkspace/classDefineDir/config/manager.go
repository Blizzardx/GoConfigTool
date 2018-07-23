package config

import (
	"fmt"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/excelHandler"
	"github.com/Blizzardx/GoConfigTool/common"
	"github.com/Blizzardx/GoConfigTool/decoder"
)

var handlerMap = map[string]func(common.ConfigDecoder, [][]string) ([]byte, error){}

func init() {

	registerParserFunc("BasicItem_Common", parserConfig_BasicItem_CommonConfig)

}
func registerParserFunc(configName string, handler func(common.ConfigDecoder, [][]string) ([]byte, error)) {
	handlerMap[configName] = handler
}
func ParserConfig(configName string, inputDir string, outputPath string) {
	handler := handlerMap[configName]
	if nil == handlerMap {
		fmt.Println("error on get handler by config name ", configName)
		return
	}
	// read config
	content, err := excelHandler.ReadExcelFile(inputDir + "/" + configName + ".xlsx")
	if nil != err {
		fmt.Println("error on get file provison file ", err.Error(), configName)
		return
	}
	result, err := handler(&decoder.PBDecodeC{}, excelHandler.FixExcelFile(content))
	if nil != err {
		fmt.Println("error on parser config ", err.Error(), configName)
		return
	}
	err = common.WriteFileByName(outputPath+"/"+configName+".bytes", result)
	if nil != err {
		fmt.Println("error on write result config ", err.Error(), configName)
		return
	}
}
func ParserAllConfig(outputPath string, inputDir string) {
	for fileName := range handlerMap {
		ParserConfig(fileName, inputDir, outputPath)
	}
}
