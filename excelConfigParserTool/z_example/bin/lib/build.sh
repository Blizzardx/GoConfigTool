cd ../../../../../../../
export WORKSPACE=$(cd `dirname $0`; pwd)

cd ${WORKSPACE}
export GOARCH=amd64
export GOOS=linux

export GOPATH=${WORKSPACE}

go build -o linux/1_genConfigProvision github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/1_genConfigProvision
go build -o linux/2_genConfigParser github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/2_genConfigParser
go build -o linux/3_genConfigDefine github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/3_genConfigDefine