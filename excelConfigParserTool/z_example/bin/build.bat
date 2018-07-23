
set GOARCH=amd64
set GOOS=windows

set CURR=%cd%
cd ..\..\..\..\..\..\..\

set GOPATH=%cd%
cd %CURR%


rd /s/q windows
mkdir windows

go build -o windows/1_genConfigProvision.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/1_genConfigProvision
go build -o windows/2_genConfigParser.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/2_genConfigParser
go build -o windows/3_genConfigDefine.exe github.com/Blizzardx/GoConfigTool/excelConfigParserTool/z_example/3_genConfigDefine


@IF %ERRORLEVEL% NEQ 0 pause