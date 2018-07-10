package configManager

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type VersionConfig struct {
	Sign     string                  `xml:"sign"`
	FileList []*VersionConfigElement `xml:"fileList"`
}
type VersionConfigElement struct {
	FilePath string `xml:"filePath"`
	Sign     string `xml:"sign"`
}

func loadVersionConfig(filePath string) (*VersionConfig, error) {
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
	v := &VersionConfig{}
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
