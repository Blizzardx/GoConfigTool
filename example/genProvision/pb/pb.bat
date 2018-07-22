set CURR_DIR=%cd%

set GO_OUTDIR="../go/"

rd /s/q %GO_OUTDIR%
mkdir %GO_OUTDIR%

protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %GO_OUTDIR% --proto_path "." *.proto
pause