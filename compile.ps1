$DirName = "./service"

if (Test-Path $DirName) {
    rm -r $DirName
}

# protoc -Iprotos/ecommerce product_info.proto --go-grpc_out=.
protoc -Iprotos/ecommerce product_info.proto --go_out=. --go-grpc_out=.