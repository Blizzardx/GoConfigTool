package genParserCode

import (
	"bytes"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/define"
	"go/parser"
	"go/printer"
	"go/token"
	"text/template"
)

//生成配置解析代码

func GenParserCode(importPath string, packageName string, decoder string, outputDir string, provision []*define.ConfigInfo) error {

	mainClass := &GenParserCode_Main{Decoder: decoder, PackageName: packageName}
	for _, elem := range provision {
		content, err := genParserCode(packageName, elem)

		if err != nil {
			return err
		}
		err = common.WriteFileByName(outputDir+"/"+elem.TableName+"parser.go", []byte(content))
		if nil != err {
			return err
		}
		mainClass.ConfigList = append(mainClass.ConfigList, &GenParserCode_MainConfigElem{ConfigName: elem.TableName})
	}

	content, err := generateCode(codeTemplate_Main, mainClass, true)
	if nil != err {
		return err
	}
	err = common.WriteFileByName(outputDir+"/manager.go", []byte(content))
	if nil != err {
		return err
	}
	launch := &GenParserCode_Launch{ImportPackage: importPath, PackageName: packageName}
	content, err = generateCode(codeTemplate_Launch, launch, true)
	if nil != err {
		return err
	}
	err = common.WriteFileByName(outputDir+"/../main.go", []byte(content))
	if nil != err {
		return err
	}

	return nil
}
func genParserCode(packageName string, provision *define.ConfigInfo) (string, error) {

	className := provision.TableName + "Config"
	lineClassName := provision.TableName + "LineInfo"

	tableInfo := &GenParserCodeTableInfo{
		PackageName:   packageName,
		ClassName:     className,
		LineClassName: lineClassName,
	}

	// if is map
	if provision.GlobalInfo.TableType == "map" {
		// is list
		tableMap := &GenParserCode_TableMap{
			ClassName:       className,
			LineClassName:   lineClassName,
			MapKeyFieldName: provision.GlobalInfo.TableKeyFieldName,
		}
		tableListContent, err := generateCode(codeTemplate_TableMap, tableMap, false)
		if err != nil {
			return "", err
		}
		tableInfo.TableTemplate = tableListContent
	} else {
		// is list
		tableList := &GenParserCode_TableList{ClassName: className}
		tableListContent, err := generateCode(codeTemplate_TableList, tableList, false)
		if err != nil {
			return "", err
		}
		tableInfo.TableTemplate = tableListContent
	}

	for _, fieldElem := range provision.LineInfo {
		fieldTemplateInfo := &GenParserCodeFieldInfo{}
		if fieldElem.FieldName == "" {
			fieldTemplateInfo.FieldTemplate = codeTemplate_FieldEmpty
			tableInfo.LineFieldList = append(tableInfo.LineFieldList, fieldTemplateInfo)
			continue
		}

		// is not list
		if !fieldElem.IsList {

			//parser field list
			fieldInfo := &GenParserCode_Field{
				FieldType: fieldElem.FieldType,
				FieldName: fieldElem.FieldName,
			}
			// is need check limit
			if common.IsTypeCanCheckLimit(fieldElem.FieldType) {
				checkLimitInfo := &GenParserCode_CheckFieldLimit{
					FieldType:     fieldElem.FieldType,
					FieldName:     fieldElem.FieldName,
					FieldLimitMin: fieldElem.FieldValueRangeLimitMin,
					FieldLimitMax: fieldElem.FieldValueRangeLimitMax,
				}
				checkContent, err := generateCode(codeTemplate_CheckFieldLimit, checkLimitInfo, false)
				if err != nil {
					return "", err
				}
				fieldInfo.LimitCheckTemplate = checkContent
			}
			lineContent, err := generateCode(codeTemplate_Field, fieldInfo, false)
			if err != nil {
				return "", err
			}
			fieldTemplateInfo.FieldTemplate = lineContent
		} else {
			// is list
			tableInfo.IsNeedImportString = "\"strings\""

			//parser field list
			fieldInfo := &GenParserCode_FieldList{
				FieldType: fieldElem.FieldType,
				FieldName: fieldElem.FieldName,
			}
			// is need check limit
			if common.IsTypeCanCheckLimit(fieldElem.FieldType) {
				checkLimitInfo := &GenParserCode_CheckFieldListLimit{
					FieldType:     fieldElem.FieldType,
					FieldLimitMin: fieldElem.FieldValueRangeLimitMin,
					FieldLimitMax: fieldElem.FieldValueRangeLimitMax,
				}
				checkContent, err := generateCode(codeTemplate_CheckFieldListLimit, checkLimitInfo, false)
				if err != nil {
					return "", err
				}
				fieldInfo.LimitCheckTemplate = checkContent
			}
			lineContent, err := generateCode(codeTemplate_FieldList, fieldInfo, false)
			if err != nil {
				return "", err
			}
			fieldTemplateInfo.FieldTemplate = lineContent
		}

		tableInfo.LineFieldList = append(tableInfo.LineFieldList, fieldTemplateInfo)
	}

	tableContent, err := generateCode(codeTemplate, tableInfo, true)

	return tableContent, err
}

//根据模板生成代码
func generateCode(templateStr string, model interface{}, needFormat bool) (string, error) {

	var err error

	var bf bytes.Buffer

	tpl, err := template.New("Template").Parse(templateStr)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&bf, model)
	if err != nil {
		return "", err
	}

	if needFormat {
		if err = formatCode(&bf); err != nil {
			fmt.Println("format golang code err", err)
		}
	}

	return string(bf.Bytes()), nil
}

//格式化go文件
func formatCode(bf *bytes.Buffer) error {

	fset := token.NewFileSet()

	ast, err := parser.ParseFile(fset, "", bf, parser.ParseComments)
	if err != nil {
		return err
	}

	bf.Reset()

	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(bf, fset, ast)
	if err != nil {
		return err
	}

	return nil
}
