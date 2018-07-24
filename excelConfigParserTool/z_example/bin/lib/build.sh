
cd ..

#create directory input
if [ -f "input" ];then
	echo directory input already exist
else
	echo create input
	mkdir input
fi

#create directory tmp
if [ -f "tmp" ];then
	echo directory tmp already exist
else
	echo create tmp
		mkdir tmp
fi

cd tmp

export GOARCH=amd64
export GOOS=linux

export CURR=`pwd`
cd ../../../../../../../../

export GOPATH=`pwd`

cd ${CURR}

rm -rf linux
mkdir linux

go build -o linux/1_genConfigProvision github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/1_genConfigProvision
go build -o linux/2_genConfigParser github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/2_genConfigParser
go build -o linux/3_genConfigDefine github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/3_genConfigDefine

