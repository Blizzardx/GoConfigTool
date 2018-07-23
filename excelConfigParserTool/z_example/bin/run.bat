cd config
set configPath=%cd%

cd ../classDefine
set classDefinePath=%cd%

cd ../windows
1_genConfigProvision.exe %configPath% %classDefinePath% pb config

@IF %ERRORLEVEL% NEQ 0 pause

cd ../protobufTool

set GO_OUTDIR="go/"

rd /s/q %GO_OUTDIR%
mkdir %GO_OUTDIR%

protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %GO_OUTDIR% --proto_path  %classDefinePath% %classDefinePath%/*.proto

@IF %ERRORLEVEL% NEQ 0 pause
echo 这里需要删除import文件夹 并创建目录结构

cd ../import/config
set parserOutputPath=%cd%

cd ../../windows
2_genConfigParser.exe %configPath% %parserOutputPath% pb config "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/import/config"

@IF %ERRORLEVEL% NEQ 0 pause

cd ../protobufTool/go
set copyGoPath=%cd%

echo begin copy...

copy *.go %parserOutputPath%

cd ../../

set GOARCH=amd64
set GOOS=windows

set CURR=%cd%
cd ..\..\..\..\..\..\..\

set GOPATH=%cd%
cd %CURR%

go build -o windows/import.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/import


@IF %ERRORLEVEL% NEQ 0 pause
