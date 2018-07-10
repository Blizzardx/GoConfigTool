package main

import (
	"fmt"
	_ "github.com/Blizzardx/GoConfigTool/auto"
	"github.com/Blizzardx/GoConfigTool/configManager"
	"github.com/Blizzardx/GoConfigTool/decoder"
	"time"
)

func main() {
	fmt.Println("xxxx1")

	configManager.Init("log", "version.txt", new(decoder.PBDecodeC))

	fmt.Println("xxxx2")
	time.Sleep(1 * time.Minute)
}
