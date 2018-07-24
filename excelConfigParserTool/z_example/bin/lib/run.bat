cd ../input
set configPath=%cd%
cd ../tmp/

rd /s/q classDefine
mkdir classDefine

cd classDefine
set classDefinePath=%cd%

cd ../windows
1_genConfigProvision.exe %configPath% %classDefinePath% pb config

@IF %ERRORLEVEL% NEQ 0 pause

cd ../../lib/protobufTool

set GO_OUTDIR="../../tmp/go/"

rd /s/q %GO_OUTDIR%
mkdir %GO_OUTDIR%

protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %GO_OUTDIR% --proto_path  %classDefinePath% %classDefinePath%/*.proto

@IF %ERRORLEVEL% NEQ 0 pause

cd ../../tmp
rd /s/q import
mkdir import
cd import
mkdir config

cd config
set parserOutputPath=%cd%

cd ../../windows
2_genConfigParser.exe %configPath% %parserOutputPath% pb config "github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import/config"

@IF %ERRORLEVEL% NEQ 0 pause

cd ../go
set copyGoPath=%cd%

echo begin copy...

copy *.go %parserOutputPath%

cd ..

set GOARCH=amd64
set GOOS=windows

set CURR=%cd%
cd ..\..\..\..\..\..\..\..\

set GOPATH=%cd%
cd %CURR%

go build -o windows/import.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/bin/tmp/import

@IF %ERRORLEVEL% NEQ 0 pause

cd ..
rd /s/q output
mkdir output

cd output
set OutputPath=%cd%

cd ../tmp/windows
import.exe %OutputPath% %configPath%

@IF %ERRORLEVEL% NEQ 0 pause