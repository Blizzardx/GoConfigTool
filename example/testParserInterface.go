package main

import (
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"github.com/Blizzardx/GoConfigTool/decoder"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/excelHandler"
	"os"
)

var handlerMap = map[string]func(common.ConfigDecoder, [][]string) ([]byte, error){}

func registerParserFunc(configName string, handler func(common.ConfigDecoder, [][]string) ([]byte, error)) {
	handlerMap[configName] = handler
}
func init() {
	registerParserFunc("itemConfigTable", parserConfig_itemConfigTable)
}
func parserConfig(configName string, inputDir string, outputPath string) {
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
func parserAllConfig(outputPath string, inputDir string) {
	for fileName := range handlerMap {
		parserConfig(fileName, inputDir, outputPath)
	}
}
func main() {
	if len(os.Args) < 3 {

		fmt.Println("error input args ,need outputDir,inputDir ")
		os.Exit(1)
	}

	outputDir := os.Args[1]
	inputDir := os.Args[2]
	targetConfig := ""
	if len(os.Args) > 3 {
		targetConfig = os.Args[3]
	}

	if targetConfig == "" {
		// load all
		parserAllConfig(outputDir, inputDir)
	} else {
		parserConfig(targetConfig, inputDir, outputDir)
	}
}
