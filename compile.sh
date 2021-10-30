# protoc product_info.proto --go_opt=M./product_info.proto=github.com/dimitarsi/hello-grpc/service --go_out=. --go-grpc_out=.
protoc --proto_path=. protos/ecommerce/*.proto --go_out=. --go-grpc_out=.