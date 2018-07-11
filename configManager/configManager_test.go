package configManager

import (
	"encoding/xml"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/common"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func Test_Get(t1 *testing.T) {
	var files []string

	files = append(files, "E:/Project/Go/configTool/configManager/configManager.go")

	fs := token.NewFileSet()
	for _, filename := range files {

		file, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = Parse(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
func Parse(fileNode *ast.File) error {

	fmt.Println(fileNode.Name.Name)

	ast.Inspect(fileNode, func(n ast.Node) bool {

		switch typeSpec := n.(type) {
		case *ast.TypeSpec:

			switch typeSpecType := typeSpec.Type.(type) {
			case *ast.StructType:
				fmt.Println(typeSpec.Name.Name)
				fmt.Println(typeSpecType)
			}
		}

		return true
	})

	return nil
}
func Test_XmlEncode(t1 *testing.T) {
	file, err := os.Open("servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := common.VersionConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
func Test_XmlDncode(t1 *testing.T) {

	v := &common.VersionConfig{}
	v.Sign = "123123"
	for i := 0; i < 10; i++ {
		v.FileList = append(v.FileList, &common.VersionConfigElement{
			FilePath: "config/test" + strconv.Itoa(i) + ".cfg",
			Sign:     "ssss",
		})
	}
	content, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(content))
}
