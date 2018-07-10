package main

import (
	_"github.com/Blizzardx/GoConfigTool/auto"
	"fmt"
	"github.com/Blizzardx/GoConfigTool/configManager"
	"github.com/Blizzardx/GoConfigTool/decoder"
)

func main(){
	fmt.Println("xxxx")

	configManager.Init("","version.xml",new (decoder.PBDecodeC))
}
