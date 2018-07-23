package protobuf

import (
	"bytes"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/define"
	"text/template"
)

const defaultCodeTemplate = `// Generated by gen-tool
// DO NOT EDIT!
syntax = "proto3";

package {{.PackageName}};

{{if gt (.ImportPackageList|len) 0}}

{{range .ImportPackageList}}
import {{.ImportPackage}};
{{end}}

{{end}}

{{range .ClassList}}
{{.ClassType}} {{.ClassName}}{

{{range .FieldList}}
	{{.FieldType}} {{.FieldName}} = {{.FieldIndex}};
	{{end}}

}
	{{end}}

`

type PBTemplateInfo struct {
	PackageName       string
	ImportPackageList []*PBImportPkgInfo
	ClassList         []*PBClassInfo
}
type PBImportPkgInfo struct {
	ImportPackage string
}
type PBClassInfo struct {
	ClassType string
	ClassName string
	FieldList []*PBField
}
type PBField struct {
	FieldType  string
	FieldName  string
	FieldIndex int16
}
type ProtobufTemplateGenTool struct {
}

func (self *ProtobufTemplateGenTool) GenProvision(templateStr string, provisionClass *define.TemplateInfo) ([]byte, error) {
	if templateStr == "" {
		templateStr = defaultCodeTemplate
	}
	templateInfo := convertToPBInfo(provisionClass)

	return generateCode(templateStr, templateInfo)
}
func (self *ProtobufTemplateGenTool) GetSuffix() string {
	return ".proto"
}
func convertToPBInfo(provisionClass *define.TemplateInfo) *PBTemplateInfo {
	result := &PBTemplateInfo{
		PackageName: provisionClass.PackageName,
	}
	for _, importElem := range provisionClass.ImportPackageList {
		result.ImportPackageList = append(result.ImportPackageList, &PBImportPkgInfo{ImportPackage: importElem.ImportPackage})
	}
	for _, classElem := range provisionClass.ClassList {
		pbClassInfo := &PBClassInfo{
			ClassType: "message",
			ClassName: classElem.ClassName,
		}
		for _, fieldElem := range classElem.FieldList {
			pbFieldInfo := &PBField{
				FieldType:  convertToPBFieldType(fieldElem.FieldType),
				FieldName:  fieldElem.FieldName,
				FieldIndex: fieldElem.FieldIndex,
			}
			if fieldElem.IsList {
				pbFieldInfo.FieldType = "repeated " + pbFieldInfo.FieldType
			} else if fieldElem.IsMap {
				//map<keyType, valueType>
				pbFieldInfo.FieldType = "map<" + fieldElem.MapKeyType + "," + pbFieldInfo.FieldType + ">"
			}
			pbClassInfo.FieldList = append(pbClassInfo.FieldList, pbFieldInfo)
		}

		result.ClassList = append(result.ClassList, pbClassInfo)
	}
	return result
}
func convertToPBFieldType(fieldTypeStr string) string {
	fieldType := define.ConvertStrToFieldType(fieldTypeStr)

	switch fieldType {
	case define.FieldType_Int32:
		return "int32"
	case define.FieldType_Int64:
		return "int64"
	case define.FieldType_Float32:
		return "float"
	case define.FieldType_Float64:
		return "double"
	case define.FieldType_String:
		return "string"
	case define.FieldType_Bool:
		return "bool"
	case define.FieldType_Custom:
		return fieldTypeStr
	default:
		return fieldTypeStr
	}
}

//根据模板生成代码
func generateCode(templateStr string, model interface{}) ([]byte, error) {

	var err error

	var bf bytes.Buffer

	tpl, err := template.New("Template").Parse(templateStr)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(&bf, model)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bf.Bytes()))

	return bf.Bytes(), nil
}