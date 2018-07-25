package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var workDir = ""
var configInputPath = ""
var classDefineOutputPath = ""
var runtimePlatform = ""

func main() {
	if runtime.GOOS == "windows" {
		runtimePlatform = "windows"
	} else {
		runtimePlatform = "linux"
	}

	workDir = getCurrentPath()
	if len(os.Args) > 1 {
		workDir = os.Args[1]
	}

	fixWorkPath()

	if parserParentPath(workDir, 9)+"/GoConfigTool/src/github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/lib" != workDir {
		fmt.Println("error work space dir ,work dir must at /GoConfigTool/src/github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/lib")
		return
	}

	configInputPath = parserParentPath(workDir, 1) + "/input/"
	classDefineOutputPath = parserParentPath(workDir, 1) + "/tmp/classDefine/"

	importConfig("BasicItem_Common")
	return
	build()

	run()
}

func build() {

	tmpDir := parserParentPath(workDir, 1) + "/tmp/"
	ensureFolder(tmpDir)

	inputDir := parserParentPath(workDir, 1) + "/input/"
	ensureFolder(inputDir)

	compileGoProject("1_genConfigProvision", "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/1_genConfigProvision")
	compileGoProject("2_genConfigParser", "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/2_genConfigParser")
	compileGoProject("3_genConfigDefine", "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/3_genConfigDefine")

}

func run() {

	//生成 *。pb 文件
	clearFolder(classDefineOutputPath)

	callApplication("1_genConfigProvision", configInputPath, classDefineOutputPath, "pb", "config")

	//调用 pb 生成代码
	goOutDir := parserParentPath(workDir, 1) + "/tmp/go/"
	clearFolder(goOutDir)

	genProtoBuf(goOutDir, classDefineOutputPath)

	// copy go file to dir file
	importProjectPath := parserParentPath(workDir, 1) + "/tmp/import/"
	clearFolder(importProjectPath)
	importProjectConfigPath := parserParentPath(workDir, 1) + "/tmp/import/config"
	clearFolder(importProjectConfigPath)
	copyDir(goOutDir, importProjectConfigPath+"/")

	//export import project
	callApplication("2_genConfigParser", configInputPath, importProjectConfigPath, "pb", "config", "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import/config")

	//compile import project
	compileGoProject("import", "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import")

	importConfig("")
}

func importConfig(targetConfig string) {

	outputDir := parserParentPath(workDir, 1) + "/output/"

	if targetConfig == "" {
		clearFolder(outputDir)
		callApplication("import", outputDir, configInputPath)
	} else {
		callApplication("import", outputDir, configInputPath, targetConfig)
	}
}

// compile go project tool
func compileGoProject(outputName string, projectPath string) {
	platform := "windows"
	platformSuffix := ".bat"
	execProgramSuffix := ".exe"

	if runtimePlatform == "linux" {
		platform = "linux"
		platformSuffix = ".sh"
		execProgramSuffix = ""
	}
	outputDir := parserParentPath(workDir, 1) + "/tmp/" + platform + "/" + outputName + execProgramSuffix

	execCmd(workDir+"/tool/compileGoProject"+platformSuffix, "amd64", platform, parserParentPath(workDir, 8), outputDir, projectPath)
}
func genProtoBuf(outputPath string, pbInputPath string) {
	protoBufToolPath := workDir + "/protobufTool/"
	suffix := ".bat"
	if runtimePlatform == "linux" {
		suffix = ".sh"
	}

	execCmd(protoBufToolPath+"genPB"+suffix, outputPath, pbInputPath, protoBufToolPath)
}
func callApplication(appName string, arg ...string) {
	appSuffix := ".exe"
	if runtimePlatform == "linux" {
		appSuffix = ""
	}
	execCmd(parserParentPath(workDir, 1)+"/tmp/"+runtimePlatform+"/"+appName+appSuffix, arg...)
}

// path tool
func fixWorkPath() {
	workDir = strings.Replace(workDir, "\\", "/", -1)
	if workDir[len(workDir)-1] == '/' {
		workDir = workDir[0 : len(workDir)-2]
	}
}
func parserParentPath(sourcePath string, count int) string {
	fixedPath := strings.Replace(sourcePath, "\\", "/", -1)
	pathList := strings.Split(fixedPath, "/")
	length := len(pathList) - count
	if length <= 0 {
		return ""
	}
	resPath := ""
	for i := 0; i < length; i++ {
		resPath += pathList[i]
		if i < length-1 {
			resPath += "/"
		}
	}
	return resPath
}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if nil != err {
		fmt.Println(err)
		return ""
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

// folder tool
func ensureFolder(path string) {
	exist, err := pathExists(path)
	if err != nil {
		fmt.Printf("get dir error %v \n", err)
		return
	}

	if exist {
		fmt.Printf("has dir %v \n", path)
	} else {
		fmt.Printf("no dir %v \n", path)
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed %v \n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
func clearFolder(path string) {
	exist, err := pathExists(path)
	if err != nil {
		fmt.Printf("get dir error %v \n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", path)
		os.RemoveAll(path)
	}
	fmt.Printf("no dir![%v]\n", path)
	// 创建文件夹
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	} else {
		fmt.Printf("mkdir success!\n")
	}
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// cmd tool
func execCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	//err := cmd.Run()
	//if nil != err {
	//	fmt.Println(err)
	//	return err
	//}
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", string(output), err.Error())
		return err
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", string(output))
	return nil
}
func test(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	stdin.Write([]byte("go text for grep\n"))
	stdin.Write([]byte("go test text for grep\n"))
	stdin.Close()

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:" + string(out_bytes))
}

// file tool
func copyDir(src, dst string) {
	files, _ := ioutil.ReadDir(src)
	for _, fileElem := range files {
		if fileElem.IsDir() {
			continue
		}
		copy(src+fileElem.Name(), dst+fileElem.Name())
	}
}
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
