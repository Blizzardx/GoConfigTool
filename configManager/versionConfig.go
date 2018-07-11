package configManager

import (
	"encoding/xml"
	"github.com/Blizzardx/GoConfigTool/common"
	"io/ioutil"
	"log"
	"os"
)

func loadVersionConfig(filePath string) (*common.VersionConfig, error) {
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	v := &common.VersionConfig{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	info, err := file.Stat()
	if nil == err && info != nil {
		lastModifyVersionFileTime = info.ModTime()
		log.Println("refresh last change time ", lastModifyVersionFileTime)
	}

	return v, nil
}
