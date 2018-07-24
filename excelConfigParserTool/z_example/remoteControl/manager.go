package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(getCurrentPath())
	return
	execCmd("E:/porject/GoConfigTool/src/github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/lib/tool/compileGoProject.bat",
		"amd64",
		"windows",
		"E:/porject/GoConfigTool/",
		"windows/3_genConfigDefine",
		"github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/3_genConfigDefine")

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
