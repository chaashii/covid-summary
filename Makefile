
## pbgen: genrate protobug file
.PHONY: pbgen
pbgen:
	protoc -I. --proto_path=internals/api/v1 --go_out=pkg/api/v1 --go_opt paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=pkg/api/v1 --go-grpc_opt paths=source_relative summary.proto
# protoc --proto_path=internals/api/v1 --proto_path=thirdparty --go_out=pkg/api/v1 --go_opt paths=source_relative --go-grpc_out=pkg/api/v1 --go-grpc_opt paths=source_relative \
# --grpc-gateway_out=logtostderr=true:pkg/api/v1 --openapiv2_out=logtostderr=true,allow_merge=true,merge_file_name=api:swagger summary.proto

# go get \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
#     google.golang.org/protobuf/cmd/protoc-gen-go \
#     google.golang.org/grpc/cmd/protoc-gen-go-grpc

# github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options