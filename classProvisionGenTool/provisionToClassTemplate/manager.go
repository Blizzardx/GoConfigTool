package provisionToClassTemplate

//通过excel的描述文件 生成 *.thrift 或 *.pb 文件

import (
	"errors"
	"github.com/Blizzardx/GoConfigTool/classProvisionGenTool/define"
	"github.com/Blizzardx/GoConfigTool/common"
)

type ConfigTemplateGenerator interface {
	// 将数据转换为字节数组
	GenProvision(templateStr string, provisionClass *define.TemplateInfo) ([]byte, error)

	GetSuffix() string
}

func convertConfigInfo(packageName string, configInfo *define.ConfigInfo) (*define.TemplateInfo, error) {
	if configInfo.TableName == "" {
		return nil, errors.New(" table name can't be empty ")
	}

	result := &define.TemplateInfo{
		PackageName: packageName,
	}
	//table
	tableClass := &define.ClassInfo{
		ClassName: configInfo.TableName + "Config",
	}
	tableFieldInfo := &define.FieldInfo{
		FieldIndex: 1,
		FieldName:  "Content",
		FieldType:  configInfo.TableName + "LineInfo",
	}
	if configInfo.GlobalInfo.TableType == "map" {
		tableFieldInfo.IsMap = true
		tmpIsFoundField := false
		for _, lineElem := range configInfo.LineInfo {
			if lineElem.FieldName == configInfo.GlobalInfo.TableKeyFieldName {
				tableFieldInfo.MapKeyType = lineElem.FieldType
				tmpIsFoundField = true
				break
			}
		}
		if !tmpIsFoundField {
			return nil, errors.New("  table key field name '" + configInfo.GlobalInfo.TableKeyFieldName + "' is not found at field list")
		}
	} else if configInfo.GlobalInfo.TableType == "list" {
		tableFieldInfo.IsList = true
	} else {
		return nil, errors.New(configInfo.GlobalInfo.TableType + " is error table type ,must be 'map' or 'list' ")
	}

	tableClass.FieldList = append(tableClass.FieldList, tableFieldInfo)
	result.ClassList = append(result.ClassList, tableClass)

	//line
	lineClass := &define.ClassInfo{
		ClassName: configInfo.TableName + "LineInfo",
	}
	for index, lineElem := range configInfo.LineInfo {
		lineClass.FieldList = append(lineClass.FieldList, &define.FieldInfo{
			FieldType:  lineElem.FieldType,
			FieldName:  lineElem.FieldName,
			FieldIndex: int16(index + 1),
			IsList:     lineElem.IsList,
		})
	}
	result.ClassList = append(result.ClassList, lineClass)
	return result, nil
}

func ProvisionToClassTemplate(configInfoList []*define.ConfigInfo, outputDir string, packageName string, genTool ConfigTemplateGenerator) error {
	errorMsg := ""
	for _, configElem := range configInfoList {
		templateElem, err := convertConfigInfo(packageName, configElem)
		if nil != err {
			errorMsg += "error on config " + configElem.TableName + " " + err.Error()
			continue
		}
		content, err := genTool.GenProvision("", templateElem)
		if err != nil {
			errorMsg += "error on config " + configElem.TableName + " " + err.Error()
			continue
		}
		// write
		err = common.WriteFileByName(outputDir+"/"+configElem.TableName+genTool.GetSuffix(), content)
		if err != nil {
			errorMsg += "error on config " + configElem.TableName + " " + err.Error()
			continue
		}
	}
	return nil
}
