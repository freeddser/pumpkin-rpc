#! /bin/bash

# 编译test.proto
#protoc -I . --go_out=plugins=grpc:. test/test.proto

# 编译hello.proto,如果不用http方式时
#protoc -I . --go_out=plugins=grpc:. hello/hello.proto

# 编译google api，新版编译器可以省略M参数
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/freeddser/pumpkin-rpc/proto/google/api:. hello_http/*.proto

# 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. hello_http/hello_http.proto

# 编译echo_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/freeddser/pumpkin-rpc/proto/google/api:. echo_http/*.proto

# 编译echo_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. echo_http/echo_http.proto


# customer_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/freeddser/pumpkin-rpc/proto/google/api:. customer_http/*.proto

# customer_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. customer_http/customer_http.proto

# monitor_http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/freeddser/pumpkin-rpc/proto/google/api:. monitor_http/*.proto

# monitor_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. monitor_http/monitor_http.proto
