package main

import (
	"fmt"
	"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/sampleWorkspace/classDefineDir/auto"
	"os"
)

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
		auto.ParserAllConfig(outputDir, inputDir)
	} else {
		auto.ParserConfig(targetConfig, inputDir, outputDir)
	}
}
