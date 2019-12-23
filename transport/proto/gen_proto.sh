protoc -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1/third_party/googleapis -I . --go_out=plugins=grpc:. *.proto
