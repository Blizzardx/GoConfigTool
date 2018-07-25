
cd ..

cd tmp

set GOARCH=amd64
set GOOS=windows

set CURR=%cd%
cd ..\..\..\..\..\..\..\..\

set GOPATH=%cd%
cd %CURR%

go build -o ../lib/remoteCtrl.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/remoteControl


@IF %ERRORLEVEL% NEQ 0 pause
