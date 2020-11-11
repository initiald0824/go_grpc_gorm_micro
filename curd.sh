#!/bin/bash

# $1 => ${tableName}

# 生成模版文件
cd cmd/
go run main.go curd -t $1

# proto文件生成go文件，包含model、validator等信息
cd ../proto/
protoc --go_out=plugins=grpc:. $1_model.proto
protoc-go-inject-tag -input=./proto/$1_model.pb.go
protoc --go_out=plugins=grpc:. --swagger_out=./proto --grpc-gateway_out=./proto $1_service.proto

# 文件中字符串替换
# 因为目前找不到好办法替换proto中的多余struct引起的报错，所以这只是暂时方法，欢迎大家有好思路提醒一下
cd ../temp/
go run main.go users