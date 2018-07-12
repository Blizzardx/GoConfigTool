package configDirectoryMonitor

import (
	"github.com/Blizzardx/GoConfigTool/common"
	"io/ioutil"
	"log"
	"time"
)

func checkDirectoryChange() {
	go common.SafeCall(func() {
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				beginCheckDirectoryState(workspaceDirectory)
			}
		}
	})
}
func beginCheckDirectoryState(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("error on check directory change " + err.Error())
		return
	}

	for _, file := range files {
		if file.IsDir() {
			beginCheckDirectoryState(directory + "/" + file.Name())
			continue
		}
		// if is version file change ,ignore
		if file.Name() == versionConfigOutputName {
			continue
		}
		// check file mod time
		if fileElem, ok := currentFileState[directory+file.Name()]; ok {
			if fileElem.modTime.Equal(file.ModTime()) {
				continue
			}
		}
		// refresh mod time
		currentFileState[directory+file.Name()] = &fileState{
			filePath: directory,
			modTime:  file.ModTime(),
		}
		// add to change list
		onFileChange(directory, file.Name())
	}
}
func onFileChange(filePath string, fileName string) {
	// add to wait queue
	needLoadFileQueue.Offer(&simpleFileInfo{filePath: filePath, fileName: fileName})
}
