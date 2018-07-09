package configManager

import (
	"testing"
	"go/ast"
	"fmt"
	"go/parser"
	"os"
	"go/token"
)

func Test_Get(t1 *testing.T) {
	var files []string

	files = append(files,"E:/Project/Go/configTool/configManager/configManager.go")

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
func  Parse(fileNode *ast.File) error {

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