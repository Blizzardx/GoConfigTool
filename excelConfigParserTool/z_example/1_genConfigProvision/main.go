package main

import (
	"fmt"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/define"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/excelHandler"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/provisionToClassTemplate"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/provisionToClassTemplate/protobuf"
	"io/ioutil"
	"os"
)

//通过传进来的 xlsx文件夹路径，读取全部 excel文件，生成 类描述文件 到输出目录
// 参数 1： inputDir xlsx 文件夹路径
// 参数 2： outputdir 生成的 *.thrift ,*.proto 文件的输出目录
// 参数 3： 需要的编码类型 thrift 或者 protobuf
// 参数 4： packageName

func main() {

	if len(os.Args) < 5 {

		fmt.Println("error input args ,need inputDir,outputdir,codeType(thrift or pb) ")
		os.Exit(1)
		return
	}

	inputDir := os.Args[1]
	outputDir := os.Args[2]
	codeType := os.Args[3]
	packageName := os.Args[4]

	var coder provisionToClassTemplate.ConfigTemplateGenerator = nil
	if codeType == "thrift" {

	} else if codeType == "pb" {
		coder = &protobuf.ProtobufTemplateGenTool{}
	} else {

		fmt.Println("error input args codeType(thrift or pb) ")
		os.Exit(1)
		return
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
	err := provisionToClassTemplate.ProvisionToClassTemplate(provisionInfoList, outputDir, packageName, coder)
	if nil != err {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}
