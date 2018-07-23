package main

import (
	"fmt"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/define"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/excelHandler"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/genParserCode"
	"io/ioutil"
	"os"
)

// 通过传进来的 xlsx文件夹路径，读取全部 excel文件，生成文件解析 代码 到输出目录
// 参数 1： inputDir xlsx 文件夹路径
// 参数 2： outputdir 生成的 *.go 解析工程的输出路径
// 参数 3： 需要的编码类型 thrift  protobuf xml json

func main() {
	if len(os.Args) < 6 {

		fmt.Println("error input args ,need inputDir,outputdir,codeType(thrift or pb) packageName importPath")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	outputDir := os.Args[2]
	decoderType := getDecoderName(os.Args[3])
	packageName := os.Args[4]
	importPath := os.Args[5]

	if decoderType == "" {

		fmt.Println("error input args codeType (thrift pb json xml msgpack) ")
		os.Exit(1)
	}
	var provisionInfoList []*define.ConfigInfo
	//读取input文件夹下的 一级目录下所有配置文件
	files, _ := ioutil.ReadDir(inputDir)
	for _, fileElem := range files {
		if fileElem.IsDir() {
			continue
		}

		content, err := excelHandler.ReadExcelFile(inputDir + "/" + fileElem.Name())
		if err != nil {
			fmt.Println("error on parser config by name ", inputDir+fileElem.Name(), err)
			continue
		}
		provisionInfo, err := excelHandler.ParserExcelToConfigProvision(content, fileElem.Name())
		if err != nil {
			fmt.Println("error on parser config by name ", inputDir+fileElem.Name(), err)
			continue
		}
		provisionInfoList = append(provisionInfoList, provisionInfo)
	}
	if len(provisionInfoList) == 0 {
		fmt.Println("error on parser config ")

		os.Exit(1)
		return
	}

	err := genParserCode.GenParserCode(importPath, packageName, decoderType, outputDir, provisionInfoList)
	if err != nil {
		fmt.Println("error on GenParserCode config ", err)

		os.Exit(1)
	}
}
func getDecoderName(decoderType string) string {
	switch decoderType {
	case "pb":
		return "PBDecodeC"
	case "thrift":
		return "ThriftDecodeC"
	case "json":
		return "JsonDecodeC"
	case "xml":
		return "XmlDecodeC"
	case "msgpack":
		return "MsgPackDecodeC"
	}
	return ""
}
