package configManager

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type VersionConfig struct {
	Sign 		string					`xml:"sign"`
	FileList 	[]*VersionConfigElement	`xml:"fileList"`
}
type VersionConfigElement struct {
	FilePath 	string	`xml:"filePath"`
	Sign 		string	`xml:"sign"`
}

func loadVersionConfig(filePath string )(*VersionConfig,error){
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil,err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil,err
	}
	v := &VersionConfig{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil,err
	}


	fmt.Println(v)

	return v,nil
}