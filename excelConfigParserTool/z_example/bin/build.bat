
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

:: 创建 classDefine
if exist classDefine (
		echo 目录 classDefine 已存在，无需创建
	) else (
		echo 创建classDefine
		md classDefine
	)

:: 创建 config
if exist config (
		echo 目录 config 已存在，无需创建
	) else (
		echo 创建config
		md config
	)