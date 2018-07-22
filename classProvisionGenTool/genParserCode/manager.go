package genParserCode

import "github.com/Blizzardx/GoConfigTool/classProvisionGenTool/define"

//生成配置解析代码

type ExcelConfigInfo struct {
	Content   [][]string
	Provision *define.ConfigInfo
}

func GenParserCode(outputDir string, configInfoList []*ExcelConfigInfo) error {

	return nil
}
