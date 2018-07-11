package configDirectoryWatcher

func onFileChange(filePath string, fileName string) {
	// add to wait queue
	needLoadFileQueue.Offer(&simpleFileInfo{filePath: filePath, fileName: fileName})
}
