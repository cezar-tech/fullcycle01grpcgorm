

gen-proto: 
	protoc --go_out=./go/src/ --go_opt=paths=source_relative     --go-grpc_out=./go/src/ --go-grpc_opt=paths=source_relative    proto/product.proto