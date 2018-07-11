package main

import (
	_ "github.com/Blizzardx/GoConfigTool/auto"
	"github.com/Blizzardx/GoConfigTool/configManager"
	"github.com/Blizzardx/GoConfigTool/decoder"
	"log"
	"time"
)

func main() {
	log.Println("xxxx1")

	configManager.Init("config", "version.cfg", new(decoder.MsgPackDecodeC))

	log.Println("xxxx2")
	time.Sleep(1 * time.Minute)
}
