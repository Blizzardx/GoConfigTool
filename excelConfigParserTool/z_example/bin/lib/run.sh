cd ../input
export configPath=`pwd`
cd ../tmp/

rm -rf classDefine
mkdir classDefine

cd classDefine
export classDefinePath=`pwd`

cd ../linux
sh 1_genConfigProvision ${configPath} ${classDefinePath} pb config

@IF %ERRORLEVEL% NEQ 0 pause

cd ../../lib/protobufTool

export GO_OUTDIR="../../tmp/go/"

rm -rf ${GO_OUTDIR}
mkdir ${GO_OUTDIR}

protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %GO_OUTDIR% --proto_path  %classDefinePath% %classDefinePath%/*.proto

@IF %ERRORLEVEL% NEQ 0 pause

cd ../../tmp
rm -rf import
mkdir import
cd import
mkdir config

cd config
export parserOutputPath=`pwd`

cd ../../linux
sh 2_genConfigParser ${configPath} ${parserOutputPath} pb config "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import/config"

@IF %ERRORLEVEL% NEQ 0 pause

cd ../go
export copyGoPath=`pwd`

echo begin copy...

copy *.go ${parserOutputPath}

cd ..

export GOARCH=amd64
export GOOS=linux

export CURR=`pwd`
cd ../../../../../../../../

export GOPATH=`pwd`
cd ${CURR}

go build -o linux/import github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import

@IF %ERRORLEVEL% NEQ 0 pause

cd ..
rm -rf output
mkdir output

cd output
export OutputPath=`pwd`

cd ../tmp/windows
sh import ${OutputPath} ${configPath}

@IF %ERRORLEVEL% NEQ 0 pause