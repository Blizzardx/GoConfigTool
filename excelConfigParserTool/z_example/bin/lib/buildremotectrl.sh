
cd ..

cd tmp

export GOARCH=amd64
export GOOS=linux

export CURR=`pwd`
cd ../../../../../../../../

export GOPATH=`pwd`

cd ${CURR}


go build -o ../lib/remoteCtrl github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/remoteControl

