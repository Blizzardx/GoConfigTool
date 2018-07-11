package main

import (
	"github.com/Blizzardx/GoConfigTool/configManager"
	"github.com/Blizzardx/GoConfigTool/decoder"
	_ "github.com/Blizzardx/GoConfigTool/example/auto"
	"log"
	"time"
)

func main() {
	log.Println("xxxx1")

	configManager.Init("example/config", "version.cfg", new(decoder.MsgPackDecodeC))

	log.Println("xxxx2")
	time.Sleep(1 * time.Minute)
}
