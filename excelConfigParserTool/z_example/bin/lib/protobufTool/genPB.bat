echo GO_OUTDIR = %1
echo classDefinePath = %2
cd %3
protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe --go_out %1 --proto_path  %2 %2/*.proto